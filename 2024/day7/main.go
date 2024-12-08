package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

type EQ struct {
	total int
	nums  []int
}

func isEquationGood(eq EQ, part2 bool) bool {
	var maxBinary int
	if part2 {
		maxBinary = int(math.Pow(3, float64(len(eq.nums))))
	} else {
		maxBinary = int(math.Pow(2, float64(len(eq.nums))))
	}

	for i := range maxBinary {
		var binaryStr string
		currTotal := eq.nums[0]
		if part2 {
			ternaryStr := strconv.FormatInt(int64(i), 3)
			binaryStr = strings.Repeat("0", len(eq.nums)-len(ternaryStr)) + ternaryStr
		} else {
			binaryStr = fmt.Sprintf("%0*b", len(eq.nums), i)
		}

		for idx, b := range binaryStr {
			if idx == 0 {
				continue
			}

			if idx >= len(eq.nums) {
				break
			}

			if b == '0' {
				currTotal *= eq.nums[idx]
			} else if b == '1' {
				currTotal += eq.nums[idx]
			} else {
				resultStr := strconv.Itoa(currTotal) + strconv.Itoa(eq.nums[idx])
				result, _ := strconv.Atoi(resultStr)
				currTotal = result
			}

			if currTotal > eq.total {
				break
			}
		}

		if currTotal == eq.total {
			return true
		}
	}

	return false
}

func main() {
	var Equations []EQ
	sum1 := 0
	sum2 := 0
	lines := utils.ReadFile("./input")

	for _, line := range lines {
		if line == "" {
			continue
		}

		nums := utils.ExtractNums(line)
		total := nums[0]
		nums = nums[1:]

		Equations = append(Equations, EQ{
			total,
			nums,
		})
	}

	for _, eq := range Equations {
		if isEquationGood(eq, false) {
			sum1 += eq.total
		}
	}

	for _, eq := range Equations {
		if isEquationGood(eq, true) {
			sum2 += eq.total
		}
	}

	fmt.Println("Part One:", sum1)
	fmt.Println("Part Two:", sum2)
}
