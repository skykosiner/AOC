package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

type Tile struct {
	X, Y int
}

func (t Tile) areaWithOtherTile(other Tile) int {
	return (int(math.Abs(float64(t.X)-float64(other.X)) + 1)) * (int(math.Abs(float64(t.Y)-float64(other.Y)) + 1))
}

type Section struct {
	A, B Tile
}

func (s Section) overLappingRect(rectA, rectB Tile) bool {
	recMinX := min(rectA.X, rectB.X) + 1
	recMaxX := max(rectA.X, rectB.X) - 1
	recMinY := min(rectA.Y, rectB.Y) + 1
	recMaxY := max(rectA.Y, rectB.Y) - 1

	secMinX := min(s.A.X, s.B.X)
	secMaxX := max(s.A.X, s.B.X)
	secMinY := min(s.A.Y, s.B.Y)
	secMaxY := max(s.A.Y, s.B.Y)

	if secMaxX < recMinX || secMinX > recMaxX {
		return false
	}

	if secMaxY < recMinY || secMinY > recMaxY {
		return false
	}

	return true
}

func getGreenSections(tiles []Tile) []Section {
	var sections []Section

	for i := range len(tiles) - 1 {
		section := Section{A: tiles[i], B: tiles[i+1]}
		sections = append(sections, section)
	}

	sections = append(sections, Section{A: tiles[len(tiles)-1], B: tiles[0]})

	return sections
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

	var redTiles []Tile

	for _, line := range utils.ReadFile(inputFile) {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		redTiles = append(redTiles, Tile{
			X: x,
			Y: y,
		})
	}

	partOne := 0
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			area := redTiles[i].areaWithOtherTile(redTiles[j])
			if area > partOne {
				partOne = area
			}
		}
	}

	fmt.Println("Part One: ", partOne)

	partTwo := 0
	greenSections := getGreenSections(redTiles)

	for i := 0; i < len(redTiles)-1; i++ {
	main:
		for j := i + 1; j < len(redTiles); j++ {
			area := redTiles[i].areaWithOtherTile(redTiles[j])
			if area < partTwo {
				continue
			}

			for _, green := range greenSections {
				if green.overLappingRect(redTiles[i], redTiles[j]) {
					continue main
				}
			}

			partTwo = area
		}
	}

	fmt.Println(partTwo)
}
