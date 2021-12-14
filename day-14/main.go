package main

import (
	"os"
	"strings"
)

func strToPairs(s string) map[string]uint64 {
	chs := make(map[string]uint64)
	// we count the first char of the pair only, so trailing " " won't be counted
	s += " "
	for i := 0; i < len(s)-1; i++ {
		chs[s[i:i+2]]++
	}
	return chs
}

func countChars(pairs map[string]uint64) [26]uint64 {
	var cnts [26]uint64
	for ch, cnt := range pairs {
		cnts[ch[0]-'A'] += cnt
	}
	return cnts
}

func solve(pairs map[string]uint64, subs map[string]string) map[string]uint64 {
	newpairs := make(map[string]uint64)
	for s, cnt := range pairs {
		if ins, ok := subs[s]; ok {
			left, right := s[0:1]+ins, ins+s[1:]
			newpairs[left] += cnt
			newpairs[right] += cnt
		} else {
			newpairs[s] += uint64(cnt)
		}
	}
	return newpairs
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	tmpl := lines[0]

	subs := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		chs := strings.FieldsFunc(lines[i], func(r rune) bool {
			return r == ' ' || r == '-' || r == '>'
		})
		subs[chs[0]] = chs[1]
	}

	printf("tmpl: %s", tmpl)
	printf("subs: %+v", subs)

	steps := []int{10, 40}
	stepix := 0

	pairs := strToPairs(tmpl)
	for step := 1; step <= steps[len(steps)-1]; step++ {
		pairs = solve(pairs, subs)
		if step == steps[stepix] {
			cnts := countChars(pairs)
			low, high := uint64(999999999999999), uint64(0)
			for i := 0; i < len(cnts); i++ {
				if cnts[i] > 0 {
					if cnts[i] > high {
						high = cnts[i]
					}
					if low > cnts[i] {
						low = cnts[i]
					}
				}
			}
			printf("low: %d, high: %d", low, high)
			printf("the answer after %d steps is: %d", step, high-low)
			stepix++
		}
	}
}
