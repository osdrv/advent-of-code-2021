package main

import (
	"os"
)

const (
	EAST  = '>'
	SOUTH = 'v'
	EMPTY = '.'
)

func evolve(field [][]byte) ([][]byte, int) {
	height, width := sizeByteField(field)
	cp := makeByteField(height, width)
	moves := 0

	// move east first
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if field[i][j] != EAST {
				continue
			}
			ni, nj := i, j+1
			if nj == width {
				nj = 0
			}
			if field[ni][nj] == 0 {
				cp[ni][nj] = field[i][j]
				cp[i][j] = 0
				moves++
			} else {
				cp[i][j] = field[i][j]
			}
		}
	}

	// move south then
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if field[i][j] != SOUTH {
				continue
			}
			ni, nj := i+1, j
			if ni == height {
				ni = 0
			}
			// if it is empty at first or there was an EAST-facing cucumber that has moved
			if (field[ni][nj] == 0 && cp[ni][nj] == 0) || (field[ni][nj] == EAST && cp[ni][nj] == 0) {
				cp[ni][nj] = field[i][j]
				cp[i][j] = 0
				moves++
			} else {
				cp[i][j] = field[i][j]
			}
		}
	}

	return cp, moves
}

func parseField(lines []string) [][]byte {
	field := makeByteField(len(lines), len(lines[0]))
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] != EMPTY {
				field[i][j] = lines[i][j]
			}
		}
	}
	return field
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field := parseField(lines)
	print(printByteFieldWithSubs(field, "", map[byte]string{SOUTH: "v", EAST: ">", 0: "."}))
	moves := 0
	step := 0
	for {
		field, moves = evolve(field)
		step++
		//print(printByteFieldWithSubs(field, "", map[byte]string{SOUTH: "v", EAST: ">", 0: "."}))
		if moves == 0 {
			printf("stopped at step %d", step)
			break
		}
	}
}
