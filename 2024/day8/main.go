package main

import (
	"fmt"

	"github.com/skykosiner/AOC/pkg/utils"
)

func inBounds(x, y, length int) bool {
	return x >= 0 && x < length && y >= 0 && y < length
}

func getAntinodes(a, b [2]int, grid []string) [][2]int {
	var antinodes [][2]int
	ax, ay := a[0], a[1]
	bx, by := b[0], b[1]

	cx, cy := ax-(bx-ax), ay-(by-ay)
	dx, dy := bx+(bx-ax), by+(by-ay)

	// Check if within bounds and add to result
	if inBounds(cx, cy, len(grid)) {
		antinodes = append(antinodes, [2]int{cx, cy})
	}
	if inBounds(dx, dy, len(grid)) {
		antinodes = append(antinodes, [2]int{dx, dy})
	}

	return antinodes
}

func partTwo(a, b [2]int, grid []string) [][2]int {
	var antinodes [][2]int
	ax, ay := a[0], a[1]
	bx, by := b[0], b[1]

	dx, dy := bx-ax, by-ay

	i := 0
	for inBounds(ax-dx*i, ay-dy*i, len(grid)) {
		antinodes = append(antinodes, [2]int{ax - dx*i, ay - dy*i})
		i++
	}

	i = 0
	for inBounds(bx+dx*i, by+dy*i, len(grid)) {
		antinodes = append(antinodes, [2]int{bx + dx*i, by + dy*i})
		i++
	}

	return antinodes
}

func main() {
	var cleanGrid []string
	grid := utils.ReadFile("./input")

	for _, line := range grid {
		if line != "" {
			cleanGrid = append(cleanGrid, line)
		}
	}

	grid = cleanGrid

	allLocs := make(map[rune][][2]int)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] != '.' {
				allLocs[rune(grid[x][y])] = append(allLocs[rune(grid[x][y])], [2]int{x, y})
			}
		}
	}

	antinodesSet := make(map[[2]int]struct{})
	antinodesSetTwo := make(map[[2]int]struct{})

	for _, locs := range allLocs {
		for i := 0; i < len(locs); i++ {
			for j := i + 1; j < len(locs); j++ {
				antinodes := getAntinodes(locs[i], locs[j], grid)
				for _, antinode := range antinodes {
					antinodesSet[antinode] = struct{}{}
				}

				antinodesTwo := partTwo(locs[i], locs[j], grid)
				for _, antinode := range antinodesTwo {
					antinodesSetTwo[antinode] = struct{}{}
				}
			}
		}
	}

	fmt.Println("Part One:", len(antinodesSet))
	fmt.Println("Part Two:", len(antinodesSetTwo))
}
