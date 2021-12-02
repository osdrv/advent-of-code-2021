package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	FORWARD = "forward "
	DOWN    = "down "
	UP      = "up "
)

func noerr(err error) {
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
}

func main() {
	f, err := os.Open("INPUT-TST")
	noerr(err)
	fmt.Println("vim-go")
	scanner := bufio.NewScanner(f)
	lines := make([]string, 0, 1)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\n\r\t")
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("failed to scan: %s", err))
	}
	x, y := computeTravel(lines)
	log.Printf("x: %d, y: %d, x*y: %d", x, y, x*y)

	x, y = computeWithAim(lines)
	log.Printf("with aim: a: %d, y: %d, x*y: %d", x, y, x*y)
}

func parseInt(s string) int {
	v, err := strconv.Atoi(s)
	noerr(err)
	return v
}

func startsWith(line, pref string) (int, bool) {
	if strings.HasPrefix(line, pref) {
		v, err := strconv.Atoi(line[len(pref):])
		noerr(err)
		return v, true
	}
	return 0, false
}

func computeTravel(lines []string) (int, int) {
	x, y := 0, 0
	for _, line := range lines {
		if dy, ok := startsWith(line, UP); ok {
			y -= dy
		} else if dy, ok := startsWith(line, DOWN); ok {
			y += dy
		} else if dx, ok := startsWith(line, FORWARD); ok {
			x += dx
		} else {
			panic(fmt.Sprintf("unknown move: %s", line))
		}
	}
	return x, y
}

func computeWithAim(lines []string) (int, int) {
	h, d, a := 0, 0, 0
	for _, line := range lines {
		if da, ok := startsWith(line, UP); ok {
			a -= da
		} else if da, ok := startsWith(line, DOWN); ok {
			a += da
		} else if dh, ok := startsWith(line, FORWARD); ok {
			h += dh
			d += a * dh
		} else {
			panic(fmt.Sprintf("unknown move: %s", line))
		}
	}
	return h, d
}
