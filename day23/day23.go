package main

import "fmt"

const example = "389125467"
const input = "712643589"

func main() {
	for _, s := range []string{example, input} {
		fmt.Println("== input", s, "==")
		fmt.Println("Part 1:", part1(s))
		fmt.Println("Part 2:", part2(s))
		fmt.Println()
	}
}

func part1(seed string) int {
	next := play(seed, len(seed), 100)
	res := 0
	for x := next[1]; x != 1; x = next[x] {
		res = 10*res + x
	}
	return res
}

func part2(seed string) int {
	next := play(seed, 1e6, 1e7)
	return next[1] * next[next[1]]
}

func play(seed string, cups, rounds int) []int {
	// next[i] is the label of the cup that comes next (clockwise)
	// after the cup labeled i
	next := make([]int, 1+cups)

	last := 0
	for i := 0; i < cups; i++ {
		x := i + 1
		if i < len(seed) {
			x = int(seed[i] - '0')
		}

		next[last] = x
		last = x
	}

	current := next[0]
	next[last] = current
	next[0] = -1e9 // poison

	for i := 0; i < rounds; i++ {
		next1 := next[current]
		next2 := next[next1]
		next3 := next[next2]
		next4 := next[next3]

		dest := current
		for {
			dest--
			if dest == 0 {
				dest = cups
			}
			if dest != next1 && dest != next2 && dest != next3 {
				break
			}
		}

		next[current] = next4
		current = next4

		next[next3] = next[dest]
		next[dest] = next1
	}

	return next
}
