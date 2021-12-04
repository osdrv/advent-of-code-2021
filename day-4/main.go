package main

import (
	"bufio"
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

const (
	S_UINT8 = 8
)

type Board struct {
	numbers [][]int
	ixs     map[int][2]int
	cross   []uint8
}

func NewBoard(lines []string) *Board {
	numbers := make([][]int, len(lines))
	ixs := make(map[int][2]int)
	cross := make([]uint8, len(lines))
	for i := 0; i < len(lines); i++ {
		numbers[i] = parseNumbers(lines[i])
		if len(numbers[i]) > S_UINT8 {
			panic("did not expect that")
		}
		for j := 0; j < len(numbers[i]); j++ {
			ixs[numbers[i][j]] = [2]int{i, j}
		}
	}
	return &Board{
		numbers: numbers,
		ixs:     ixs,
		cross:   cross,
	}
}

func (b *Board) Play(num int) {
	if ix, ok := b.ixs[num]; ok {
		i, j := ix[0], ix[1]
		b.cross[i] |= (1 << j)
	}
}

func (b *Board) IsWin() bool {
	var msk uint8 = (1 << len(b.numbers[0])) - 1
	and := msk
	for _, cc := range b.cross {
		if cc == msk {
			return true
		}
		and &= cc
	}
	return and > 0
}

func (b *Board) GetUnmarked() []int {
	nums := make([]int, 0, 1)
	for i := 0; i < len(b.numbers); i++ {
		for j := 0; j < len(b.numbers[i]); j++ {
			if b.cross[i]&(1<<j) == 0 {
				nums = append(nums, b.numbers[i][j])
			}
		}
	}
	return nums
}

func fields(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == ','
	})
}

func parseNumbers(s string) []int {
	nums := make([]int, 0, 1)
	for _, ch := range fields(s) {
		num, err := strconv.Atoi(ch)
		noerr(err)
		nums = append(nums, num)
	}
	return nums
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	numbers := parseNumbers(scanner.Text())
	log.Printf("Numbers: %+v", numbers)

	scanner.Scan()

	lines := make([]string, 0, 1)
	boards := make([]*Board, 0, 1)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			board := NewBoard(lines)
			boards = append(boards, board)
			lines = make([]string, 0, 1)
			continue
		}
		lines = append(lines, line)
	}
	log.Printf("boards: %+v", boards)

	numix := 0
	for len(boards) > 0 && numix < len(numbers) {
		next := make([]*Board, 0, 1)
		for _, bb := range boards {
			num := numbers[numix]
			bb.Play(num)
			if bb.IsWin() {
				log.Printf("board: %+v, wining number: %d", bb, num)
				sum := 0
				for _, n := range bb.GetUnmarked() {
					sum += n
				}
				log.Printf("The answer is: %d * %d = %d", sum, num, sum*num)
			} else {
				next = append(next, bb)
			}
		}
		boards = next
		numix++
	}
}
