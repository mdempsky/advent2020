package main

import (
	"fmt"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	lines := advent.InputLines()

	earliest := int(advent.Atoi(lines[0]))

	var routes []int
	for _, route := range strings.Split(lines[1], ",") {
		if route == "x" {
			routes = append(routes, 0)
			continue
		}
		routes = append(routes, int(advent.Atoi(route)))
	}

Outer:
	for i := 0; ; i++ {
		for _, route := range routes {
			if route == 0 {
				continue
			}
			if (earliest+i)%route == 0 {
				fmt.Println("Part 1:", i*route)
				break Outer
			}
		}
	}

	t, modulo := 0, 1
	for i, route := range routes {
		if route != 0 {
			for (t+i)%route != 0 {
				t += modulo
			}
			modulo *= route / gcd(modulo, route)
		}
	}
	fmt.Println("Part 2:", t)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}