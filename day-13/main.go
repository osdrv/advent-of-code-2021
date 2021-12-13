package main

import (
	"os"
	"strings"
)

func parseFold(s string) *Point2 {
	x, y := 0, 0
	if strings.HasPrefix(s, "fold along x=") {
		x = parseInt(s[13:])
	} else {
		y = parseInt(s[13:])
	}
	return NewPoint2(x, y)
}

func foldField(field [][]int, fold *Point2) [][]int {
	if fold.x > 0 {
		return foldFieldVert(field, fold.x)
	} else {
		return foldFieldHor(field, fold.y)
	}
}

func foldFieldVert(field [][]int, x int) [][]int {
	printf("fold vertically along x=%d", x)
	folded := makeIntField(len(field), x)
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if j < x {
				folded[i][j] = field[i][j]
			} else if j > x {
				newj := x - (j - x)
				folded[i][newj] = min(field[i][j]+folded[i][newj], 1)
			}
		}
	}
	return folded
}

func foldFieldHor(field [][]int, y int) [][]int {
	printf("fold vertically along y=%d", y)
	folded := makeIntField(y, len(field[0]))
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if i < y {
				folded[i][j] = field[i][j]
			} else if i > y {
				newi := y - (i - y)
				folded[newi][j] = min(field[i][j]+folded[newi][j], 1)
			}
		}
	}
	return folded
}

func countDots(field [][]int) int {
	cnt := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if field[i][j] > 0 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	points := make([]*Point2, 0, 1)

	ix := 0
	maxx, maxy := 0, 0
	for {
		if lines[ix] == "" {
			break
		}
		nums := parseInts(lines[ix])
		point := NewPoint2(nums[0], nums[1])
		points = append(points, point)
		ix++
		if point.x > maxx {
			maxx = point.x
		}
		if point.y > maxy {
			maxy = point.y
		}
	}

	ix++
	folds := make([]*Point2, 0, 1)
	for ix < len(lines) {
		folds = append(folds, parseFold(lines[ix]))
		ix++
	}

	printf("points: %+v", points)
	printf("folds: %+v", folds)

	field := makeIntField(maxy+1, maxx+1)
	for _, point := range points {
		field[point.y][point.x] = 1
	}

	for ix, fold := range folds {
		field = foldField(field, fold)
		if ix == 0 {
			printf("number of dots after the first fold: %d", countDots(field))
		}
	}

	print(printIntFieldWithSubs(field, "", map[int]string{0: " ", 1: "#"}))
}
