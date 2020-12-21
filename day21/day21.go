package main

import (
	"fmt"
	"sort"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	counts := map[string]int{}

	possible := map[string]map[string]bool{}

	for _, line := range advent.InputLines() {
		parts := strings.SplitN(line, " (contains ", 2)
		var allergens []string
		if len(parts) >= 2 {
			allergens = strings.Split(strings.TrimSuffix(parts[1], ")"), ", ")
		}
		parts = strings.Split(parts[0], " ")

		for _, part := range parts {
			counts[part]++
		}

		for _, allergen := range allergens {
			m2 := map[string]bool{}
			for _, part := range parts {
				m2[part] = true
			}

			if m, ok := possible[allergen]; ok {
				for k := range m {
					if !m2[k] {
						delete(m, k)
					}
				}
			} else {
				possible[allergen] = m2
			}
		}
	}

	fmt.Println(possible)

	type occur struct {
		allergen    string
		ingredients []string
	}
	var occurs []occur
	for allergen, m := range possible {
		o := occur{allergen: allergen}
		for ingredient := range m {
			o.ingredients = append(o.ingredients, ingredient)
		}
		occurs = append(occurs, o)
	}
	sort.Slice(occurs, func(i, j int) bool {
		return len(occurs[i].ingredients) < len(occurs[j].ingredients)
	})

	contains := map[string]string{}
	found := map[string]bool{}

	var brute func(int) bool
	brute = func(i int) bool {
		if i == len(occurs) {
			return true
		}
		occur := occurs[i]
		for _, ingredient := range occur.ingredients {
			if x, ok := contains[ingredient]; ok {
				if x == occur.allergen && brute(i+1) {
					return true
				}
			} else if !found[occur.allergen] {
				contains[ingredient] = occur.allergen
				found[occur.allergen] = true
				if brute(i + 1) {
					return true
				}
				delete(contains, ingredient)
				delete(found, occur.allergen)
			}
		}
		return false
	}
	if !brute(0) {
		panic("failed")
	}

	fmt.Println(occurs)
	fmt.Println(contains)

	count := 0
	var dangerous []string
	for name, x := range counts {
		if contains[name] == "" {
			count += x
		} else {
			dangerous = append(dangerous, name)
		}
	}
	fmt.Println("Part 1:", count)

	sort.Slice(dangerous, func(i, j int) bool {
		return contains[dangerous[i]] < contains[dangerous[j]]
	})
	fmt.Println("Part 2:", strings.Join(dangerous, ","))
}
