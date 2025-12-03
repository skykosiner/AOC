package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

type IncOrDec int

const (
	incerase IncOrDec = iota
	decrease
)

type Level struct {
	value    int
	incOrDec IncOrDec
}

func PartOne() {
	var safeLevels int
	lines := utils.ReadFile("./input")

OuterLoop:
	for _, line := range lines {
		var levels []Level
		rawLevel := strings.Split(line, " ")

		if len(rawLevel) <= 1 {
			break
		}

		for idx, level := range rawLevel {
			l, _ := strconv.Atoi(level)
			if len(levels) == 0 {
				levels = append(levels, Level{
					value: l,
				})

				continue
			}

			if levels[idx-1].value < l {
				levels = append(levels, Level{
					value:    l,
					incOrDec: decrease,
				})
			} else {
				levels = append(levels, Level{
					value:    l,
					incOrDec: incerase,
				})
			}

		}

		if levels[0].value > levels[1].value {
			levels[0].incOrDec = incerase
		} else {
			levels[0].incOrDec = decrease
		}

		for idx := range levels {
			if idx != 0 {
				diff := int(math.Abs(float64(levels[idx].value) - float64(levels[idx-1].value)))
				if (diff < 1 || diff > 3) || (levels[idx].incOrDec != levels[idx-1].incOrDec) {
					continue OuterLoop
				}
			}
		}

		safeLevels += 1
	}

	fmt.Println(safeLevels)
}

func PartTwo() {
	var safeLevels int
	lines := utils.ReadFile("./input")

OuterLoop:
	for _, line := range lines {
		var levels []Level
		rawLevel := strings.Split(line, " ")

		if len(rawLevel) <= 1 {
			break
		}

		for idx, level := range rawLevel {
			l, _ := strconv.Atoi(level)
			if len(levels) == 0 {
				levels = append(levels, Level{
					value: l,
				})

				continue
			}

			if levels[idx-1].value < l {
				levels = append(levels, Level{
					value:    l,
					incOrDec: decrease,
				})
			} else {
				levels = append(levels, Level{
					value:    l,
					incOrDec: incerase,
				})
			}

		}

		if levels[0].value > levels[1].value {
			levels[0].incOrDec = incerase
		} else {
			levels[0].incOrDec = decrease
		}

		badLevel := false
		for idx := range levels {
			if idx != 0 {
				diff := int(math.Abs(float64(levels[idx].value - levels[idx-1].value)))
				if (diff < 1 || diff > 3) || (levels[idx].incOrDec != levels[idx-1].incOrDec) {
					if badLevel {
						continue OuterLoop
					} else {
						badLevel = true
					}
				}
			}
		}

		safeLevels += 1
	}

	fmt.Println(safeLevels)
}

func main() {
	PartOne()
	PartTwo()
}
