package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

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

	// var ranges []Range
	sort.Slice(fresh, func(i, j int) bool {
		return fresh[i][0] < fresh[j][0]
	})

	merged := fresh[:0]
	currStart, currEnd := fresh[0][0], fresh[0][1]
	for i := 1; i < len(fresh); i++ {
		start, end := fresh[i][0], fresh[i][1]
		if start <= currEnd { // it overlaps
			if end > currEnd {
				currEnd = end
			}
		} else {
			merged = append(merged, []int{currStart, currEnd})
			currStart, currEnd = start, end
		}
	}

	merged = append(merged, []int{currStart, currEnd})

	for _, r := range merged {
		fmt.Println(r)
		partTwo += r[1] - r[0] + 1
	}

	fmt.Println(partOne)
	fmt.Println(partTwo)
}
