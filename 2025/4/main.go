package main

import (
	"flag"
	"fmt"

	"github.com/skykosiner/AOC/pkg/utils"
)

var directions = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func neighbourCount(lines []string, x, y int) bool {
	if lines[x][y] != '@' {
		return false
	}

	count := 0
	for _, dir := range directions {
		xx := x + dir[0]
		yy := y + dir[1]

		if xx >= 0 && yy >= 0 && xx < len(lines) && yy < len(lines[x]) && lines[xx][yy] == '@' {
			count++
		}
	}

	return count < 4
}

func countNeighbours(lines []string, partTwo bool) int {
	count := 0
	for x := range lines {
		for y := 0; y < len(lines[x]); y++ {
			if ok := neighbourCount(lines, x, y); ok {
				if partTwo {
					lines[x] = utils.ReplaceAt(lines[x], y, '.')
				}

				count++
			}
		}
	}

	return count
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

	partOne, partTwo := countNeighbours(lines, false), 0

	number := 1
	for {
		if number == 0 {
			break
		}

		number = countNeighbours(lines, true)
		partTwo += number
	}

	fmt.Println(partOne, partTwo)
}
