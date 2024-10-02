package Cells

import (
	"image"
	"image/color"
	Utils "mazeGenerator/Utils/2DVector"
)

type edgeType int

const (
	NONE  edgeType = 0
	WALL  edgeType = 1
	ROAD  edgeType = 2
	ENTRY edgeType = 3
	EXIT  edgeType = 4
)

type MazeEdge struct {
	Direction    Utils.Vector2D
	WallRotation float64
	WallDistance float32
	EdgeType     edgeType
	Object       *MazeCell
}

type MazeCell struct {
	Position   Utils.Vector2D
	Direction  Utils.Vector2D
	Rotation   float64
	Edges      []MazeEdge
	IsExitPath bool
}

type Maze struct {
	Height int
	Width  int
	Cells  []MazeCell
	Shape  Cell
}

type Cell interface {
	GetOpositeDirection(d Utils.Vector2D) Utils.Vector2D
	GetDistanceBetweenCells() Utils.Vector2D
	DrawShape(img *image.RGBA, position Utils.Vector2D, size int, rotation float64, cellColor color.RGBA)
	DrawWall(img *image.RGBA, direction Utils.Vector2D, rotation float64, cellPosition Utils.Vector2D, distance float32, size int, thickness float32, color color.RGBA)
	GetNewDirection(currentCellDirection Utils.Vector2D, edgeDirection Utils.Vector2D) Utils.Vector2D
	GetStartingDirection() Utils.Vector2D
	GetEdges() []MazeEdge
}
