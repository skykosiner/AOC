package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

func main() {
	var stones []int
	line := utils.ReadFile("./input")[0]

	for _, num := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(num)
		stones = append(stones, n)
	}

	for i := range 75 {
		fmt.Println(i)
		newStones := make([]int, 0, len(stones)*2) // Preallocate for possible splits

		for idx := range stones {
			if stones[idx] == 0 {
				newStones = append(newStones, 1)
				continue
			}

			str := strconv.Itoa(stones[idx])
			if len(str)%2 == 0 {
				midpoint := 1
				for j := 1; j < len(str)/2; j++ {
					midpoint *= 10
				}

				stoneOne := stones[idx] / midpoint
				stoneTwo := stones[idx] % midpoint

				newStones = append(newStones, stoneOne, stoneTwo)

				// Insert the split stones in the slice
				// stones[idx] = stoneOne
				// stones = append(stones[:idx+1], append([]int{stoneTwo}, stones[idx+1:]...)...)
				// idx++ // Skip the newly inserted stoneTwo
			} else {
				newStones = append(newStones, stones[idx]*2024)
			}
		}

		stones = newStones
	}

	fmt.Println(len(stones))
}
