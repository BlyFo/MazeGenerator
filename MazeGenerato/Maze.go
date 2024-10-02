package mazeDrawer

import (
	"fmt"
	"math"
	"math/rand/v2"
	Cells "mazeGenerator/MazeGenerato/Cells"
	Utils "mazeGenerator/Utils/2DVector"
)

type Vector2D = Utils.Vector2D
type MazeCell = Cells.MazeCell
type MazeEdge = Cells.MazeEdge

func isPositionOutOfBounds(height int, width int, position Vector2D) bool {
	//values are not exact to take in to consideration rounding errors
	return float32(width)+0.05 < position.X || float32(height)+0.05 < position.Y || position.X < -0.05 || position.Y < -0.05
}

func getCellByPosition(cells []*MazeCell, position Vector2D) (*MazeCell, int) {
	if len(cells) == 0 {
		return nil, -1
	}

	for i := 0; i < len(cells); i++ {
		distance := cells[i].Position.Distance(position)
		if distance <= 0.10 {
			return cells[i], i
		}
	}

	return nil, -1
}

func getEdgeByDirection(edges *[]MazeEdge, direction Vector2D) *MazeEdge {
	// this assumes that the directions are normalized
	for i := 0; i < len(*edges); i++ {
		distance := (*edges)[i].Direction.DotProduct(direction)
		if distance >= 0.90 { //1 means that they are the same vector (same direction)
			return &(*edges)[i]
		}
	}
	fmt.Println(edges, direction)
	panic("this should never happen")
}

func getCellEmptyEdges(cell MazeCell) []Vector2D {
	var result []Vector2D
	for _, edge := range cell.Edges {
		if edge.EdgeType == Cells.NONE {
			result = append(result, edge.Direction)
		}
	}
	return result
}

func getRandomUnfinishCell(unfinishCells []*MazeCell, exitFound bool) (int, *MazeCell) {
	var cellIndex int
	if exitFound {
		cellIndex = rand.IntN(len(unfinishCells))
	} else {
		cellIndex = len(unfinishCells) - 1
	}
	return cellIndex, unfinishCells[cellIndex]
}

func vectorToRadians(vector Vector2D) float64 {
	// to make sure the rotation starts from (0,1) in clockwise order
	radians := math.Atan2(float64(vector.Y), float64(vector.X)) - math.Pi/2

	if radians < 0 {
		radians += 2 * math.Pi // Asegurarse que esté en [0, 2π]
	}

	return radians
}

func createNewCell(prviousCell *Cells.MazeCell, shape Cells.Cell, direction Vector2D, position Vector2D) Cells.MazeCell {
	var newCellDirection Vector2D
	if prviousCell == nil {
		newCellDirection = shape.GetStartingDirection()
	} else {
		newCellDirection = shape.GetNewDirection(prviousCell.Direction, direction)
	}

	newCell := MazeCell{
		Position:   position,
		Direction:  newCellDirection,
		Rotation:   vectorToRadians(newCellDirection),
		Edges:      shape.GetEdges(),
		IsExitPath: false,
	}

	newCellEdge := getEdgeByDirection(&newCell.Edges, shape.GetOpositeDirection(direction))
	newCellEdge.EdgeType = Cells.ROAD
	newCellEdge.Object = prviousCell

	if prviousCell == nil {
		newCellEdge.EdgeType = Cells.ENTRY
	}

	return newCell
}

func checkEdgesAndRemoveFromListIfEmpty(cell *Cells.MazeCell, cellIndex int, cellList *[]*Cells.MazeCell, exitFound bool) bool {
	emptyEdges := getCellEmptyEdges(*cell)
	wasCellRemoved := false
	//if that was the last empty edge remove the cell from the list
	if len(emptyEdges) == 0 {
		//if no exit was found mark it as no exith path
		cell.IsExitPath = cell.IsExitPath && exitFound
		(*cellList)[cellIndex] = (*cellList)[len(*cellList)-1]
		*cellList = (*cellList)[:len(*cellList)-1]
		wasCellRemoved = true
	}
	return wasCellRemoved
}

func CreateMaze(height int, width int, shape Cells.Cell) Cells.Maze {
	mazeCells := make([]MazeCell, 0, height*width)
	notFinishCells := make([]*MazeCell, 0, height*width/2) //worst case escenario (I think)
	CellSeparation := shape.GetDistanceBetweenCells()
	exitFound := false

	firstCell := createNewCell(nil, shape, shape.GetStartingDirection(), Vector2D{X: float32(rand.IntN(width + 1)), Y: 0})
	firstCell.IsExitPath = true

	mazeCells = append(mazeCells, firstCell)
	notFinishCells = append(notFinishCells, &mazeCells[len(mazeCells)-1])

	for len(notFinishCells) > 0 {

		randomCellIndex, randomCell := getRandomUnfinishCell(notFinishCells, exitFound)
		if checkEdgesAndRemoveFromListIfEmpty(randomCell, randomCellIndex, &notFinishCells, exitFound) {
			continue
		}
		emptyEdges := getCellEmptyEdges(*randomCell)

		directionVector := emptyEdges[rand.IntN(len(emptyEdges))]
		randomCellEdge := getEdgeByDirection(&randomCell.Edges, directionVector)

		nextCellRotation := directionVector.Rotate(vectorToRadians(randomCell.Direction))
		nextCellDistance := Vector2D{X: nextCellRotation.X * CellSeparation.X, Y: nextCellRotation.Y * CellSeparation.Y}
		nextCellPosition := randomCell.Position.Add(nextCellDistance)

		if isPositionOutOfBounds(height, width, nextCellPosition) {
			if nextCellPosition.Y > float32(height) && !exitFound {
				randomCellEdge.EdgeType = Cells.EXIT
				exitFound = true
				continue
			}

			randomCellEdge.EdgeType = Cells.WALL
			continue
		}

		nextCell, nextCellIndex := getCellByPosition(notFinishCells, nextCellPosition)

		if nextCell == nil {
			//if no cell was found a new one will be created
			newCell := createNewCell(randomCell, shape, directionVector, nextCellPosition)
			newCell.IsExitPath = !exitFound

			mazeCells = append(mazeCells, newCell)
			notFinishCells = append(notFinishCells, &mazeCells[len(mazeCells)-1])

			randomCellEdge.EdgeType = Cells.ROAD
			randomCellEdge.Object = &mazeCells[len(mazeCells)-1]

			continue
		}

		checkEdgesAndRemoveFromListIfEmpty(nextCell, nextCellIndex, &notFinishCells, exitFound)

		//if one was found mark both edges as walls
		nextCellEdge := getEdgeByDirection(&nextCell.Edges, shape.GetOpositeDirection(directionVector))
		randomCellEdge.EdgeType = Cells.WALL
		nextCellEdge.EdgeType = Cells.WALL
	}

	return Cells.Maze{
		Height: height,
		Width:  width,
		Cells:  mazeCells,
		Shape:  shape,
	}
}
