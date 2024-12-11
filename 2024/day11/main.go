package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

func partOne() int {
	line := utils.ReadFile("./input")[0]
	stones := make(map[int]int)

	for _, num := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(num)
		stones[n]++
	}

	for range 25 {
		cache := make(map[int]int)
		for stone, count := range stones {
			if stone == 0 {
				cache[1] += count
				continue
			} else {
				str := strconv.Itoa(stone)
				if len(str)%2 == 0 {
					mid := len(str) / 2
					first, _ := strconv.Atoi(str[:mid])
					second, _ := strconv.Atoi(str[mid:])
					cache[first] += count
					cache[second] += count
				} else {
					multiplied := stone * 2024
					cache[multiplied] += count
				}
			}
		}

		stones = cache
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
}

func partTwo() int {
	line := utils.ReadFile("./input")[0]
	stones := make(map[int]int)

	for _, num := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(num)
		stones[n]++
	}

	for range 75 {
		cache := make(map[int]int)
		for stone, count := range stones {
			if stone == 0 {
				cache[1] += count
				continue
			} else {
				str := strconv.Itoa(stone)
				if len(str)%2 == 0 {
					mid := len(str) / 2
					first, _ := strconv.Atoi(str[:mid])
					second, _ := strconv.Atoi(str[mid:])
					cache[first] += count
					cache[second] += count
				} else {
					multiplied := stone * 2024
					cache[multiplied] += count
				}
			}
		}

		stones = cache
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
}

func main() {
	fmt.Println("Part One:", partOne())
	fmt.Println("Part Two:", partTwo())
}
