package main

import (
	"bytes"
	"fmt"
	"log"
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
		if x, ok := memo[s]; ok {
			return x
		}

		n := len(memo)
		memo[s] = n
		chars = append(chars, nil)
		pairs = append(pairs, nil)
		aliases = append(aliases, nil)

		if rule, ok := rules[s]; ok {
			s = rule
		}

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
	const N = 100
	if len(message) >= N {
		log.Fatal(message)
	}

	const M = 200
	if len(rules.chars) > M {
		log.Fatal("rules length", len(rules.chars))
	}

	// CYK algorithm

	// m is indexed: [start][end][rule]
	var m [N][N][M]bool

	for start := 0; start < len(message); start++ {
		for rule, chars := range rules.chars {
			if bytes.IndexByte(chars, message[start]) >= 0 {
				m[start][start+1][rule] = true
			}
		}
	}

	for w := 2; w <= len(message); w++ {
		for start := 0; start+w <= len(message); start++ {
			end := start + w
		Rules:
			for rule, pairs := range rules.pairs {
				for _, pair := range pairs {
					for mid := start + 1; mid < end; mid++ {
						if m[start][mid][pair[0]] && m[mid][end][pair[1]] {
							m[start][end][rule] = true
							continue Rules
						}
					}
				}
			}
		}
	}

	return m[0][len(message)][0]
}
