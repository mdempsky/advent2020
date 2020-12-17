package main

import (
	"fmt"
	"sort"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	var nums []int
	for _, line := range advent.InputLines() {
		nums = append(nums, int(advent.Atoi(line)))
	}

	var found int
Outer:
	for i, n := range nums[25:] {
		preamble := nums[i : i+25]
		for j, a := range preamble {
			for _, b := range preamble[j+1:] {
				if a+b == n {
					continue Outer
				}
			}
		}
		found = n
		break
	}

	fmt.Println(found)

	i, j := 0, 0
	var sum int
Outer2:
	for {
		switch {
		case sum < found:
			sum += nums[j]
			j++
		case sum > found:
			sum -= nums[i]
			i++
		default:
			break Outer2
		}
	}

	sort.Ints(nums[i:j])
	fmt.Println(nums[i] + nums[j-1])
}
