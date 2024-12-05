package main

import (
	"fmt"
	"slices"

	"github.com/skykosiner/AOC/pkg/utils"
)

func isCorrect(page []int, rules [][]int) bool {
	for _, rule := range rules {
		pageA, pageB := rule[0], rule[1]
		positionA, positionB := -1, -1

		// Find positions of A and B in the page
		for i, p := range page {
			if p == pageA {
				positionA = i
			}
			if p == pageB {
				positionB = i
			}
		}

		if positionA != -1 && positionB != -1 && positionA > positionB {
			return false
		}
	}

	return true
}

func main() {
	var rules [][]int
	var pages [][]int
	var sum1 int
	var sum2 int

	lines := utils.ReadFile("./input")

	inRules := true
	for _, line := range lines {
		nums := utils.ExtractNums(line)
		if line == "" {
			inRules = false
			continue
		}

		if inRules {
			rules = append(rules, nums)
		} else {
			pages = append(pages, nums)
		}
	}

	ruleMap := make(map[int][]int)

	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	for _, page := range pages {
		if !isCorrect(page, rules) {
			slices.SortFunc(page, func(a, b int) int {
				if slices.Contains(ruleMap[b], a) {
					return 1
				}

				return -1
			})

			sum2 += page[len(page)/2]
		} else {
			sum1 += page[len(page)/2]
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
