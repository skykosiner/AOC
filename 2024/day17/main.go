package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/skykosiner/AOC/pkg/utils"
)

func main() {
	var program []int
	// var registerA, registerB, registerC int

	lines := utils.ReadFile("./input.test")
	r, _ := regexp.Compile("[0-9]")
	for _, char := range r.FindAllString(lines[4], -1) {
		num, _ := strconv.Atoi(char)
		program = append(program, num)
	}

	fmt.Println(program)
}
