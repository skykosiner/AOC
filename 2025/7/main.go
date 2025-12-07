package main

import (
	"flag"
	"fmt"

	"github.com/skykosiner/AOC/pkg/utils"
)

func laser(grid [][]byte) (int, int) {
	splits := 0
	waysTo := make([]int, len(grid[0]))
	for _, row := range grid {
		fmt.Println(waysTo)
		for j, ch := range row {
			if ch == 'S' {
				waysTo[j] = 1
			} else if ch == '^' && waysTo[j] > 0 {
				splits++
				waysTo[j-1] += waysTo[j]
				waysTo[j+1] += waysTo[j]
				waysTo[j] = 0
			}
		}
	}

	partTwo := 0
	for _, item := range waysTo {
		partTwo += item
	}

	return splits, partTwo
}

func main() {
	test := flag.Bool("test", false, "Run test case or not")
	flag.Parse()

	var inputFile string
	if *test {
		inputFile = "./input.test"
	} else {
		inputFile = "./input"
	}

	lines := utils.ReadFile(inputFile)

	var grid [][]byte
	for y := range lines {
		var row []byte
		for x := range lines[y] {
			row = append(row, lines[y][x])
		}

		grid = append(grid, row)
	}

	fmt.Println(laser(grid))
}
