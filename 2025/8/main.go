package main

import (
	"flag"
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

	var nums [][]int
	for _, line := range utils.ReadFile(inputFile) {
		n := strings.Split(line, ",")
		x, _ := strconv.Atoi(n[0])
		y, _ := strconv.Atoi(n[1])
		z, _ := strconv.Atoi(n[2])

		nums = append(nums, []int{
			x, y, z,
		})
	}

	indexes
}
