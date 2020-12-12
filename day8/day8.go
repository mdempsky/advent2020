package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func run(prog []string) (rax int64, loop bool) {
	seen := map[int64]bool{}
	var pc int64

	for {
		if pc == int64(len(prog)) {
			return rax, false
		}
		if seen[pc] {
			return rax, true
		}
		seen[pc] = true

		instr := strings.Fields(prog[pc])
		if len(instr) < 2 {
			log.Fatalf("why? %q", prog[pc])
		}
		pc++

		n, err := strconv.ParseInt(instr[1], 0, 64)
		if err != nil {
			log.Fatal(err)
		}

		switch instr[0] {
		case "acc":
			rax += n
		case "nop":
		case "jmp":
			pc += n - 1
		default:
			log.Fatalf("unknown instruction %q", instr[0])
		}
	}
}

func main() {
	prog := advent.InputLines()

	// Part 1.
	rax, _ := run(prog)
	fmt.Println(rax)

	// Part 2.
	for i, instr := range prog {
		switch strings.Fields(instr)[0] {
		case "jmp":
			prog[i] = "nop" + instr[3:]
		case "nop":
			prog[i] = "jmp" + instr[3:]
		default:
			continue
		}

		rax, loop := run(prog)
		if !loop {
			fmt.Println(rax)
			break
		}

		prog[i] = instr
	}
}
