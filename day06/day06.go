package main

import (
	"fmt"
	"math/bits"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	var anysum, everysum int

	for _, group := range strings.Split(advent.Input(), "\n\n") {
		var anyone, everyone uint64
		everyone = ^uint64(0) // 0xffffffffffffffff
		for _, person := range strings.Fields(group) {
			var x uint64 = 0
			for _, ch := range person {
				if n := uint(ch - 'a'); n <= 26 {
					x |= 1 << n
				}
			}
			anyone |= x
			everyone &= x
		}

		anysum += bits.OnesCount64(anyone)
		everysum += bits.OnesCount64(everyone)
	}

	fmt.Println(anysum, everysum)
}
