package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

type Item struct {
	value int
	seen  int
}

func (i Item) String() string {
	return fmt.Sprintf("%d %d\n", i.value, i.seen)
}

func leftRightsides(lines []string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range lines {
		sides := strings.Split(line, "   ")
		if len(sides) == 1 {
			break
		}

		l, _ := strconv.Atoi(sides[0])
		r, _ := strconv.Atoi(sides[1])
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func PartOne() {
	lines := utils.ReadFile("./input")
	left, right := leftRightsides(lines)

	slices.Sort(left)
	slices.Sort(right)

	distances := make([]int, len(left))
	for idx := range left {
		distances[idx] = int(math.Abs(float64(left[idx] - right[idx])))
	}

	fmt.Println(utils.Sum(distances))
}

func PartTwo() {
	var items []Item
	var similarities []int

	lines := utils.ReadFile("./input")
	left, right := leftRightsides(lines)

	for _, l := range left {
		seen := 0
		for _, r := range right {
			if l == r {
				seen++
			}
		}

		items = append(items, Item{
			value: l,
			seen:  seen,
		})
	}

	for _, item := range items {
		similarities = append(similarities, item.seen*item.value)
	}

	fmt.Println(utils.Sum(similarities))
}

func main() {
	PartOne()
	PartTwo()
}
