package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var grid [][]byte

	bytes, err := os.ReadFile("./input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(bytes), "\n\n")

	for _, line := range strings.Split(lines[0], "\n") {
		grid = append(grid, []byte(line))
	}

	moves := strings.Split(strings.TrimSpace(lines[1]), "")

	var startX, startY int
	found := false
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '@' {
				startX, startY = x, y
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	directions := map[string][2]int{
		"^": {-1, 0},
		"v": {1, 0},
		"<": {0, -1},
		">": {0, 1},
	}

	for _, move := range moves {
		dir, exists := directions[move]
		if !exists {
			continue
		}

		dr, dc := dir[0], dir[1]
		newX, newY := startX+dr, startY+dc

		if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) || grid[newX][newY] == '#' {
			continue // Skip invalid move
		}

		if grid[newX][newY] == '.' {
			grid[startX][startY] = '.'
			grid[newX][newY] = '@'
			startX, startY = newX, newY
		} else if grid[newX][newY] == 'O' {
			if canMoveBoxes(grid, newX, newY, dr, dc) {
				moveBoxes(grid, newX, newY, dr, dc)
				grid[startX][startY] = '.'
				grid[newX][newY] = '@'
				startX, startY = newX, newY
			}
		}
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	sum := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'O' {
				sum += 100*x + y
			}
		}
	}

	fmt.Println(sum)
}

func canMoveBoxes(grid [][]byte, x, y, dr, dc int) bool {
	for x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
		if grid[x][y] == '.' {
			return true
		} else if grid[x][y] == '#' {
			return false
		}
		x += dr
		y += dc
	}
	return false
}

func moveBoxes(grid [][]byte, x, y, dr, dc int) {
	var positions [][]int

	for x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
		if grid[x][y] == 'O' {
			positions = append(positions, []int{x, y})
		} else if grid[x][y] == '.' {
			break
		}
		x += dr
		y += dc
	}

	for i := len(positions) - 1; i >= 0; i-- {
		px, py := positions[i][0], positions[i][1]
		grid[px+dr][py+dc] = 'O'
		grid[px][py] = '.'
	}
}
