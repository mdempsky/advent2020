package main

import (
	"fmt"
	"io"
	"log"
	"sort"
)

func main() {
	var num []int
	for {
		var x int
		_, err := fmt.Scanf("%d", &x)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		num = append(num, x)
	}

	// Naive solution. O(N^2)
	for i, a := range num {
		for _, b := range num[i+1:] {
			if a+b == 2020 {
				fmt.Println("found:", a, b, a*b)
			}
		}
	}

	// Part 2: Naive solution.
	for i, a := range num {
		for j, b := range num[i+1:] {
			for _, c := range num[i+j+2:] {
				if a+b+c == 2020 {
					fmt.Println("found pt2:", a, b, c, a*b*c)
				}
			}
		}
	}

	// Asymptotically faster solution. O(N lg N)
	sort.Ints(num)
	for i, a := range num {
		if a > 1010 {
			// Short-circuit.
			break
		}
		if a == 1010 {
			if i+1 < len(num) && num[i+1] == 1010 {
				fmt.Println("faster:", 1010, 1010, 1010*1010)
			}
			continue
		}

		b := 2020 - a
		j := sort.Search(len(num), func(i int) bool { return num[i] >= b })
		if j < len(num) && num[j] == b {
			fmt.Println("faster:", a, b, a*b)
		}
	}
}
