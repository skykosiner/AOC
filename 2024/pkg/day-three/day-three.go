package daythree

import (
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

func PartOne() {
	var sum int
	var numsToTimes [][]int
	lines, err := utils.ReadFile("./pkg/day-three/input")
	if err != nil {
		slog.Error("Error reading input file.", "error", err)
		return
	}

	for _, line := range lines {
		regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		matches := regex.FindAllStringIndex(line, -1)

		if matches != nil {
			for _, match := range matches {
				mul := line[match[0]:match[1]]
				mul = strings.TrimSuffix(strings.TrimPrefix(mul, "mul("), ")")
				nums := strings.Split(mul, ",")
				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])
				numsToTimes = append(numsToTimes, []int{num1, num2})
			}
		}
	}

	for _, nums := range numsToTimes {
		sum += nums[0] * nums[1]
	}

	fmt.Println(sum)
}

func PartTwo() {
	var sum int
	var numsToTimes [][]int
	lines, err := utils.ReadFile("./pkg/day-three/input")
	if err != nil {
		slog.Error("Error reading input file.", "error", err)
		return
	}

	for _, line := range lines {
		seen := make([]string, 0)
		enabled := true
		for _, char := range strings.Split(line, "") {
		}
	}

	for _, nums := range numsToTimes {
		sum += nums[0] * nums[1]
	}

	fmt.Println(sum)
}
