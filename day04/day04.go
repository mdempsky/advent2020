package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	passports := strings.Split(strings.Trim(string(buf), "\n"), "\n\n")

	var valid int
Outer:
	for _, passport := range passports {
		want := map[string]*regexp.Regexp{
			"byr": regexp.MustCompile("^(19[2-9][0-9]|200[0-2])$"),
			"iyr": regexp.MustCompile("^20(1[0-9]|20)$"),
			"eyr": regexp.MustCompile("^20([1-2][0-9]|30)$"),
			"hgt": regexp.MustCompile("^(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)$"),
			"hcl": regexp.MustCompile("^#[0-9a-f]{6}$"),
			"ecl": regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$"),
			"pid": regexp.MustCompile("^[0-9]{9}$"),
			"cid": regexp.MustCompile("^.*$"),
		}

		for _, field := range strings.Fields(passport) {
			split := strings.SplitN(field, ":", 2)
			name, value := split[0], strings.Trim(split[1], " \n")
			pattern, ok := want[name]
			if !ok {
				continue Outer
			}
			if !pattern.MatchString(value) {
				fmt.Printf("%v / %v didn't match %q\n", name, pattern, value)
				continue Outer
			}
			delete(want, name)
		}
		delete(want, "cid")
		if len(want) == 0 {
			valid++
		}
	}
	fmt.Println(valid)
}
