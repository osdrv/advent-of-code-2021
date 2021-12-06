package main

import (
	"log"
	"os"
)

func simulateFast(init []int, days int) int {
	var popcnt [9]int
	for _, v := range init {
		popcnt[v]++
	}

	for days > 0 {
		var newpopcnt [9]int
		for i := 0; i < len(popcnt); i++ {
			ix := i - 1
			if ix < 0 {
				ix = 6
			}
			newpopcnt[ix] += popcnt[i]
			if i == 0 {
				newpopcnt[8] += popcnt[i]
			}
		}
		popcnt = newpopcnt
		days--
	}

	res := 0
	for _, cnt := range popcnt {
		res += cnt
	}
	return res
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	nums := parseInts(readFile(f))

	log.Printf("nums: %+v", nums)

	res := simulateFast(nums, 256)

	log.Printf("the answer is: %d", res)
}
