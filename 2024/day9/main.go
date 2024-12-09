package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/skykosiner/AOC/pkg/utils"
)

func correct(blocks []string) bool {
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == "." {
			for j := i + 1; j < len(blocks); j++ {
				if blocks[j] != "." {
					return false
				}
			}
			return true
		}
	}
	return true
}

func main() {
	t := time.Now()
	var blocks []string
	sum := 0
	line := utils.ReadFile("./input")[0]

	currID := 0
	for idx, char := range line {
		if idx%2 == 0 {
			for range int(char - '0') {
				blocks = append(blocks, strconv.Itoa(currID))
			}

			currID++
		} else {
			for range int(char - '0') {
				blocks = append(blocks, ".")
			}
		}
	}

	for {
		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i] != "." {
				block := blocks[i]

				for j := 0; j < len(blocks); j++ {
					if blocks[j] == "." {
						blocks[j] = block
						blocks[i] = "."
						fmt.Println(blocks)
						break
					}
				}

				break
			}
		}

		if correct(blocks) {
			break
		}
	}

	for idx, block := range blocks {
		num, err := strconv.Atoi(block)
		if err != nil {
			continue
		}

		sum += idx * num
	}

	fmt.Println(sum)
	since := time.Since(t)
	fmt.Println(since.Minutes(), since.Seconds())
}
