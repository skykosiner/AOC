package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

func tokensForItem(ax, ay, bx, by, px, py, offset int) int {
	px, py = px+offset, py+offset

	det := ax*by - ay*bx
	ad := by*px - bx*py
	bd := ax*py - ay*px

	if (3*ad)%det == 0 && bd%det == 0 {
		return (3*ad + bd) / det
	}

	return 0
}

func main() {
	var partOne int
	var partTwo int

	bytes, err := os.ReadFile("./input")
	if err != nil {
		slog.Error("Can't read file", "error", err)
		os.Exit(1)
	}

	for _, game := range strings.Split(string(bytes), "\n\n") {
		nums := utils.ExtractNums(game)
		partOne += tokensForItem(nums[0], nums[1], nums[2], nums[3], nums[4], nums[5], 0)
		partTwo += tokensForItem(nums[0], nums[1], nums[2], nums[3], nums[4], nums[5], 10000000000000)
	}

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
