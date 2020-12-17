package main

import "fmt"

func solve(input []int, want int) int {
	// seen maps numbers to the turn we last saw it, if any.
	var seen []int

	set := func(i, val int) {
		if i >= len(seen) {
			seen = append(seen, make([]int, i+1-len(seen))...)
		}
		seen[i] = val
	}

	for i, x := range input[:len(input)-1] {
		set(x, 1+i)
	}

	turn := len(input)
	next := input[len(input)-1]

	for {
		prevSeen, ok := 0, false
		if next < len(seen) && seen[next] != 0 {
			prevSeen, ok = seen[next], true
		}
		set(next, turn)

		if !ok {
			next = 0
		} else {
			next = turn - prevSeen
		}
		turn++

		if turn == want {
			return next
		}
	}
}

func main() {
	var example = []int{0, 3, 6}
	var input = []int{16, 1, 0, 18, 12, 14, 19}

	if false {
		fmt.Println(solve(example, 2020))
		fmt.Println(solve(input, 2020))
	}

	const big = 30000000
	fmt.Println(solve(input, big))

	return

	fmt.Println(solve([]int{0, 3, 6}, big), 175594)
	fmt.Println(solve([]int{1, 3, 2}, big), 2578)
	fmt.Println(solve([]int{2, 1, 3}, big), 3544142)
	fmt.Println(solve([]int{1, 2, 3}, big), 261214)
	fmt.Println(solve([]int{2, 3, 1}, big), 6895259)
	fmt.Println(solve([]int{3, 2, 1}, big), 18)
	fmt.Println(solve([]int{3, 1, 2}, big), 362)
}
