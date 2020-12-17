package main

import (
	"fmt"
	"log"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

const Max = 20

// valid[i][j] means that rule i might be for ticket field j.
var valid [Max][Max]bool

func main() {
	lines := advent.InputLines()
	nfields := 0
	for lines[nfields] != "" {
		nfields++
	}

	your := ints(lines[nfields+2])
	nearbys := lines[nfields+5:]

	var rules []rule
	for _, line := range lines[:nfields] {
		var rule rule
		fields := strings.Split(line, ": ")
		rule.name = fields[0]
		_, err := fmt.Sscanf(fields[1], "%d-%d or %d-%d", &rule.lo1, &rule.hi1, &rule.lo2, &rule.hi2)
		if err != nil {
			log.Fatalf("sscanf: %q: %v", fields[1], err)
		}
		rule.name = strings.TrimSuffix(rule.name, ":")
		rules = append(rules, rule)
	}

	for i := 0; i < nfields; i++ {
		for j := 0; j < nfields; j++ {
			valid[i][j] = true
		}
	}

	errorRate := 0
	for _, nearby := range nearbys {
		bad := false
		fields := ints(nearby)
	Fields:
		for _, field := range fields {
			for _, rule := range rules {
				if rule.valid(field) {
					continue Fields
				}
			}
			errorRate += field
			bad = true
		}
		if bad {
			continue
		}

		for j, field := range fields {
			for i, rule := range rules {
				if valid[i][j] && !rule.valid(field) {
					valid[i][j] = false
				}
			}
		}
	}

	fmt.Println("Part 1:", errorRate)

	doneRules := make([]bool, nfields)
	doneFields := make([]bool, nfields)
	fieldPlaces := make([]int, nfields)
	todo := nfields

	for todo > 0 {
	DoneRules:
		for i, done := range doneRules {
			if done {
				continue
			}

			x := -1
			for j := 0; j < nfields; j++ {
				if valid[i][j] {
					if x >= 0 {
						continue DoneRules
					}
					x = j
				}
			}
			if x == -1 {
				continue DoneRules
			}

			doneRules[i] = true
			doneFields[x] = true
			fieldPlaces[i] = x
			for i2 := 0; i2 < nfields; i2++ {
				if i2 != i {
					valid[i2][x] = false
				}
			}
			todo--
		}

	DoneFields:
		for j, done := range doneFields {
			if done {
				continue
			}

			x := -1
			for i := 0; i < nfields; i++ {
				if valid[i][j] {
					if x >= 0 {
						continue DoneFields
					}
					x = i
				}
			}
			if x == -1 {
				continue DoneFields
			}

			doneRules[x] = true
			doneFields[j] = true
			fieldPlaces[x] = j
			for j2 := 0; j2 < nfields; j2++ {
				if j2 != j {
					valid[x][j2] = false
				}
			}
			todo--
		}
	}

	product := 1
	for i, rule := range rules {
		if strings.HasPrefix(rule.name, "departure") {
			product *= your[fieldPlaces[i]]
		}
	}
	fmt.Println("Part 2:", product)
}

type rule struct {
	name               string
	lo1, hi1, lo2, hi2 int
}

func (r *rule) valid(x int) bool {
	return x >= r.lo1 && x <= r.hi1 ||
		x >= r.lo2 && x <= r.hi2
}

func ints(s string) []int {
	vals := strings.Split(s, ",")
	res := make([]int, len(vals))
	for i, val := range vals {
		res[i] = int(advent.Atoi(val))
	}
	return res
}
