package main

import (
	"bytes"
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

type chomsky struct {
	chars [][]byte
	pairs [][][2]int
}

func compile(lines []string) chomsky {
	rules := map[string]string{}
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		rules[fields[0]] = fields[1]
	}

	// Compile into Chomsky normal form.

	memo := map[string]int{}
	var chars [][]byte
	var pairs [][][2]int
	var aliases [][]int

	// Eliminate rules with nonsolitary terminals and right-hand sides
	// with more than 2 nonterminals.
	var do func(s string) int
	do = func(s string) int {
		if rule, ok := rules[s]; ok {
			s = rule
		}

		if n, ok := memo[s]; ok {
			return n
		}

		n := len(memo)
		memo[s] = n
		chars = append(chars, nil)
		pairs = append(pairs, nil)
		aliases = append(aliases, nil)

		for _, f := range strings.Split(s, " | ") {
			if f := strings.SplitN(f, " ", 2); len(f) > 1 {
				pairs[n] = append(pairs[n], [2]int{do(f[0]), do(f[1])})
				continue
			}
			if f[0] == '"' {
				chars[n] = append(chars[n], f[1])
				continue
			}
			aliases[n] = append(aliases[n], do(f))
		}

		return n
	}
	do("0")

	// Eliminate unit rules.
	var elim func(int)
	elim = func(key int) {
		for _, val := range aliases[key] {
			elim(val)
			chars[key] = append(chars[key], chars[val]...)
			pairs[key] = append(pairs[key], pairs[val]...)
		}
		aliases[key] = nil
	}
	for key := range aliases {
		elim(key)
	}

	return chomsky{chars, pairs}
}

func match(rules chomsky, message string) bool {
	// Simulate a non-deterministic automata.

	var step func([][]int, byte, int, []int) [][]int
	step = func(dst [][]int, char byte, rule int, rest []int) [][]int {
		if bytes.IndexByte(rules.chars[rule], char) >= 0 {
			dst = append(dst, rest)
		}
		for _, pair := range rules.pairs[rule] {
			dst = step(dst, char, pair[0], append(pair[1:], rest...))
		}
		return dst
	}

	nstate := [][]int{{0}}

	for _, char := range []byte(message) {
		var next [][]int
		for _, state := range nstate {
			if len(state) > 0 {
				next = step(next, char, state[0], state[1:])
			}
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
