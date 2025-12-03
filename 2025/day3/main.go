package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

func main() {
	var batteries [][]int
	for _, line := range utils.ReadFile("./input") {
		var nums []int
		numsStr := strings.Split(line, "")
		if len(numsStr) == 0 {
			break
		}
		for _, numStr := range numsStr {
			number, _ := strconv.Atoi(numStr)
			nums = append(nums, number)
		}

		batteries = append(batteries, nums)
	}

	partOne := 0

	for _, battery := range batteries {
		maxNum := slices.Max(battery[:len(battery)-1])
		maxIndex := slices.Index(battery, maxNum)
		partOne += maxNum*10 + slices.Max(battery[maxIndex+1:])
	}

	partTwo := 0

	for _, bank := range utils.ReadFile("./input") {
		if bank == "" {
			break
		}

		s, p := "", 0
		for b := 12; b > 0; b-- {
			for i := p; i <= len(bank)-b; i++ {
				if bank[i] > bank[p] {
					p = i
				}
			}

			s += string(bank[p])
			p++
		}

		num, _ := strconv.Atoi(s)
		partTwo += num
	}

	fmt.Println(partOne)
	fmt.Println(partTwo)
}
