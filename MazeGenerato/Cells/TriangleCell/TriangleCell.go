package TriangleCell

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

type TriangleShape struct {
}

func GetShape() TriangleShape {
	return TriangleShape{}
}

func (sqs TriangleShape) GetDistanceBetweenCells() Vector2D {
	return Vector2D{X: 0.5, Y: 1.0}
}

func (sqs TriangleShape) GetEdges() []Cells.MazeEdge {
	edges := make([]Cells.MazeEdge, 3)
	edges[0] = Cells.MazeEdge{Direction: WEST, WallRotation: 1.1071, WallDistance: 0.5, EdgeType: Cells.NONE, Object: nil}
	edges[1] = Cells.MazeEdge{Direction: SOUTH, WallRotation: math.Pi, WallDistance: 1, EdgeType: Cells.NONE, Object: nil}
	edges[2] = Cells.MazeEdge{Direction: EAST, WallRotation: 5.17603, WallDistance: 0.5, EdgeType: Cells.NONE, Object: nil}
	return edges
}

func (sqs TriangleShape) GetOpositeDirection(direction Vector2D) Vector2D {
	if direction.Equals(NORTH) {
		return SOUTH
	}
	return direction
}

func (sqs TriangleShape) GetNewDirection(currentCellDirection Vector2D, edgeDirection Vector2D) Vector2D {
	return currentCellDirection.Mult(-1.0)
}

func (sqs TriangleShape) GetStartingDirection() Vector2D {
	return NORTH
}

func (sqs TriangleShape) DrawShape(img *image.RGBA, position Vector2D, size int, rotation float64, cellColor color.RGBA) {

	rightBorder := Vector2D{X: 2.0, Y: 1.0}.Rotate(rotation)
	leftBorder := Vector2D{X: -2.0, Y: 1.0}.Rotate(rotation)
	bottomBorder := Vector2D{X: 0.0, Y: -1.0}.Rotate(rotation)

	halfSize := float32(size) / 2

	for y := -size / 2; y < size/2; y++ {
		newY := float32(y)
		for x := -size / 2; x < size/2; x++ {
			newX := float32(x)
			if bottomBorder.X*newX+bottomBorder.Y*newY <= halfSize && leftBorder.X*newX+leftBorder.Y*newY <= halfSize && rightBorder.X*newX+rightBorder.Y*newY <= halfSize {
				img.Set(x+int(position.X), y+int(position.Y), cellColor)
			}
		}
	}
}

func (sqs TriangleShape) DrawWall(img *image.RGBA, direction Vector2D, rotation float64, position Vector2D, distance float32, size int, thickness float32, color color.RGBA) {

	halfSize := float32(size / 2)

	topBorder := Vector2D{X: 0.0, Y: 1.0}.Rotate(rotation)
	bottomBorder := Vector2D{X: 0.0, Y: -1.0}.Rotate(rotation)
	leftBorder := Vector2D{X: -1.0, Y: 0.0}.Rotate(rotation).Mult(1 / 1.1)
	rightBorder := Vector2D{X: 1.0, Y: 0.0}.Rotate(rotation).Mult(1 / 1.1)

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
