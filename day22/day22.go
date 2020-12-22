package main

import (
	"fmt"
	"reflect"

	advent "github.com/mdempsky/advent2020"
)

type game [2][]int

func (g *game) step() {
	var x [2]int
	for i := range x {
		x[i], g[i] = g[i][0], g[i][1:]
	}
	for i := range x {
		if x[i] > x[1-i] {
			g[i] = g[i][:len(g[i]):len(g[i])]
			g[i] = append(g[i], x[i], x[1-i])
			return
		}
	}
	panic("huh?")
}

func (g *game) done() bool {
	return len(g[0]) == 0 || len(g[1]) == 0
}

func (g *game) score() int {
	for i, winner := range *g {
		if len(g[1-i]) == 0 {
			score := 0
			for i, x := range winner {
				score += x * (len(winner) - i)
			}
			return score
		}
	}
	panic("no winner yet")
}

func (g *game) same(other *game) bool {
	return reflect.DeepEqual(g, other)
}

func main() {
	var g game
	for i, para := range advent.InputParas() {
		g[i] = advent.Ints(para[1:])
	}

	// Part 1.
	{
		g := g
		for !g.done() {
			g.step()
		}

		fmt.Println("Part 1:", g.score())
	}

	// Part 2.
	{
		g := g
		g.play2()

		fmt.Println(g)

		fmt.Println("Part 2:", g.score())
	}
}

// returns winner index,
// or -1 if loop (i.e., player 1 wins by default)
func (g *game) play2() int {
	slow := *g
	even := false

	for !g.done() {
		g.step2()

		if even {
			slow.step2()
		}
		even = !even

		if g.same(&slow) {
			return 0
		}
	}

	for i := range g {
		if len(g[1-i]) == 0 {
			return i
		}
	}
	panic("bad")
}

func (g *game) step2() {
	i := g.round2()
	g[i] = g[i][:len(g[i]):len(g[i])]
	g[i] = append(g[i][1:], g[i][0], g[1-i][0])
	g[1-i] = g[1-i][1:]
}

func (g *game) round2() int {
	if len(g[0]) > g[0][0] && len(g[1]) > g[1][0] {
		recur := game{
			g[0][1 : 1+g[0][0]],
			g[1][1 : 1+g[1][0]],
		}
		return recur.play2()
	}
	for i := range g {
		if g[i][0] > g[1-i][0] {
			return i
		}
	}
	panic("bad")
}
