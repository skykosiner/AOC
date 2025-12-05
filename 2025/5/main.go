package main

import (
	"flag"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

func merge(fresh [][]int) []int {
	var mergedArr []int

	for _, f := range fresh {
		slices.Sort(f)
	}

	for i := 1; i < len(fresh); i++{
		if (fresh[i][0] >= fresh[i-1][0]) && (fresh[i][0] <= fresh[i-1][1]) {
			fmt.Println(math.Max(float64(fresh[i-1][1]), float64(fresh[i][1])))
		}
	}

	return mergedArr
}

func main() {
	test := flag.Bool("test", false, "Run test case or not")
	flag.Parse()

	var inputFile string
	if *test {
		inputFile = "./input.test"
	} else {
		inputFile = "./input"
	}

	lines := utils.ReadFile(inputFile)

	var fresh [][]int
	var IDs []int

	inFresh := true
	for _, line := range lines {
		if line == "" {
			inFresh = false
			continue
		}

		if inFresh {
			numsStr := strings.Split(line, "-")
			x, _ := strconv.Atoi(numsStr[0])
			y, _ := strconv.Atoi(numsStr[1])
			// nums := utils.Range(x, y)

			fresh = append(fresh, []int{
				x, y,
			})
		} else {
			id, _ := strconv.Atoi(line)
			IDs = append(IDs, id)
		}
	}

	partOne := 0
	for _, id := range IDs {
		seen := -1
		for _, fresh := range fresh {
			if id >= fresh[0] && id <= fresh[len(fresh)-1] {
				if seen == id {
					continue
				}

				seen = id
				partOne++
			}
		}
		seen = -1
	}

	partTwo := 0
	// Go through each fresh
	// If there is an overlap then merge the items
	// 3-5
	// run 5-3+1 = 3
	// Merge these so they become 10 - 20
	// 10-14
	// 16-20
	// 12-18
	// Run 20-10+1=11
	// 3+5+9=19

	fmt.Println(merge(fresh))

	fmt.Println(partOne)
	fmt.Println(partTwo)
}
