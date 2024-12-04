package main

import (
	"fmt"

	"github.com/skykosiner/AOC/pkg/utils"
)

type Pattern struct {
	value string
	row   int
	col   int
}

func countXmas(lines []string, x, y int) int {
	count := 0

	for xOff := -1; xOff <= 1; xOff++ {
		for yOff := -1; yOff <= 1; yOff++ {
			word := ""

			currX, currY := x, y
			if currX >= 0 && currX < len(lines) && currY >= 0 && currY < len(lines[currX]) {
				word = string(lines[x][y])
			}

			for range 3 {
				currX += xOff
				currY += yOff
				if currX >= 0 && currX < len(lines) && currY >= 0 && currY < len(lines[currX]) {
					word += string(lines[currX][currY])
				}
			}

			if word == "XMAS" {
				count++
			}

		}
	}

	return count
}

func findXMas(x, y int, pattern []Pattern, lines []string) bool {
	for _, p := range pattern {
		newX := p.row + x
		newY := p.col + y

		if newX < 0 || newX >= len(lines) || newY < 0 || newY >= len(lines[newX]) {
			return false
		}

		if string(lines[newX][newY]) != p.value {
			return false
		}
	}
	return true
}

func main() {
	// Part One
	countPartOne := 0
	lines := utils.ReadFile("./input")

	for x := range lines {
		for y := range lines[x] {
			if lines[x][y] == 'X' {
				countPartOne += countXmas(lines, x, y)
			}
		}
	}

	fmt.Println(countPartOne)

	patterns := [][]Pattern{
		{
			{"M", -1, -1},
			{"S", -1, 1},
			{"A", 0, 0},
			{"M", 1, -1},
			{"S", 1, 1},
		},
		{
			{"S", -1, -1},
			{"M", 1, -1},
			{"A", 0, 0},
			{"S", -1, 1},
			{"M", 1, 1},
		},
		{
			{"S", 1, -1},
			{"M", 1, 1},
			{"A", 0, 0},
			{"S", -1, -1},
			{"M", -1, 1},
		},
		{
			{"M", -1, 1},
			{"S", 1, 1},
			{"A", 0, 0},
			{"M", -1, -1},
			{"S", 1, -1},
		},
	}

	// Part Two
	countPartTwo := 0
	for x := range lines {
		for y := range lines[x] {
			for _, pattern := range patterns {
				if findXMas(x, y, pattern, lines) {
					countPartTwo += 1
				}
			}
		}
	}

	fmt.Println(countPartTwo)
}
