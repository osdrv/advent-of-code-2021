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
	f, err := os.Open("INPUT")
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

func computeTravel(lines []string) (int, int) {
	x, y := 0, 0
	for _, line := range lines {
		if strings.HasPrefix(line, UP) {
			dy, err := strconv.Atoi(line[len(UP):])
			noerr(err)
			y -= dy
		} else if strings.HasPrefix(line, DOWN) {
			dy, err := strconv.Atoi(line[len(DOWN):])
			noerr(err)
			y += dy
		} else if strings.HasPrefix(line, FORWARD) {
			dx, err := strconv.Atoi(line[len(FORWARD):])
			noerr(err)
			x += dx
		}
	}
	return x, y
}

func computeWithAim(lines []string) (int, int) {
	h, d, a := 0, 0, 0
	for _, line := range lines {
		if strings.HasPrefix(line, UP) {
			da, err := strconv.Atoi(line[len(UP):])
			noerr(err)
			a -= da
		} else if strings.HasPrefix(line, DOWN) {
			da, err := strconv.Atoi(line[len(DOWN):])
			noerr(err)
			a += da
		} else if strings.HasPrefix(line, FORWARD) {
			dh, err := strconv.Atoi(line[len(FORWARD):])
			noerr(err)
			h += dh
			d += a * dh
		}
	}
	return h, d
}
