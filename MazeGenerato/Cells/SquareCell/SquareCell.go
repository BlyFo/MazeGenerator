package SquareCell

import (
	"image"
	"image/color"
	"math"
	"mazeGenerator/MazeGenerato/Cells"
	Utils "mazeGenerator/Utils/2DVector"
)

type Vector2D = Utils.Vector2D

var (
	NORTH Vector2D = Vector2D{X: 0.0, Y: 1.0}
	WEST  Vector2D = Vector2D{X: -1.0, Y: 0.0}
	SOUTH Vector2D = Vector2D{X: 0.0, Y: -1.0}
	EAST  Vector2D = Vector2D{X: 1.0, Y: 0.0}
)

type squareShape struct {
}

func GetShape() squareShape {
	return squareShape{}
}

func (sqs squareShape) GetDistanceBetweenCells() Vector2D {
	return Vector2D{X: 1.0, Y: 1.0}
}

func (sqs squareShape) GetEdges() []Cells.MazeEdge {
	edges := make([]Cells.MazeEdge, 4)
	edges[0] = Cells.MazeEdge{Direction: NORTH, WallRotation: 0.0, WallDistance: 1, EdgeType: Cells.NONE, Object: nil}
	edges[1] = Cells.MazeEdge{Direction: WEST, WallRotation: 3 * math.Pi / 2, WallDistance: 1, EdgeType: Cells.NONE, Object: nil}
	edges[2] = Cells.MazeEdge{Direction: SOUTH, WallRotation: math.Pi, WallDistance: 1, EdgeType: Cells.NONE, Object: nil}
	edges[3] = Cells.MazeEdge{Direction: EAST, WallRotation: math.Pi / 2, WallDistance: 1, EdgeType: Cells.NONE, Object: nil}
	return edges
}

func (sqs squareShape) GetOpositeDirection(direction Vector2D) Vector2D {
	return direction.Mult(-1)
}
func (sqs squareShape) GetNewDirection(currentCellDirection Vector2D, edgeDirection Vector2D) Vector2D {
	return NORTH
}

func (sqs squareShape) GetStartingDirection() Vector2D {
	return NORTH
}

func (sqs squareShape) DrawShape(img *image.RGBA, position Vector2D, size int, rotation float64, cellColor color.RGBA) {

	topBorder := Vector2D{X: 0.0, Y: 1.0}.Rotate(rotation)
	bottomBorder := Vector2D{X: 0.0, Y: -1.0}.Rotate(rotation)
	leftBorder := Vector2D{X: -1.0, Y: 0.0}.Rotate(rotation)
	rightBorder := Vector2D{X: 1.0, Y: 0.0}.Rotate(rotation)

	halfSize := float32(size) / 2

	for y := -size / 2; y < size/2; y++ {
		newY := float32(y)
		for x := -size / 2; x < size/2; x++ {
			newX := float32(x)
			if topBorder.X*newX+topBorder.Y*newY <= halfSize && bottomBorder.X*newX+bottomBorder.Y*newY <= halfSize && leftBorder.X*newX+leftBorder.Y*newY <= halfSize && rightBorder.X*newX+rightBorder.Y*newY <= halfSize {
				img.Set(x+int(position.X), y+int(position.Y), cellColor)
			}
		}
	}
}

func (sqs squareShape) DrawWall(img *image.RGBA, direction Vector2D, rotation float64, position Vector2D, distance float32, size int, thickness float32, color color.RGBA) {

	halfSize := float32(size / 2)

	topBorder := Vector2D{X: 0.0, Y: 1.0}.Rotate(rotation)
	bottomBorder := Vector2D{X: 0.0, Y: -1.0}.Rotate(rotation)
	leftBorder := Vector2D{X: -1.0, Y: 0.0}.Rotate(rotation)
	rightBorder := Vector2D{X: 1.0, Y: 0.0}.Rotate(rotation)

	wallPosition := position.Add(direction.Mult(halfSize).Mult(distance))

	for y := -size / 2; y < size/2; y++ {
		newY := float32(y)
		for x := -size / 2; x < size/2; x++ {
			newX := float32(x)
			if topBorder.X*newX+topBorder.Y*newY <= thickness && bottomBorder.X*newX+bottomBorder.Y*newY <= thickness && leftBorder.X*newX+leftBorder.Y*newY <= halfSize && rightBorder.X*newX+rightBorder.Y*newY <= halfSize {
				img.Set(x+int(wallPosition.X), y+int(wallPosition.Y), color)
			}
		}
	}
}
