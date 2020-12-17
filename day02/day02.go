package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	var valid1, valid2 int
	for {
		var lo, hi int
		var ch byte
		var password string

		_, err := fmt.Scanf("%d-%d %c: %s", &lo, &hi, &ch, &password)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		n := strings.Count(password, string(ch))
		if n >= lo && n <= hi {
			valid1++
		}

		if ( /*lo <= len(password) && */ password[lo-1] == ch) !=
			( /*hi <= len(password) && */ password[hi-1] == ch) {
			valid2++
		}
	}
	fmt.Println("valid:", valid1)
	fmt.Println("valid:", valid2)
}
