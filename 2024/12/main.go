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

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func containsCord(region [][]int, r, c float64) bool {
	for _, point := range region {
		if point[0] == int(r) && point[1] == int(c) {
			return true
		}
	}
	return false
}

type Plot struct {
	letter rune
	cords  [][]int
}

func (p *Plot) getArea() int {
	return len(p.cords)
}

func (p *Plot) getPerimeter() int {
	if len(p.cords) == 1 {
		return 4
	}

	totalSides := len(p.cords) * 4
	sharedSides := 0

	for i := 0; i < len(p.cords); i++ {
		for j := i + 1; j < len(p.cords); j++ {
			if (p.cords[i][0] == p.cords[j][0] && abs(p.cords[i][1]-p.cords[j][1]) == 1) ||
				(p.cords[i][1] == p.cords[j][1] && abs(p.cords[i][0]-p.cords[j][0]) == 1) {
				sharedSides++
			}
		}
	}

	return totalSides - (sharedSides * 2)
}

func (p *Plot) getPricePartOne() int {
	return p.getArea() * p.getPerimeter()
}

func (p *Plot) getPricePartTwo() int {
	cornerCandidates := make(map[[2]float64]bool)

	for _, p := range p.cords {
		r, c := p[0], p[1]

		cornerCandidates[[2]float64{float64(r) - 0.5, float64(c) - 0.5}] = true
		cornerCandidates[[2]float64{float64(r) + 0.5, float64(c) - 0.5}] = true
		cornerCandidates[[2]float64{float64(r) + 0.5, float64(c) + 0.5}] = true
		cornerCandidates[[2]float64{float64(r) - 0.5, float64(c) + 0.5}] = true
	}

	corners := 0

	for corner := range cornerCandidates {
		cr, cc := corner[0], corner[1]
		config := []bool{
			containsCord(p.cords, cr-0.5, cc-0.5),
			containsCord(p.cords, cr+0.5, cc-0.5),
			containsCord(p.cords, cr+0.5, cc+0.5),
			containsCord(p.cords, cr-0.5, cc+0.5),
		}

		number := 0
		for _, val := range config {
			if val {
				number++
			}
		}

		if number == 1 {
			corners++
		} else if number == 2 {
			if (config[0] && !config[1] && config[2] && !config[3]) || (!config[0] && config[1] && !config[2] && config[3]) {
				corners += 2
			}
		} else if number == 3 {
			corners++
		}
	}

	return p.getArea() * corners
}

func floodFill(lines []string, x, y int, visited map[[2]int]bool, letter rune) [][]int {
	stack := [][2]int{{x, y}}
	region := [][]int{}

	for len(stack) > 0 {
		// Pop the top element from the stack
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		cx, cy := curr[0], curr[1]
		if visited[[2]int{cx, cy}] {
			continue
		}
		visited[[2]int{cx, cy}] = true
		region = append(region, []int{cx, cy})

		for _, dir := range directions {
			nx, ny := cx+dir[0], cy+dir[1]
			if nx >= 0 && nx < len(lines) && ny >= 0 && ny < len(lines[nx]) &&
				!visited[[2]int{nx, ny}] && rune(lines[nx][ny]) == letter {
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}

	return region
}

func main() {
	var lines []string
	for _, l := range utils.ReadFile("./input") {
		if l != "" {
			lines = append(lines, l)
		}
	}

	plots := make([]Plot, 0)
	visited := make(map[[2]int]bool)

	for x := range lines {
		for y := range lines[x] {

			if visited[[2]int{x, y}] {
				continue
			}

			letter := rune(lines[x][y])
			region := floodFill(lines, x, y, visited, letter)
			plots = append(plots, Plot{
				letter: letter,
				cords:  region,
			})
		}
	}

	partOne := 0
	partTwo := 0

	for _, plot := range plots {
		partOne += plot.getPricePartOne()
		partTwo += plot.getPricePartTwo()
	}

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
