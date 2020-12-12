package main

import (
	"fmt"
	"log"

	advent "github.com/mdempsky/advent2020"
)

const (
	Floor = '.'
	Empty = 'L'
	Occup = '#'
)

var state0 [1000][1000]byte
var state [1000][1000]byte

var rows, cols int

const part2 = true

func main() {
	lines := advent.InputLines()

	rows = 1 + len(lines) + 1
	cols = 1 + len(lines[0]) + 1

	for i := 0; i < cols; i++ {
		state[0][i] = Floor
		state[rows-1][i] = Floor
	}

	for i, line := range lines {
		if len(line) != len(lines[0]) {
			log.Fatal("huh?")
		}
		row := 1 + i
		state[row][0] = Floor
		copy(state[row][1:], line)
		state[row][cols-1] = Floor
	}

	for {
		changed := false

		state0 = state

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				old := state0[row][col]
				if old == Floor {
					continue
				}

				neighbors := 0
				for dr := -1; dr <= 1; dr++ {
				Neighbors:
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}
						r, c := row+dr, col+dc
						for r >= 0 && r < rows && c >= 0 && c < cols {

							nearby := state0[r][c]
							switch nearby {
							case Occup:
								neighbors++
								continue Neighbors
							case Empty:
								continue Neighbors
							}

							r, c = r+dr, c+dc
						}
					}
				}

				if old == Empty && neighbors == 0 {
					state[row][col] = Occup
					changed = true
				} else if old == Occup && neighbors >= 5 {
					state[row][col] = Empty
					changed = true
				}

			}
		}

		if !changed {
			break
		}
	}

	total := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if state[row][col] == Occup {
				total++
			}
		}
	}

	fmt.Println(total)
}
