package utils

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X float32
	Y float32
}

func (b Vector2D) String() string { return fmt.Sprintf("( x: %f, y: %f )\n", b.X, b.Y) }

func (b Vector2D) Add(vectorToAdd Vector2D) Vector2D {
	return Vector2D{b.X + vectorToAdd.X, b.Y + vectorToAdd.Y}
}

func (b Vector2D) AddC(constToAdd float32) Vector2D {
	return Vector2D{b.X + constToAdd, b.Y + constToAdd}
}

func (b Vector2D) Substract(vectorTosubstract Vector2D) Vector2D {
	return Vector2D{b.X - vectorTosubstract.X, b.Y - vectorTosubstract.Y}
}

func (b Vector2D) Mult(constant float32) Vector2D {
	return Vector2D{b.X * constant, b.Y * constant}
}

func (b Vector2D) Distance(vector Vector2D) float32 {
	distanceX := math.Pow(float64(vector.X-b.X), 2)
	distancey := math.Pow(float64(vector.Y-b.Y), 2)
	return float32(math.Sqrt(distanceX + distancey))
}

func (b Vector2D) Rotate(rotation float64) Vector2D {
	newX := float64(b.X)*math.Cos(rotation) - float64(b.Y)*math.Sin(rotation)
	newY := float64(b.X)*math.Sin(rotation) + float64(b.Y)*math.Cos(rotation)
	return Vector2D{X: float32(newX), Y: float32(newY)}
}

func (b Vector2D) DotProduct(vector Vector2D) float32 {
	return b.X*vector.X + b.Y*vector.Y
}

func (b Vector2D) Equals(vector Vector2D) bool {
	return b.X == vector.X && b.Y == vector.Y
}
