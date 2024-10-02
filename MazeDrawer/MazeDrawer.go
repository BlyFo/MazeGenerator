package mazege

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	Cells "mazeGenerator/MazeGenerato/Cells"
	Utils "mazeGenerator/Utils/2DVector"
	"os"
	"sync"
)

type Vector2D = Utils.Vector2D
type MazeCell = Cells.MazeCell
type Maze = Cells.Maze

var wg sync.WaitGroup

func drawSlice(img *image.RGBA, maze []MazeCell, shape Cells.Cell, cellSize int, offSet float32, exitPathColor color.RGBA, showExitPath bool) {
	defer wg.Done()

	sizeFloat := float32(cellSize)
	wallThickness := sizeFloat * 0.025

	cellColor := color.RGBA{120, 120, 120, 255}
	wallColor := color.RGBA{0, 0, 0, 255}
	startEndWallColor := color.RGBA{255, 0, 0, 255}

	for i := 0; i < len(maze); i++ {

		center := maze[i].Position.Mult(sizeFloat).AddC(offSet)

		if showExitPath && maze[i].IsExitPath {
			shape.DrawShape(img, center, cellSize, maze[i].Rotation, exitPathColor)
		} else {
			shape.DrawShape(img, center, cellSize, maze[i].Rotation, cellColor)
		}

		for j := 0; j < len(maze[i].Edges); j++ {

			edge := maze[i].Edges[j]
			wallDirection := edge.Direction.Rotate(maze[i].Rotation)
			wallRotation := edge.WallRotation + maze[i].Rotation
			wallDistance := edge.WallDistance

			if edge.EdgeType == Cells.WALL {
				shape.DrawWall(img, wallDirection, wallRotation, center, wallDistance, cellSize, wallThickness, wallColor)
				continue
			}
			if edge.EdgeType == Cells.ENTRY || edge.EdgeType == Cells.EXIT {
				shape.DrawWall(img, wallDirection, wallRotation, center, wallDistance, cellSize, wallThickness, startEndWallColor)
			}
		}
	}
}

func saveImage(image *image.RGBA, saveLocation string) {
	myfile, err := os.Create(saveLocation)
	if err != nil {
		panic(err)
	}
	defer myfile.Close()
	png.Encode(myfile, image)
}

func DrawMaze(maze Maze, cellSize int, showExitPath bool, saveLocation string) {

	maxBatchSize := 1000
	cellsList := maze.Cells
	sizeFloat := float32(cellSize)
	offSet := int(sizeFloat/2) + 10

	backgroundColor := color.RGBA{255, 255, 255, 255}
	exitPathColor := color.RGBA{60, 60, 60, 255}

	mazeImage := image.NewRGBA(image.Rect(0, 0, (maze.Width+1)*cellSize+offSet, (maze.Height+1)*cellSize+offSet)) //background
	draw.Draw(mazeImage, mazeImage.Bounds(), &image.Uniform{backgroundColor}, image.ZP, draw.Src)

	batches := int(math.Ceil(float64(len(cellsList)) / float64(maxBatchSize)))

	for i := 0; i < batches; i++ {
		wg.Add(1)

		//in case there are less than maxBatchSize in a batch
		sliceEnd := min(maxBatchSize, len(cellsList)-i*maxBatchSize)
		mazeSlice := cellsList[i*maxBatchSize : i*maxBatchSize+sliceEnd]

		go drawSlice(mazeImage, mazeSlice, maze.Shape, cellSize, float32(offSet), exitPathColor, showExitPath)
	}

	wg.Wait()

	saveImage(mazeImage, saveLocation)
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
