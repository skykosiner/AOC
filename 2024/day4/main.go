package main

import (
	"fmt"

	"github.com/skykosiner/AOC/pkg/utils"
)

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

func main() {
	count := 0
	lines := utils.ReadFile("./input")

	for x := range lines {
		for y := range lines[x] {
			if lines[x][y] == 'X' {
				count += countXmas(lines, x, y)
			}
		}
	}

	fmt.Println(count)
}
