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

func ReplaceAt(s string, index int, replacement rune) string {
	runes := []rune(s) // Convert string to rune slice
	if index < 0 || index >= len(runes) {
		return s // Return the original string if the index is out of range
	}

	runes[index] = replacement // Replace the character
	return string(runes)       // Convert back to string and return
}
