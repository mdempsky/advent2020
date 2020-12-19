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

type pair struct {
	a, b int
}

type char struct {
	x byte
}

func (pair) isAlt() {}
func (char) isAlt() {}

type alt interface{ isAlt() }

type chomsky [][]alt

func compile(lines []string) chomsky {
	rules := map[string]string{}
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		rules[fields[0]] = fields[1]
	}

	memo := map[string]int{}
	alias := map[int][]int{}
	var res chomsky

	var do func(s string) int
	do = func(s string) int {
		if x, ok := memo[s]; ok {
			return x
		}

		n := len(res)
		res = append(res, nil)
		memo[s] = n

		if rule, ok := rules[s]; ok {
			s = rule
		}

		for _, f := range strings.Split(s, " | ") {
			if f := strings.SplitN(f, " ", 2); len(f) > 1 {
				res[n] = append(res[n], pair{do(f[0]), do(f[1])})
				continue
			}
			if f[0] == '"' {
				res[n] = append(res[n], char{f[1]})
				continue
			}
			alias[n] = append(alias[n], do(f))
		}

		return n
	}
	do("0")

	// Eliminate unit rules.
	var elim func(int)
	elim = func(key int) {
		if vals, ok := alias[key]; ok {
			for _, val := range vals {
				elim(val)
				res[key] = append(res[key], res[val]...)
			}
			delete(alias, key)
		}
	}
	for key := range alias {
		elim(key)
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
	for w := 1; w <= len(message); w++ {
		for start := 0; start+w <= len(message); start++ {
			end := start + w
			for rule, alts := range rules {
				found := false
				for _, alt := range alts {
					switch alt := alt.(type) {
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
					default:
						panic(alt)
					}
					if found {
						break
					}
				}
				m[start][end][rule] = found
			}
		}
	}

	return m[0][len(message)][0]
}
