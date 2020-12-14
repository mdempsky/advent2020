package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	lines := advent.InputLines()

	var set, clear int64

	part1 := make(map[int]int)
	part2 := make(map[int]int)
	for _, cmd := range lines[:] {
		if strings.HasPrefix(cmd, "mask = ") {
			mask := strings.Fields(cmd)[2]

			var err error
			set, err = strconv.ParseInt(strings.ReplaceAll(mask, "X", "0"), 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			clear, err = strconv.ParseInt(strings.ReplaceAll(mask, "X", "1"), 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			continue
		}

		fields := strings.Split(strings.TrimPrefix(cmd, "mem["), "] = ")
		addr := int(advent.Atoi(fields[0]))
		data := int(advent.Atoi(fields[1]))

		part1[addr] = (data & int(clear)) | int(set)

		var do func(int, int)
		do = func(addr, float int) {
			if float == 0 {
				part2[addr] = data
				return
			}

			newfloat := float & (float - 1)
			floatlsb := newfloat ^ float

			do(addr&^floatlsb, newfloat)
			do(addr|floatlsb, newfloat)
		}
		do(addr|int(set), int(set^clear))
	}

	var sum1 int
	for _, v := range part1 {
		sum1 += v
	}
	fmt.Println("Part 1:", sum1)

	var sum2 int
	for _, v := range part2 {
		sum2 += v
	}
	fmt.Println("Part 2:", sum2)
}
