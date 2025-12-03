package main

import (
	"fmt"
	"math"

	"github.com/skykosiner/AOC/pkg/utils"
)

const WIDTH = 101
const HEIGHT = 103

type cords struct {
	x, y int
}

type Robot struct {
	P cords
	V cords
}

func (r *Robot) moveCords() {
	r.P.x = (r.P.x + r.V.x + WIDTH) % WIDTH
	r.P.y = (r.P.y + r.V.y + HEIGHT) % HEIGHT
}

func NewRobot(px, py, vx, vy int) *Robot {
	return &Robot{
		P: cords{
			x: px,
			y: py,
		},
		V: cords{
			x: vx,
			y: vy,
		},
	}
}

func solve(partTwo bool) int {
	var robots []*Robot
	lines := utils.ReadFile("./input")

	for _, line := range lines {
		nums := utils.ExtractNums(line)
		if len(nums) == 0 {
			break
		}

		robots = append(robots, NewRobot(nums[0], nums[1], nums[2], nums[3]))
	}

	if partTwo {
		i := 0

		for {
			i++
			seen := map[[2]int]bool{}
			for _, robot := range robots {
				robot.moveCords()
				seen[[2]int{robot.P.x, robot.P.y}] = true
			}

			if len(seen) == len(robots) {
				break
			}
		}

		return i
	}

	for range 100 {
		for _, robot := range robots {
			robot.moveCords()
		}
	}

	total := 1
	midX, midY := int(math.Floor(WIDTH)/2), int(math.Floor(HEIGHT/2))

	quadrantCounts := [4]int{0, 0, 0, 0}

	for _, robot := range robots {
		if robot.P.x == midX || robot.P.y == midY {
			continue
		}
		if robot.P.x < midX && robot.P.y < midY {
			quadrantCounts[0]++
		} else if robot.P.x >= midX && robot.P.y < midY {
			quadrantCounts[1]++
		} else if robot.P.x < midX && robot.P.y >= midY {
			quadrantCounts[2]++
		} else if robot.P.x >= midX && robot.P.y >= midY {
			quadrantCounts[3]++
		}
	}

	for _, count := range quadrantCounts {
		total *= count
	}

	return total
}

func main() {
	fmt.Println("Part One:", solve(false))
	fmt.Println("Part Two:", solve(true))
}
