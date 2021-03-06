package main

import (
	"fmt"
	"sort"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	joltages := advent.InputInts()
	sort.Ints(joltages)

	// counts[x] is the number of x-jolt differences
	var counts [4]int

	// Loop invariant:
	// di is the number of distinct ways to get prev-i
	prev := 0
	d0, d1, d2 := 1, 0, 0

	for _, joltage := range joltages {
		counts[joltage-prev]++

		for prev < joltage {
			x := d0 + d1 + d2
			d0, d1, d2 = 0, d0, d1

			prev++
			if prev == joltage {
				d0 = x
			}
		}
	}

	counts[3]++
	fmt.Println("Part 1:", counts[1]*counts[3])

	fmt.Println("Part 2:", d0)
}
