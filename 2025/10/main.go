package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

type Machine struct {
	Buttons      []uint64
	InitialState uint64
	N            int
}

func NewMachine(inputButtons [][]int, finalState []string) Machine {
	n := len(finalState)
	initialState := uint64(0)
	for i, char := range finalState {
		if char == "#" {
			initialState |= (1 << i)
		}
	}

	buttons := make([]uint64, 0, len(inputButtons))
	for _, sec := range inputButtons {
		mask := uint64(0)
		for _, idx := range sec {
			if idx >= 0 && idx < n {
				mask |= (1 << idx)
			}
		}

		buttons = append(buttons, mask)
	}

	return Machine{
		Buttons:      buttons,
		InitialState: initialState,
		N:            n,
	}
}

func (m Machine) solve() int {
	if m.InitialState == 0 {
		return 0
	}

	queue := []uint64{m.InitialState}
	visted := make(map[uint64]bool)
	visted[m.InitialState] = true
	steps := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for range levelSize {
			currentState := queue[0]
			queue = queue[1:]

			for _, buttonMask := range m.Buttons {
				newState := currentState ^ buttonMask
				if newState == 0 {
					return steps + 1
				}

				if _, found := visted[newState]; !found {
					visted[newState] = true
					queue = append(queue, newState)
				}
			}
		}
		steps++
	}

	return 0
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

	finalStateRegex := regexp.MustCompile(`\[(.*?)\]`)
	buttonsRegex := regexp.MustCompile(`\((.*?)\)`)

	var machines []Machine
	for _, line := range utils.ReadFile(inputFile) {
		finalStateMatch := finalStateRegex.FindStringSubmatch(line)
		finalState := strings.Split(finalStateMatch[1], "")

		var buttons [][]int
		buttonMatches := buttonsRegex.FindAllStringSubmatch(line, -1)
		for _, match := range buttonMatches {
			var buttonSector []int
			buttonsStr := strings.Split(match[1], ",")
			for _, button := range buttonsStr {
				b, _ := strconv.Atoi(button)
				buttonSector = append(buttonSector, b)
			}

			buttons = append(buttons, buttonSector)
		}

		machines = append(machines, NewMachine(buttons, finalState))
	}

	partOne := 0
	for _, machine := range machines {
		partOne += machine.solve()
	}

	fmt.Println(partOne)
}
