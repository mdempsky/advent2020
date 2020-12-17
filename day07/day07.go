package main

import (
	"fmt"
	"log"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	inward := map[string]map[string]int{}
	outward := map[string][]string{}

	for _, rule := range strings.SplitAfter(advent.Input(), "\n") {
		if rule == "" {
			continue
		}
		rule = strings.TrimSuffix(rule, ".\n")

		p := strings.SplitN(rule, " bags contain ", 2)
		if len(p) < 2 {
			log.Fatalf("split %q into %q", rule, p)
		}
		outer, inners := p[0], p[1]

		if inners == "no other bags" {
			continue
		}

		for _, inner := range strings.Split(inners, ", ") {
			inner = strings.TrimSuffix(inner, "s")
			inner = strings.TrimSuffix(inner, " bag")

			parts := strings.SplitN(inner, " ", 2)
			var count int
			_, err := fmt.Sscanf(parts[0], "%d", &count)
			if err != nil {
				log.Fatal(err)
			}
			bag := parts[1]

			outward[bag] = append(outward[bag], outer)

			m := inward[outer]
			if m == nil {
				m = make(map[string]int)
				inward[outer] = m
			}
			m[bag] = count
			fmt.Printf("%q %q %d\n", outer, bag, count)
		}
	}

	var todo []string
	seen := map[string]bool{}

	todo = append(todo, outward["shiny gold"]...)
	for len(todo) > 0 {
		var bag string
		bag, todo = todo[len(todo)-1], todo[:len(todo)-1]

		if !seen[bag] {
			seen[bag] = true
			todo = append(todo, outward[bag]...)
		}
	}

	fmt.Println(len(seen))

	var do func(string) int
	do = func(bag string) int {
		fmt.Printf("do(%q)\n", bag)
		sum := 1
		for inner, count := range inward[bag] {
			fmt.Printf("  %q %q %d\n", bag, inner, count)
			sum += count * do(inner)
		}
		return sum
	}
	fmt.Println(do("shiny gold") - 1) // -1, because how many *inside* the shiny gold bag
}
