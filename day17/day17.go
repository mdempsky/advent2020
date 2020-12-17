package main

import (
	"fmt"

	advent "github.com/mdempsky/advent2020"
)

const Cycles = 6
const N = Cycles + 8 + Cycles
const M = Cycles + 1 + Cycles

// Indexed: [x][y][z]
var old, new [N][N][M][M]bool

func getOld(x, y, z, w int) bool {
	if x < 0 || x >= N {
		return false
	}
	if y < 0 || y >= N {
		return false
	}
	if z < 0 || z >= M {
		return false
	}
	if w < 0 || w >= M {
		return false
	}
	return old[x][y][z][w]
}

func neighbors(x, y, z, w int) int {
	sum := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if (dx != 0 || dy != 0 || dz != 0 || dw != 0) && getOld(x+dx, y+dy, z+dz, w+dw) {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func main() {
	for y, line := range advent.InputLines() {
		for x, ch := range line {
			if ch == '#' {
				new[Cycles+x][Cycles+y][Cycles][Cycles] = true
			}
		}
	}

	for i := 0; i < Cycles; i++ {
		old = new
		for x := range new {
			for y := range new[x] {
				for z := range new[x][y] {
					for w := range new[x][y][z] {
						n := neighbors(x, y, z, w)
						if old[x][y][z][w] { // previously active
							if n < 2 || n > 3 {
								new[x][y][z][w] = false
							}
						} else { // previously inactive
							if n == 3 {
								new[x][y][z][w] = true
							}
						}
					}
				}
			}
		}
	}

	sum := 0
	for x := range new {
		for y := range new[x] {
			for z := range new[x][y] {
				for w := range new[x][y][z] {
					if new[x][y][z][w] {
						sum++
					}
				}
			}
		}
	}
	fmt.Println("Part 2:", sum)
}
