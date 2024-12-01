package dayone

import (
	"fmt"
	"log/slog"
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
	var distances []int

	lines, err := utils.ReadFile("./pkg/day-one/input")
	if err != nil {
		slog.Error("Error reading input file.", "error", err)
		return
	}

	left, right := leftRightsides(lines)

	slices.Sort(left)
	slices.Sort(right)

	for idx := range left {
		l := left[idx]
		r := right[idx]

		if l > r {
			distances = append(distances, int(left[idx]-right[idx]))
		} else {
			distances = append(distances, int(right[idx]-left[idx]))
		}
	}

	fmt.Println(utils.Sum(distances))
}

func PartTwo() {
	var items []Item
	var similarities []int

	lines, err := utils.ReadFile("./pkg/day-one/input")
	if err != nil {
		slog.Error("Error reading input file.", "error", err)
		return
	}

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
