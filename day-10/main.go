package main

import (
	"os"
	"sort"
)

var (
	PAIRING = map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	ERROR_SCORE = map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	AUTOCOMPLETE_SCORE = map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

func autoComplete(line string) []byte {
	err, q := findErr(line)
	if err > 0 {
		return nil
	}
	return reverseByteArr(mapByteArr(q, func(b byte) byte {
		return PAIRING[b]
	}))
}

func findErr(s string) (byte, []byte) {
	q := make([]byte, 0, 1)
	advance := func(b byte) bool {
		if q[len(q)-1] == b {
			q = q[:len(q)-1]
			return true
		}
		return false
	}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch ch {
		case '(', '[', '{', '<':
			q = append(q, ch)
		case ')', ']', '}', '>':
			if !advance(PAIRING[ch]) {
				return ch, nil
			}
		}
	}

	return 0, q
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	res := 0
	for _, line := range lines {
		if err, _ := findErr(line); err > 0 {
			printf("line: %s, err: %c", line, err)
			res += ERROR_SCORE[err]
		}
	}
	printf("error is: %d", res)

	scores := make([]int, 0, 1)
	for _, line := range lines {
		if auto := autoComplete(line); auto != nil {
			score := 0
			printf("line %s, auto: %s", line, string(auto))
			for _, b := range auto {
				score *= 5
				score += AUTOCOMPLETE_SCORE[b]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	printf("scores: %+v", scores)
	printf("res2: %d", scores[len(scores)/2])
}
