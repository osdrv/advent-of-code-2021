package main

import (
	"os"
)

func parseField(lines []string) [][]int {
	field := makeIntField(len(lines), len(lines[0]))
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			field[i][j] = int(lines[i][j] - '0')
		}
	}
	return field
}

// improvised heap
func partSort(q [][3]int) {
	minix := -1
	min := ALOT
	for i := 0; i < len(q); i++ {
		if q[i][2] < min {
			minix = i
			min = q[i][2]
		}
	}
	if minix > 0 {
		q[0], q[minix] = q[minix], q[0]
	}
}

func computeMinRisk(field [][]int) int {
	height, width := sizeIntField(field)

	q := make([][3]int, 0, 1)
	q = append(q, [3]int{0, 0, 0})

	minsum := ALOT
	visited := make(map[[2]int]int)
	var head [3]int
	for len(q) > 0 {
		head, q = q[0], q[1:]
		i, j, r := head[0], head[1], head[2]
		if i == height-1 && j == width-1 {
			if minsum > r {
				minsum = r
				continue
			}
		}
		ij := [2]int{i, j}
		if vv, ok := visited[ij]; ok {
			if vv <= r {
				continue
			}
		}
		visited[ij] = r
		if i < height-1 {
			q = append(q, [3]int{i + 1, j, r + field[i+1][j]})
		}
		if j < width-1 {
			q = append(q, [3]int{i, j + 1, r + field[i][j+1]})
		}
		if i > 0 {
			q = append(q, [3]int{i - 1, j, r + field[i-1][j]})
		}
		if j > 0 {
			q = append(q, [3]int{i, j - 1, r + field[i][j-1]})
		}
		partSort(q)
	}

	return minsum
}

func expandField(field [][]int, fact int) [][]int {
	height, width := sizeIntField(field)
	res := makeIntField(fact*height, fact*width)
	for ii := 0; ii < fact; ii++ {
		for jj := 0; jj < fact; jj++ {
			for i := 0; i < height; i++ {
				for j := 0; j < width; j++ {
					ni, nj := i+height*ii, j+width*jj
					nv := field[i][j]
					if ii > 0 {
						nv = res[ni-height][nj] + 1
					} else if jj > 0 {
						nv = res[ni][nj-width] + 1
					}
					if nv > 9 {
						nv = 1
					}
					res[ni][nj] = nv
				}
			}
		}
	}
	return res
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field := parseField(lines)

	print(printIntField(field, ""))

	risk := computeMinRisk(field)
	printf("the total risk is: %d", risk)

	expField := expandField(field, 5)
	risk2 := computeMinRisk(expField)
	printf("the total risk2 is: %d", risk2)
}
