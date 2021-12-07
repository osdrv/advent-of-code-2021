package main

import (
	"log"
	"os"
)

func sumN(n int) int {
	return n * (n + 1) / 2
}

func computeMinAlignmentLog(nums []int) int {
	compute := func(p int) int {
		res := 0
		for _, num := range nums {
			//res += abs(p - num)
			res += sumN(abs(p - num))
		}
		return res
	}

	maxnum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxnum {
			maxnum = nums[i]
		}
	}

	return findMin(maxnum+1, compute)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	nums := parseInts(readFile(f))

	log.Printf("min fuel(log): %d", computeMinAlignmentLog(nums))
}
