package main

import (
	"fmt"
	"log"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	var x, y = 0, 0    // ship position
	var dx, dy = 10, 1 // waypoint position (relative to ship)

	for _, instr := range advent.InputLines() {
		action, val := instr[0], int(advent.Atoi(instr[1:]))

		switch action {
		case 'N':
			dy += val
		case 'E':
			dx += val
		case 'S':
			dy -= val
		case 'W':
			dx -= val
		case 'F':
			x += dx * val
			y += dy * val
		case 'R':
			val = 360 - val
			fallthrough
		case 'L':
			for val >= 90 {
				dx, dy = -dy, dx
				val -= 90
			}

			if val != 0 {
				log.Fatalf("bad turn: %q, %v", instr, val)
			}
		default:
			log.Fatalf("bad instruction: %q", instr)
		}

		fmt.Println(instr, x, y)
	}

	fmt.Println(x, y)

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	fmt.Println(x + y)
}
