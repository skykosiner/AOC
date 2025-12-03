package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/skykosiner/AOC/pkg/utils"
)

func getJolts(bank string, n int) int {
	str, prev := "", 0
	for b := n; b > 0; b-- {
		for i := prev; i <= len(bank)-b; i++ {
			if bank[i] > bank[prev] {
				prev = i
			}
		}

		str += string(bank[prev])
		prev++
	}

	num, _ := strconv.Atoi(str)
	return num
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

	partOne, partTwo := 0, 0

	for _, bank := range utils.ReadFile(inputFile) {
		if bank == "" {
			break
		}

		partOne += getJolts(bank, 2)
		partTwo += getJolts(bank, 12)
	}

	fmt.Println(partOne, partTwo)
}
