package utils

import (
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		slog.Error("Cloudn't read file provided.", "path", path, "error", err)
		os.Exit(1)
	}

	return strings.Split(string(bytes), "\n")
}

func ExtractNums(line string) []int {
	var nums []int

	regex := regexp.MustCompile(`\d+`)
	matches := regex.FindAllString(line, -1)

	for _, match := range matches {
		number, err := strconv.Atoi(match)
		if err != nil {
			slog.Error("Couldn't convert string to number.", "match", match, "error", err)
			os.Exit(1)
		}

		nums = append(nums, number)
	}

	return nums
}

func Sum(items []int) int {
	sum := 0
	for _, item := range items {
		sum += item
	}

	return sum
}
