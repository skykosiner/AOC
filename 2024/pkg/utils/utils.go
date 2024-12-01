package utils

import (
	"os"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(bytes), "\n"), nil
}

func Sum(items []int) int {
	sum := 0
	for _, item := range items {
		sum += item
	}

	return sum
}
