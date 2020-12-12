package main

import (
	"fmt"
	"sort"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	joltages := advent.InputInts()
	sort.Ints(joltages)

	var counts [3]int

	// Loop invariant:
	// distinct[0] is the number of ways to get (previous joltage)-3
	// distinct[1] is ... (previous joltage)-2
	// distinct[2] is ... (previous joltage)-1
	// distinct[3] is ... (previous joltage)
	var distinct [4]int
	distinct[3] = 1

	for i, joltage := range joltages {
		prev := 0
		if i > 0 {
			prev = joltages[i-1]
		}
		diff := joltage - prev

		counts[diff-1]++

		for i := 0; i < diff; i++ {
			copy(distinct[:], distinct[1:])
			distinct[3] = 0
		}
		distinct[3] = distinct[0] + distinct[1] + distinct[2]
	}

	counts[2]++
	fmt.Println(counts[0] * counts[2])

	fmt.Println(distinct[3])

}
