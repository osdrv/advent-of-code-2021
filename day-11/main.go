package main

import (
	"os"
)

func simulate(field [][]int) int {
	flashed := make(map[[2]int]bool)
	rows, cols := sizeIntField(field)
	var visit func(i, j int)
	visit = func(i, j int) {
		field[i][j]++
		if field[i][j] > 9 {
			//flash
			if !flashed[[2]int{i, j}] {
				flashed[[2]int{i, j}] = true
				for _, step := range STEPS8 {
					ni, nj := i+step[0], j+step[1]
					if ni < 0 || ni >= rows || nj < 0 || nj >= cols {
						continue
					}
					visit(ni, nj)
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			visit(i, j)
		}
	}

	flashes := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if field[i][j] > 9 {
				field[i][j] = 0
				flashes++
			}
		}
	}
	return flashes
}

func parseField(lines []string) [][]int {
	field := makeIntField(len(lines), len(lines[0]))
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			field[i][j] = int(lines[i][j] - '0')
		}
	}
	return field
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	var NUM_STEPS = 100

	flashes := 0
	field := parseField(lines)
	print(printIntField(field, ""))
	step := 0
	ALL := len(field) * len(field[0])
	for {
		step++
		f := simulate(field)
		flashes += f
		if step == NUM_STEPS {
			printf("flashes after %d steps: %d", flashes, NUM_STEPS)
		}
		if f == ALL {
			printf("all octs flashed at step %d", step)
			break
		}
	}
}
