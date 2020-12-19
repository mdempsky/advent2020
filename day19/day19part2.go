package main

import (
	"fmt"
	"log"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	paras := advent.InputParas()
	rules := paras[0]
	messages := paras[1]

	cnf := compile(rules)

	count := 0
	for _, message := range messages {
		if match(cnf, message) {
			count++
		}
	}
	fmt.Println("Part 1:", count)
}

type pair struct {
	a, b int
}

type char struct {
	x byte
}

type either struct {
	a, b int
}

type alias struct {
	a int
}

func (pair) isAlt()   {}
func (char) isAlt()   {}
func (either) isAlt() {}
func (alias) isAlt()  {}

type alt interface{ isAlt() }

type chomsky []alt

func compile(lines []string) chomsky {
	rules := map[string]string{}
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		rules[fields[0]] = fields[1]
	}

	memo := map[string]int{}
	var res chomsky

	emit := func(s string, a alt) int {
		n := len(res)
		res = append(res, a)
		memo[s] = n
		return n
	}

	var do func(s string) int
	do = func(s string) int {
		if x, ok := memo[s]; ok {
			return x
		}

		if rule, ok := rules[s]; ok {
			n := len(res)
			memo[s] = n
			res = append(res, nil)
			res[n] = alias{do(rule)}
			return n
		}

		if len(s) == 3 && s[0] == '"' {
			return emit(s, char{s[1]})
		}

		if f := strings.SplitN(s, " | ", 2); len(f) > 1 {
			return emit(s, either{do(f[0]), do(f[1])})
		}

		// "a b c" -> []string{"a", "b", "c"}
		if f := strings.SplitN(s, " ", 2); len(f) > 1 {
			return emit(s, pair{do(f[0]), do(f[1])})
		}

		panic(s)
	}

	if do("0") != 0 {
		panic("wat")
	}
	return res
}

func match(rules chomsky, message string) bool {
	// m is indexed: [start][end][rule]
	const N = 100
	if len(message) >= N {
		log.Fatal(message)
	}

	const M = 450
	if len(rules) > M {
		log.Fatal("rules length", len(rules))
	}

	var m [N][N][M]bool
	for w := 0; w <= len(message); w++ {
		for start := 0; start+w <= len(message); start++ {
			end := start + w
			var done [M]bool
			var do func(int)
			do = func(rule int) {
				if done[rule] {
					return
				}

				found := false
				switch alt := rules[rule].(type) {
				case char:
					if w == 1 && message[start] == alt.x {
						found = true
					}
				case pair:
					for mid := start + 1; mid < end; mid++ {
						if m[start][mid][alt.a] && m[mid][end][alt.b] {
							found = true
						}
					}
				case either:
					do(alt.a)
					do(alt.b)
					if m[start][end][alt.a] || m[start][end][alt.b] {
						found = true
					}
				case alias:
					do(alt.a)
					found = m[start][end][alt.a]
				default:
					panic(alt)
				}
				m[start][end][rule] = found

				done[rule] = true
			}

			for rule := range rules {
				do(rule)
			}
		}
	}

	return m[0][len(message)][0]
}
