package main

import (
	"fmt"
	"regexp"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	paras := advent.InputParas()
	rules := paras[0]
	messages := paras[1]

	rx := compile(rules)

	count := 0
	for _, message := range messages {
		if rx.MatchString(message) {
			count++
		}
	}
	fmt.Println("Part 1:", count)
}

func compile(lines []string) *regexp.Regexp {
	rules := map[string]string{}
	for _, line := range lines {
		fields := strings.Split(line, ": ")
		rules[fields[0]] = fields[1]
	}

	memo := map[string]string{}
	var eval func(string) string
	eval = func(key string) string {
		if res, ok := memo[key]; ok {
			return res
		}
		var buf strings.Builder
		rule, ok := rules[key]
		if !ok {
			panic(key)
		}
		buf.WriteString("(")
		for i, alt := range strings.Split(rule, " | ") {
			if i > 0 {
				buf.WriteString("|")
			}
			for _, part := range strings.Split(alt, " ") {
				if part[0] == '"' {
					buf.WriteString(strings.Trim(part, `"`))
					continue
				}
				buf.WriteString(eval(part))
			}
		}
		buf.WriteString(")")
		res := buf.String()
		memo[key] = res
		return res
	}

	return regexp.MustCompile("^" + eval("0") + "$")
}
