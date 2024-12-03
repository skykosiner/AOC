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
		enabled := true
		regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
		matches := regex.FindAllString(line, -1)

		if matches != nil {
			for _, item := range matches {
				switch item {
				case "don't()":
					enabled = false
				case "do()":
					enabled = true
				default:
					if enabled {
						mul := strings.TrimPrefix(item, "mul(")
						mul = strings.TrimSuffix(mul, ")")
						nums := strings.Split(mul, ",")
						if len(nums) == 2 {
							num1, err := strconv.Atoi(nums[0])
							if err != nil {
								slog.Error("Failed to parse number", "number", nums[0], "error", err)
								continue
							}
							num2, err := strconv.Atoi(nums[1])
							if err != nil {
								slog.Error("Failed to parse number", "number", nums[1], "error", err)
								continue
							}
							numsToTimes = append(numsToTimes, []int{num1, num2})
						}
					}
				}
			}
		}
	}

	// Calculate the sum of the products
	for _, nums := range numsToTimes {
		sum += nums[0] * nums[1]
	}

	fmt.Println(sum)
}
