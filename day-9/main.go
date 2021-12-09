package main

import (
	"log"
	"os"
	"sort"
)

var (
	STEPS = [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
)

func findLowPoints(field [][]int) []*Point2 {
	res := make([]*Point2, 0, 1)
	for i := 0; i < len(field); i++ {
	Element:
		for j := 0; j < len(field[0]); j++ {
		Step:
			for _, step := range STEPS {
				ni, nj := i+step[0], j+step[1]
				if ni < 0 || ni >= len(field) || nj < 0 || nj >= len(field[0]) {
					continue Step
				}
				if field[ni][nj] <= field[i][j] {
					continue Element
				}
			}
			res = append(res, NewPoint2(i, j))
		}
	}
	return res
}

func makeField(lines []string) [][]int {
	res := makeIntField(len(lines), len(lines[0]))
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			res[i][j] = int(lines[i][j] - '0')
		}
	}
	return res
}

func traverseBasin(field [][]int, point *Point2) int {
	var visit func(i, j int) int
	visited := make(map[[2]int]bool)
	visit = func(i, j int) int {
		if field[i][j] >= 9 {
			return 0
		}
		k := [2]int{i, j}
		if visited[k] {
			return 0
		}
		visited[k] = true
		res := 1
		for _, step := range STEPS {
			ni, nj := i+step[0], j+step[1]
			if ni < 0 || ni >= len(field) || nj < 0 || nj >= len(field[0]) {
				continue
			}
			if field[ni][nj] > field[i][j] {
				res += visit(ni, nj)
			}
		}
		return res
	}

	return visit(point.x, point.y)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field := makeField(lines)
	lows := findLowPoints(field)

	res1 := 0
	for _, point := range lows {
		res1 += 1 + field[point.x][point.y]
	}

	log.Printf("part1 result is: %d", res1)

	basins := make([]int, 0, 1)
	for _, point := range lows {
		basin := traverseBasin(field, point)
		basins = append(basins, basin)
		log.Printf("basin for point {%d, %d}: %d", point.x, point.y, basin)
	}

	sort.Ints(basins)
	ix := len(basins) - 1
	res2 := basins[ix] * basins[ix-1] * basins[ix-2]

	log.Printf("basins: %+v", basins)

	log.Printf("part2 result is: %d", res2)
}
