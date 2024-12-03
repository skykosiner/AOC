package main

import (
	"fmt"
	"regexp"

	"github.com/skykosiner/AOC/pkg/utils"
)

func PartOne() {
	var sum int
	lines := utils.ReadFile("./input")

	for _, line := range lines {
		regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		matches := regex.FindAllStringIndex(line, -1)

		for _, match := range matches {
			mul := line[match[0]:match[1]]
			nums := utils.ExtractNums(mul)
			sum += nums[0] * nums[1]
		}
	}

	fmt.Println(sum)
}

func PartTwo() {
	var sum int
	lines := utils.ReadFile("./input")
	enabled := true

	for _, line := range lines {
		regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
		matches := regex.FindAllString(line, -1)

		for _, item := range matches {
			switch item {
			case "don't()":
				enabled = false
			case "do()":
				enabled = true
			default:
				if enabled {
					nums := utils.ExtractNums(item)
					sum += nums[0] * nums[1]
				}
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	PartOne()
	PartTwo()
}
