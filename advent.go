package advent

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var flagInput = flag.String("input", "input.txt", "input file to read from")

func Input() string {
	flag.Parse()

	fmt.Println("reading", *flagInput)

	buf, err := ioutil.ReadFile(*flagInput)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func InputLines() []string {
	return strings.Split(strings.Trim(Input(), "\n"), "\n")
}

func InputInts() []int {
	var res []int
	for _, n := range InputLines() {
		res = append(res, int(Atoi(n)))
	}
	return res
}

func Atoi(s string) int64 {
	n, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
