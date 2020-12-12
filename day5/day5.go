package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var seats []int
	for _, pass := range strings.Fields(string(buf)) {
		pass = strings.ReplaceAll(pass, "F", "0")
		pass = strings.ReplaceAll(pass, "B", "1")
		pass = strings.ReplaceAll(pass, "L", "0")
		pass = strings.ReplaceAll(pass, "R", "1")

		n, err := strconv.ParseInt(pass, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		seats = append(seats, int(n))
	}

	var max = -1
	for _, seat := range seats {
		if seat > max {
			max = seat
		}
	}
	fmt.Println("max", max)

	sort.Ints(seats)
	for i, x := range seats {
		if seats[i+1] != x+1 {
			fmt.Println("mine", x+1)
			break
		}
	}
}
