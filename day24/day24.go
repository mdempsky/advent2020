package main

import (
	"fmt"

	advent "github.com/mdempsky/advent2020"
)

type pos [2]int

func main() {
	black := map[pos]bool{}

	for _, line := range advent.InputLines() {
		const (
			Y = 0
			X = 1
		)

		var tile pos
		for i := 0; i < len(line); i++ {
			dx := 2
			switch line[i] {
			case 'n':
				tile[Y] += 2
				fallthrough
			case 's':
				tile[Y]--
				dx--
				i++
			}

			switch line[i] {
			case 'e':
				tile[X] += dx
			case 'w':
				tile[X] -= dx
			default:
				panic(line[i])
			}
		}

		black[tile] = !black[tile]
	}

	fmt.Println("Part 1:", count(black))

	for i := 0; i < 100; i++ {
		black = next(black)
	}

	fmt.Println("Part 2:", count(black))
}

func count(black map[pos]bool) int {
	res := 0
	for _, x := range black {
		if x {
			res++
		}
	}
	return res
}

var neighbors = [6]pos{
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
	{0, 2},
	{0, -2},
}

func (p pos) add(q pos) pos {
	return pos{p[0] + q[0], p[1] + q[1]}
}

func next(black map[pos]bool) map[pos]bool {

	res := map[pos]bool{}
	for tile, x := range black {
		if !x {
			continue
		}

		bn := 0
		for _, d := range &neighbors {
			n := tile.add(d)
			if black[n] {
				bn++
				continue
			}

			bn2 := 0
			for _, d := range &neighbors {
				m := n.add(d)
				if black[m] {
					bn2++
				}
			}
			if bn2 == 2 {
				res[n] = true
			}
		}

		if bn >= 1 && bn <= 2 {
			res[tile] = true
		}
	}

	return res
}
