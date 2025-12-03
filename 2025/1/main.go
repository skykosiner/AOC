package main

import (
	"fmt"

	"github.com/skykosiner/AOC/pkg/utils"
)

func partTwo(lines []string) int {
	dial := 50
	count := 0

	for _, line := range lines {
		nums := utils.ExtractNums(line)
		if len(nums) == 0 {
			break
		}

		for i := 0; i < nums[0]; i++ {
			if line[0] == 'L' {
				dial--
			} else {
				dial++
			}

			dial = (dial + 100) % 100

			if dial == 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	start := 50
	seen := 0
	lines := utils.ReadFile("./input")

	for _, line := range lines {
		nums := utils.ExtractNums(line)
		if len(nums) == 0 {
			break
		}

		moveAmount := nums[0] % 100

		if line[0] == 'L' {
			start = (start + moveAmount%100 + 100) % 100
		}

		if line[0] == 'R' {
			start = (start - moveAmount%100 + 100) % 100
		}

		if start == 0 {
			seen += 1
		}
	}

	fmt.Println(seen)
	fmt.Println(partTwo(lines))
}
