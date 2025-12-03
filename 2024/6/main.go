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

func walk(x, y, dir int, seen [][]bool, lines []string, visited *int) bool {
	nextRow, nextCol := x+directions[dir][0], y+directions[dir][1]

	if nextRow >= 0 && nextRow < len(lines) && nextCol >= 0 && nextCol < len(lines[nextRow]) && lines[nextRow][nextCol] == '#' {
		newDir := (dir + 1) % 4
		return walk(x, y, newDir, seen, lines, visited)
	}

	// Of the map
	if nextRow < 0 || nextRow >= len(lines) || nextCol < 0 || nextCol >= len(lines[nextRow]) {
		return true
	}

	// Seen before
	if !seen[nextRow][nextCol] {
		seen[nextRow][nextCol] = true
		*visited++
	}

	return walk(nextRow, nextCol, dir, seen, lines, visited)
}

func main() {
	var startX, startY int

	lines := utils.ReadFile("./input.test")
	seen := make([][]bool, len(lines))
	for i := range seen {
		seen[i] = make([]bool, len(lines[i]))
	}

	for x := range lines {
		for y := range lines[x] {
			if lines[x][y] == '^' {
				startX, startY = x, y
			}
		}
	}

	visited := 1
	seen[startX][startY] = true

	if walk(startX, startY, 0, seen, lines, &visited) {
		fmt.Println("Part One:", visited)
	} else {
		fmt.Println("Can't solve??")
	}
}
