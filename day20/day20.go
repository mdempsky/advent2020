package main

import (
	"fmt"
	"math/bits"

	advent "github.com/mdempsky/advent2020"
)

type tile struct {
	id   int
	data [10][10]bool

	n, e, s, w *edge
}

type edge struct {
	t1, t2 *tile
	match  bool
}

func (e *edge) flip() { e.match = !e.match }

func (e *edge) other(t *tile) *tile {
	switch t {
	case e.t1:
		return e.t2
	case e.t2:
		return e.t1
	}
	panic(e)
}

func main() {
	var tiles []*tile
	for _, para := range advent.InputParas() {
		var t tile
		advent.Sscanf(para[0], "Tile %d:", &t.id)
		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				t.data[r][c] = para[1+r][c] == '#'
			}
		}
		tiles = append(tiles, &t)
	}

	edges := map[uint16]*edge{}
	add := func(x uint16, t *tile, ep **edge) {
		flip := false
		if r := bits.Reverse16(x) >> 6; r < x {
			x = r
			flip = true
		}

		e, ok := edges[x]
		if ok {
			if e.t2 != nil {
				panic("non-unique edge")
			}
			e.t2 = t
		} else {
			e = &edge{t1: t}
			edges[x] = e
		}

		*ep = e
		if flip {
			e.flip()
		}
	}

	for _, t := range tiles {
		var N, E, S, W uint16
		for i := 0; i < 10; i++ {
			N = N<<1 | v(t.data[0][i])
			E = E<<1 | v(t.data[i][9])
			S = S<<1 | v(t.data[9][9-i])
			W = W<<1 | v(t.data[9-i][0])
		}
		add(N, t, &t.n)
		add(E, t, &t.e)
		add(S, t, &t.s)
		add(W, t, &t.w)
	}

	corners := 1
	var first *tile

	for _, t := range tiles {
		if t.neighbors() == 2 {
			corners *= t.id
			if first == nil {
				first = t
			}
		}
	}
	fmt.Println("Part 1:", corners)

	for first.e.other(first) == nil || first.s.other(first) == nil {
		first.rotateLeft()
	}

	var rows [][]*tile
	for this := first; ; {
		var row []*tile
		for this := this; ; {
			row = append(row, this)
			right := this.e.other(this)
			if right == nil {
				break
			}
			for right.w != this.e {
				right.rotateLeft()
			}
			if !this.e.match {
				right.flipV()
			}
			this = right
		}
		rows = append(rows, row)

		down := this.s.other(this)
		if down == nil {
			break
		}
		for down.e != this.s {
			down.rotateLeft()
		}
		if !this.s.match {
			down.flipV()
		}
		down.rotateLeft() // down.e becomes down.n, matching this.s
		this = down
	}

	if false {
		for _, row := range rows {
			for r := 0; r < 10; r++ {
				for _, t := range row {
					for c := 0; c < 10; c++ {
						x := '.'
						if t.data[r][c] {
							x = '#'
						}
						fmt.Print(string(x))
					}
					fmt.Print(" ")
				}
				fmt.Println()
			}
			fmt.Println()
		}
	}

	var img image
	var total int
	for r := 0; r < 96; r++ {
		for c := 0; c < 96; c++ {
			if rows[r/8][c/8].data[1+(r%8)][1+(c%8)] {
				img.data[r][c] = true
				total++
			}
		}
	}

	found := 0
	for try := 0; try < 8; try++ {
		for r := 0; r < 96; r++ {
			for c := 0; c < 96; c++ {
				if img.monsterAt(r, c) {
					found++
				}
			}
		}
		if found != 0 {
			break
		}

		img.rotateLeft()
		if try == 4 {
			img.flipV()
		}
	}
	if found == 0 {
		panic("no sea monsters found")
	}
	fmt.Println("Part 2:", total-found*len(monster))
}

type image struct {
	data [96][96]bool
}

func (i *image) rotateLeft() {
	new, old := &i.data, i.data
	for r := 0; r < 96; r++ {
		for c := 0; c < 96; c++ {
			new[r][c] = old[c][95-r]
		}
	}
}

func (i *image) flipV() {
	new, old := &i.data, i.data
	for r := 0; r < 96; r++ {
		for c := 0; c < 96; c++ {
			new[r][c] = old[95-r][c]
		}
	}
}

//                   #
// #    ##    ##    ###
//  #  #  #  #  #  #
var monster = [...][2]int{
	{1, 0},
	{2, 1},
	{2, 4},
	{1, 5},
	{1, 6},
	{2, 7},
	{2, 10},
	{1, 11},
	{1, 12},
	{2, 13},
	{2, 16},
	{1, 17},
	{0, 18},
	{1, 18},
	{1, 19},
}

func (i *image) monsterAt(r, c int) bool {
	if r+2 >= len(i.data) || c+19 >= len(i.data[0]) {
		return false
	}
	for _, p := range &monster {
		if !i.data[r+p[0]][c+p[1]] {
			return false
		}
	}
	return true
}

func (t *tile) rotateLeft() {
	new, old := &t.data, t.data
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			new[r][c] = old[c][9-r]
		}
	}

	t.n, t.e, t.s, t.w = t.e, t.s, t.w, t.n
}

func (t *tile) flipV() {
	new, old := &t.data, t.data
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			new[r][c] = old[9-r][c]
		}
	}

	t.n, t.s = t.s, t.n
	for _, e := range t.edges() {
		e.flip()
	}
}

func (t *tile) neighbors() int {
	sum := 0
	for _, e := range t.edges() {
		if e.t2 != nil {
			sum++
		}
	}
	return sum
}

func (t *tile) edges() [4]*edge {
	return [4]*edge{t.n, t.e, t.s, t.w}
}

func v(b bool) uint16 {
	if b {
		return 1
	}
	return 0
}
