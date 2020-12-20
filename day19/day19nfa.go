package main

import (
	"fmt"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	paras := advent.InputParas()
	rules := paras[0]
	messages := paras[1]

	cnf1 := compile(rules)
	cnf2 := compile(append(rules, []string{"8: 42 | 42 8", "11: 42 31 | 42 11 31"}...))

	count1, count2 := 0, 0
	for _, message := range messages {
		if match(cnf1, message) {
			count1++
		}
		if match(cnf2, message) {
			count2++
		}
	}
	fmt.Println("Part 1:", count1)
	fmt.Println("Part 2:", count2)
}

func compile(lines []string) map[string]string {
	rules := map[string]string{}
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		rules[fields[0]] = fields[1]
	}
	return rules
}

func match(rules map[string]string, message string) bool {
	// Simulate a non-deterministic automata.

	nstate := []string{"0"}

	for _, char := range []byte(message) {
		var next []string

		var step func(string)
		step = func(state string) {
			if state == "" {
				return
			}

			head, rest := state, ""
			if i := strings.Index(head, " "); i >= 0 {
				head, rest = head[:i], head[i:]
			}

			if head[0] == '"' {
				if head[1] == char {
					next = append(next, strings.TrimPrefix(rest, " "))
				}
			} else if rule, ok := rules[head]; ok {
				for _, alt := range strings.Split(rule, " | ") {
					step(alt + rest)
				}
			} else {
				panic(head)
			}
		}

		for _, state := range nstate {
			step(state)
		}
		nstate = next
	}

	for _, state := range nstate {
		if len(state) == 0 {
			return true
		}
	}
	return false
}
