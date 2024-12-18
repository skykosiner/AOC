package main

import (
	"fmt"

	"github.com/skykosiner/AOC/pkg/utils"
)

var directions = [][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

type Point struct {
	x, y, dir int
}

func walk(point Point, maze []string, seen [][]bool, points *int, path []Point) bool {
	if (point.x < 0 || point.x >= len(maze)) || (point.y < 0 || point.y >= len(maze[point.x])) {
		return false
	}

	if maze[x][y] == '#' {
		return false
	}

	if maze[x][y] == 'E' {
		return true
	}

	// If it's a valid tile, mark it as visited
	if !seen[x][y] {
		seen[x][y] = true
		// Increment points for moving forward
		*points += 1
	}

	return false
}

func main() {
	var maze []string

	for _, line := range utils.ReadFile("./input.test") {
		if line != "" {
			maze = append(maze, line)
		}
	}

	startX, startY := 0, 0

	found := false
	for x := range maze {
		for y := range maze[x] {
			if maze[x][y] == 'S' {
				startX, startY = x, y
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[i]))
	}

	points := 0
	walk(startX, startY, maze, seen, &points)
	fmt.Println(points)
}
