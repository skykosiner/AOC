package main

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/skykosiner/AOC/pkg/utils"
)

func correct(blocks []string) bool {
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == "." {
			for j := i + 1; j < len(blocks); j++ {
				if blocks[j] != "." {
					return false
				}
			}
			return true
		}
	}

	return true
}

func partOne(blocks []string) {
	for {
		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i] != "." {
				block := blocks[i]

				for j := 0; j < len(blocks); j++ {
					if blocks[j] == "." {
						blocks[j] = block
						blocks[i] = "."
						break
					}
				}

				break
			}
		}

		if correct(blocks) {
			break
		}
	}
}

func findFreePoints(blocks []string) [][]int {
	points := make([][]int, 0)

	length := 0
	start := -1

	for i, block := range blocks {
		if block == "." {
			if start == -1 {
				start = i
			}

			length++
		} else {
			if length > 0 {
				points = append(points, []int{start, length})
			}
			start = -1
			length = 0
		}
	}
	// Add last point if it exists
	if length > 0 {
		points = append(points, []int{start, length})
	}

	return points
}

func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func partTwo(blocks []string) {
	files := make(map[int]int)

	for _, block := range blocks {
		if block != "." {
			num, _ := strconv.Atoi(block)
			files[num] += 1
		}
	}

	var ids []int

	for key := range files {
		ids = append(ids, key)
	}

	slices.Sort(ids)
	reverse(ids)

	for _, id := range ids {
		points := findFreePoints(blocks)
		for _, point := range points {
			start, length := point[0], point[1]
			// Check if the free space is large enough
			if length >= files[id] {
				// Move the file to the free space
				for i := 0; i < files[id]; i++ {
					blocks[start+i] = strconv.Itoa(id)
				}

				// Clear the original file blocks
				c := 0
				for i := len(blocks) - 1; i >= 0; i-- {
					if blocks[i] == strconv.Itoa(id) {
						blocks[i] = "."
						c++
						if c == files[id] {
							break
						}
					}
				}
				break
			}
		}
	}
}

func main() {
	t := time.Now()

	var blocksOne []string
	var blocksTwo []string

	sum1 := 0
	sum2 := 0
	line := utils.ReadFile("./input")[0]

	currID := 0
	for idx, char := range line {
		if idx%2 == 0 {
			for range int(char - '0') {
				blocksOne = append(blocksOne, strconv.Itoa(currID))
				blocksTwo = append(blocksTwo, strconv.Itoa(currID))
			}

			currID++
		} else {
			for range int(char - '0') {
				blocksOne = append(blocksOne, ".")
				blocksTwo = append(blocksTwo, ".")
			}
		}
	}

	partOne(blocksOne)
	partTwo(blocksTwo)

	for idx, block := range blocksOne {
		num, err := strconv.Atoi(block)
		if err != nil {
			continue
		}

		sum1 += idx * num
	}

	fmt.Println(blocksTwo)

	for i, block := range blocksTwo {
		if block != "." {
			num, _ := strconv.Atoi(block)
			sum2 += i * num
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
	since := time.Since(t)
	fmt.Println(since.Minutes(), since.Seconds())
}
