package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func noerr(err error) {
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %s", err))
	}
}

func countNSums(nums []int, n int) int {
	prev := 0
	for i := 0; i < n; i++ {
		prev += nums[i]
	}
	inc := 0
	for j := n; j < len(nums); j++ {
		next := nums[j] + prev - nums[j-n]
		if next > prev {
			inc++
		}
		prev = next
	}
	return inc
}

func readInts(f io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(f)
	nums := make([]int, 0, 1)
	for scanner.Scan() {
		n, err := strconv.Atoi(strings.Trim(scanner.Text(), "\n\t\r"))
		noerr(err)
		nums = append(nums, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nums, nil
}

func main() {
	f, err := os.Open("INPUT-TST")
	noerr(err)
	defer f.Close()
	nums, err := readInts(f)
	noerr(err)

	log.Printf("%+v", nums)

	log.Printf("1-increases: %d", countNSums(nums, 1))
	log.Printf("3-increases: %d", countNSums(nums, 3))
}
