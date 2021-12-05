package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func noerr(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Pixel struct {
	x, y int
}

func NewPixel(x, y int) *Pixel {
	return &Pixel{x: x, y: y}
}

type Segment struct {
	p1, p2 *Pixel
}

func NewSegment(s string) *Segment {
	chs := strings.FieldsFunc(s, func(r rune) bool {
		return r == ',' || r == ' ' || r == '-' || r == '>'
	})

	x1, err := strconv.Atoi(chs[0])
	noerr(err)
	y1, err := strconv.Atoi(chs[1])
	noerr(err)
	x2, err := strconv.Atoi(chs[2])
	noerr(err)
	y2, err := strconv.Atoi(chs[3])
	noerr(err)

	return &Segment{
		p1: NewPixel(x1, y1),
		p2: NewPixel(x2, y2),
	}
}

func (l *Segment) String() string {
	return fmt.Sprintf("[{%d, %d} -> {%d, %d}]", l.p1.x, l.p1.y, l.p2.x, l.p2.y)
}

func (l *Segment) Raster() []*Pixel {
	res := make([]*Pixel, 0, 1)
	dx, dy := sign(l.p2.x-l.p1.x), sign(l.p2.y-l.p1.y)
	x, y := l.p1.x, l.p1.y
	for x != l.p2.x || y != l.p2.y {
		res = append(res, NewPixel(x, y))
		x += dx
		y += dy
	}
	res = append(res, NewPixel(x, y))
	return res
}

func sign(v int) int {
	if v < 0 {
		return -1
	} else if v == 0 {
		return 0
	}
	return 1
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	segments := make([]*Segment, 0, 1)
	for scanner.Scan() {
		segment := NewSegment(strings.TrimRight(scanner.Text(), "\r\n\t"))
		segments = append(segments, segment)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan failed: %s", err)
	}

	log.Printf("segments: %+v", segments)

	pCnt := make(map[Pixel]int)

	for _, segment := range segments {
		for _, p := range segment.Raster() {
			pCnt[*p]++
		}
	}

	cnt := 0
	for p, n := range pCnt {
		if n >= 2 {
			log.Printf("Pixel: {%d, %d}", p.x, p.y)
			cnt++
		}
	}

	log.Printf("The answer is: %d", cnt)
}
