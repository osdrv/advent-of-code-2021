package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func noerr(err error) {
	if err != nil {
		panic(fmt.Sprintf("unhandled error: %s", err))
	}
}

func assert(check bool, msg string) {
	if !check {
		panic(fmt.Sprintf("assert %q failed", msg))
	}
}

func parseInt(s string) int {
	num, err := strconv.Atoi(s)
	noerr(err)
	return num
}

func readFile(in io.Reader) string {
	data, err := ioutil.ReadAll(in)
	noerr(err)
	return trim(string(data))
}

func readLines(in io.Reader) []string {
	scanner := bufio.NewScanner(in)
	lines := make([]string, 0, 1)
	for scanner.Scan() {
		lines = append(lines, trim(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("scan failed: %s", err))
	}
	return lines
}

func trim(s string) string {
	return strings.TrimRight(s, "\t\n\r")
}

func parseInts(s string) []int {
	chs := strings.FieldsFunc(trim(s), func(r rune) bool {
		return r == ' ' || r == ',' || r == '\t'
	})
	nums := make([]int, 0, len(chs))
	for i := 0; i < len(chs); i++ {
		nums = append(nums, parseInt(chs[i]))
	}
	return nums
}

type Point2 struct {
	x, y int
}

func NewPoint2(x, y int) *Point2 {
	return &Point2{x, y}
}

type Point3 struct {
	x, y, z int
}

func NewPoint3(x, y, z int) *Point3 {
	return &Point3{x, y, z}
}
