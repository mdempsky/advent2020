package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(strings.Trim(string(buf), "\n"), "\n")
	fmt.Println(len(rows))

	var trees [5]int
	for i, row := range rows {
		if row[(1*i)%len(row)] == '#' {
			trees[0]++
		}
		if row[(3*i)%len(row)] == '#' {
			trees[1]++
		}
		if row[(5*i)%len(row)] == '#' {
			trees[2]++
		}
		if row[(7*i)%len(row)] == '#' {
			trees[3]++
		}
		if i%2 == 0 && row[(i/2)%len(row)] == '#' {
			trees[4]++
		}
	}
	fmt.Println(trees[1])
	fmt.Println(trees[0] * trees[1] * trees[2] * trees[3] * trees[4])

}
