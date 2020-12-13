package main

import (
	"fmt"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	lines := advent.InputLines()
	earliest := int(advent.Atoi(lines[0]))

	// Part 1.
	best := 0
	wait := func(route int) int { return mod(-earliest, route) }

	// Part 2.
	t, modulo := 0, 1

	for i, s := range strings.Split(lines[1], ",") {
		if s == "x" {
			continue
		}
		route := int(advent.Atoi(s))

		// Part 1.
		if best == 0 || wait(route) < wait(best) {
			best = route
		}

		// Part 2.
		for (t+i)%route != 0 {
			t += modulo
		}
		modulo *= route / gcd(modulo, route)
	}

	fmt.Println("Part 1:", best*wait(best))
	fmt.Println("Part 2:", t)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mod(x, m int) int {
	x %= m
	if x < m {
		x += m
	}
	return x
}
