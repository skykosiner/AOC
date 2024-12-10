package main

import (
	"fmt"

	"github.com/skykosiner/AOC/pkg/utils"
)

func main() {
	var disk []int
	fid := 0

	for idx, char := range utils.ReadFile("./input.test")[0] {
		x := int(char - '0')
		if idx%2 == 0 {
			for range x {
				disk = append(disk, fid)
			}

			fid++
		} else {
			for range x {
				disk = append(disk, -1)
			}
		}
	}

	var blanks []int

	for idx, x := range disk {
		if x == -1 {
			blanks = append(blanks, idx)
		}
	}

	for _, blank := range blanks {
		for len(disk) > 0 && disk[len(disk)-1] == -1 {
			disk = disk[:len(disk)-1]
		}

		if len(disk) <= blank {
			break
		}

		lastElement := disk[len(disk)-1]
		disk = disk[:len(disk)-1]

		disk[blank] = lastElement
		fmt.Println(disk)
	}

	total := 0
	for idx, num := range disk {
		total += num * idx
	}

	fmt.Println(total)
}
