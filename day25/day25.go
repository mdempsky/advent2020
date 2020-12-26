package main

import (
	"fmt"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	nums := advent.InputInts()
	fmt.Println("Round 1:", pow(nums[0], log(7, nums[1])))
}

func pow(base, x int) int {
	value := 1
	for i := 0; i < x; i++ {
		value *= base
		value %= 20201227
	}
	return value
}

func log(base, x int) int {
	value := 1
	loops := 0
	for value != x {
		value *= base
		value %= 20201227
		loops++
	}
	return loops
}
