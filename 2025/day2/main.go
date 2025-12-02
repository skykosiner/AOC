package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

func getNumbersBetween(x, y int) []int {
	var nums []int

	for i := x; i <= y; i++ {
		nums = append(nums, i)
	}

	return nums
}

func padEnd(maxLength int, fillString string) string {
	for len(fillString) < maxLength {
		fillString += fillString
	}

	return fillString
}

func validID(x int) bool {
	check := strconv.Itoa(x)
	halfLength := math.Floor(float64(len(check)) / 2)

	for i := 0; i < int(halfLength); i++ {
		pattern := check[0 : i+1]
		if padEnd(len(check), pattern) == check {
			return false
		}
	}

	return true
}

func main() {
	var ranges [][]int

	for _, item := range strings.Split(utils.ReadFile("./input.test")[0], ",") {
		items := strings.Split(item, "-")
		num1, _ := strconv.Atoi(items[0])
		num2, _ := strconv.Atoi(items[1])

		ranges = append(ranges, []int{num1, num2})
	}

	sum := 0
	for _, r := range ranges {
		for _, id := range getNumbersBetween(r[0], r[1]) {
			if !validID(id) {
				fmt.Println(id)
				sum += id
			}
		}
	}

	fmt.Println(sum)
}
