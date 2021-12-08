package main

import (
	"log"
	"os"
	"strings"
)

func strToBin(s string) uint8 {
	var res uint8
	for _, ch := range s {
		res |= 1 << (ch - 'a')
	}
	return res
}

func decodeLine(s string) [4]int {
	chs := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '|'
	})
	tail := make([]string, 4)
	copy(tail, chs[10:])

	digits := make(map[string]int)
	mapping := make(map[int]uint8)

	recognize := func(ch string, i int) {
		digits[ch] = i
		mapping[i] = strToBin(ch)
	}

	for _, ch := range chs {
		switch len(ch) {
		case 2:
			recognize(ch, 1)
		case 3:
			recognize(ch, 7)
		case 4:
			recognize(ch, 4)
		case 7:
			recognize(ch, 8)
		}
	}

	for _, ch := range chs {
		m := strToBin(ch)
		if len(ch) == 5 {
			//2, 3, 5
			if m&mapping[1] == mapping[1] {
				recognize(ch, 3)
			}
		}
		if len(ch) == 6 {
			// 0, 6, 9
			if m&mapping[7] != mapping[7] {
				recognize(ch, 6)
			} else if m&mapping[4] == mapping[4] {
				recognize(ch, 9)
			} else {
				recognize(ch, 0)
			}
		}
	}

	for _, ch := range chs {
		if len(ch) == 5 {
			m := strToBin(ch)
			if digits[ch] == 3 {
				continue
			}
			if mapping[9]&m == m {
				recognize(ch, 5)
			} else {
				recognize(ch, 2)
			}
		}
	}

	return [4]int{
		digits[tail[0]],
		digits[tail[1]],
		digits[tail[2]],
		digits[tail[3]],
	}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	// part1
	cnt := 0
	for _, line := range lines {
		d4 := decodeLine(line)
		for _, d := range d4 {
			switch d {
			case 1, 4, 7, 8:
				cnt++
			}
		}
	}
	log.Printf("part1 result is: %d", cnt)

	// part2
	sum := 0
	for _, line := range lines {
		d4 := decodeLine(line)
		log.Printf("decoded: %+v", d4)
		sum += d4[0]*1000 + d4[1]*100 + d4[2]*10 + d4[3]
	}
	log.Printf("part2 result is: %d", sum)
}
