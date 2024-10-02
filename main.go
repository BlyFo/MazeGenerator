package main

import (
	draw "mazeGenerator/MazeDrawer"
	mazege "mazeGenerator/MazeGenerato"

	//"mazeGenerator/MazeGenerato/Cells/SquareCell"
	"mazeGenerator/MazeGenerato/Cells/TriangleCell"
)

func main() {

	mazeHeight := 9
	mazeWidth := 9
	cellSize := 40
	cellShape := TriangleCell.GetShape()
	showExitPath := false
	saveLocation := ""

	maze := mazege.CreateMaze(mazeHeight, mazeWidth, cellShape)
	draw.DrawMaze(maze, cellSize, showExitPath, saveLocation)
}
