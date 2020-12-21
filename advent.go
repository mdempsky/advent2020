package advent

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

func InputParas() [][]string {
	var res [][]string
	for _, para := range strings.Split(strings.Trim(Input(), "\n"), "\n\n") {
		res = append(res, strings.Split(para, "\n"))
	}
	return res
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

func Sscanf(str, format string, a ...interface{}) {
	_, err := fmt.Sscanf(str, format, a...)
	if err != nil {
		log.Fatal(err)
	}
}

func NextPermutation(data sort.Interface) bool {
	n := data.Len()
	for i := n - 2; i >= 0; i-- {
		if data.Less(i, i+1) {
			j := n - 1
			for !data.Less(i, j) {
				j--
			}
			data.Swap(i, j)
			reverse(data, i+1, n)
			return true
		}
	}

	reverse(data, 0, n)
	return false
}

func reverse(data sort.Interface, beg, end int) {
	end--
	for beg < end {
		data.Swap(beg, end)

		beg++
		end--
	}
}

func EachPermutation(data sort.Interface, f func()) {
	for {
		f()
		if !NextPermutation(data) {
			break
		}
	}
}
