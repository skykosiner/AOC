package main

import (
	"fmt"
	"regexp"

	"github.com/skykosiner/AOC/pkg/utils"
)

func main() {
	var sum1 int
	var sum2 int

	lines := utils.ReadFile("./input")
	regex1 := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	regex2 := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	enabled := true

	for _, line := range lines {
		matchesOne := regex1.FindAllStringIndex(line, -1)

		for _, match := range matchesOne {
			mul := line[match[0]:match[1]]
			nums := utils.ExtractNums(mul)
			sum1 += nums[0] * nums[1]
		}

		matchesTwo := regex2.FindAllString(line, -1)

		for _, item := range matchesTwo {
			switch item {
			case "don't()":
				enabled = false
			case "do()":
				enabled = true
			default:
				if enabled {
					nums := utils.ExtractNums(item)
					sum2 += nums[0] * nums[1]
				}
			}
		}
	}

	fmt.Printf("%d\n%d\n", sum1, sum2)
}
