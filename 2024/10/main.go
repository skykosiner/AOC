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

func walk(x, y int, seen map[[2]int]bool, lines []string, partTwo bool) int {
	// Off the map
	if x < 0 || x >= len(lines) || y < 0 || y >= len(lines[x]) {
		return 0
	}

	// If the cell is already seen, stop exploring only if it's part one
	if !partTwo && seen[[2]int{x, y}] {
		return 0
	}

	seen[[2]int{x, y}] = true

	// If we've reached a '9', count it
	if lines[x][y]-'0' == 9 {
		return 1
	}

	// Explore all directions and accumulate counts
	count := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		// Only proceed if the next number is one greater
		if nx >= 0 && nx < len(lines) && ny >= 0 && ny < len(lines[nx]) &&
			lines[nx][ny]-'0' == lines[x][y]-'0'+1 {
			count += walk(nx, ny, seen, lines, partTwo)
		}
	}

	return count
}

func main() {
	var lines []string
	for _, line := range utils.ReadFile("./input") {
		if line != "" {
			lines = append(lines, line)
		}
	}

	trailHeads := make([][]int, 0)

	for x := range lines {
		for y := range lines[x] {
			if lines[x][y]-'0' == 0 {
				trailHeads = append(trailHeads, []int{x, y})
			}
		}
	}

	count, paths := 0, 0

	for _, head := range trailHeads {
		seen := make(map[[2]int]bool)
		count += walk(head[0], head[1], seen, lines, false)
		paths += walk(head[0], head[1], seen, lines, true)
	}

	fmt.Println("Part One:", count)
	fmt.Println("Part Two:", paths)
}
