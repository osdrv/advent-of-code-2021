package main

import (
	"os"
)

func parseField(lines []string) map[[2]int]bool {
	field := make(map[[2]int]bool)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			field[[2]int{i, j}] = lines[i][j] == '#'
		}
	}
	return field
}

func enhance(field map[[2]int]bool, key string, outer int) map[[2]int]bool {
	mini, minj, maxi, maxj := 0, 0, 0, 0
	for ij := range field {
		mini = min(mini, ij[0])
		minj = min(minj, ij[1])
		maxi = max(maxi, ij[0])
		maxj = max(maxj, ij[1])
	}

	newfield := make(map[[2]int]bool)
	for i := mini - 1; i <= maxi+1; i++ {
		for j := minj - 1; j <= maxj+1; j++ {
			newfield[[2]int{i, j}] = getPixel(field, i, j, key, outer)
		}
	}

	return newfield
}

func getPixel(field map[[2]int]bool, i, j int, key string, outer int) bool {
	ix := 0
	for _, step := range STEPS9 {
		ix <<= 1
		if v, ok := field[[2]int{i + step[0], j + step[1]}]; ok {
			if v {
				ix += 1
			}
		} else {
			if outer > 0 && key[0] == '#' {
				ix += 1
			}
		}
	}
	return key[ix] == '#'
}

func countPixels(field map[[2]int]bool) int {
	cnt := 0
	for _, px := range field {
		if px {
			cnt++
		}
	}
	return cnt
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	key := lines[0]

	lines = lines[2:]

	field := parseField(lines)

	printf("field: %+v", field)

	NUM_REPS := 50

	for i := 0; i < NUM_REPS; i++ {
		field = enhance(field, key, i%2)
	}

	cnt := countPixels(field)
	printf("pixels after %d reps: %d", NUM_REPS, cnt)
}
