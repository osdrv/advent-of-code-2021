package main

import (
	"fmt"
	"os"
)

type Mode uint8

const (
	LEFT Mode = iota
	RIGHT
)

type Number struct {
	v           int
	left, right *Number
}

func (n *Number) String() string {
	if n.isRegular() {
		return fmt.Sprintf("%d", n.v)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func (n *Number) Add(n2 *Number) *Number {
	return &Number{
		left:  n,
		right: n2,
	}
}

func (n *Number) isRegular() bool {
	return n.left == nil || n.right == nil
}

func (n *Number) Explode() bool {
	stack := make([]*Number, 0, 1)
	stack = append(stack, n)
	var traverse func(*Number, int) []*Number
	traverse = func(n *Number, depth int) []*Number {
		if n.isRegular() {
			return nil
		}
		if depth == 4 {
			if n.left.isRegular() && n.right.isRegular() {
				return []*Number{n}
			}
			return nil
		}
		if n.left != nil {
			if tr := traverse(n.left, depth+1); tr != nil {
				return append([]*Number{n}, tr...)
			}
		}
		if n.right != nil {
			if tr := traverse(n.right, depth+1); tr != nil {
				return append([]*Number{n}, tr...)
			}
		}
		return nil
	}
	tr := traverse(n, 0)
	if tr == nil {
		return false
	}

	printf("traverse complete: %+v", tr)

	add(tr, LEFT)
	add(tr, RIGHT)

	printf("addition complete")

	repl := &Number{v: 0}
	if tr[len(tr)-2].left == tr[len(tr)-1] {
		tr[len(tr)-2].left = repl
	} else {
		tr[len(tr)-2].right = repl
	}

	return true
}

func (n *Number) Split() bool {
	if n.isRegular() {
		if n.v >= 10 {
			left, right := n.v/2, n.v-n.v/2
			n.left = &Number{v: left}
			n.right = &Number{v: right}
			n.v = 0
			return true
		}
		return false
	}
	return n.left.Split() || n.right.Split()
}

func (n *Number) Reduce() bool {
	return n.Explode() || n.Split()
}

func add(tr []*Number, mode Mode) {
	ptr := len(tr) - 1
	var toAdd int
	if mode == LEFT {
		toAdd = tr[ptr].left.v
	} else {
		toAdd = tr[ptr].right.v
	}
	printf("add number: mode: %d, toadd: %d", mode, toAdd)
	ptr--
	for ptr >= 0 {
		if mode == LEFT {
			if tr[ptr].left == tr[ptr+1] {
				ptr--
				continue
			}
			addTo(tr[ptr].left, toAdd, RIGHT)
			return
		} else {
			if tr[ptr].right == tr[ptr+1] {
				ptr--
				continue
			}
			addTo(tr[ptr].right, toAdd, LEFT)
			return
		}
	}
}

func addTo(n *Number, v int, mode Mode) {
	if n.isRegular() {
		n.v += v
		return
	}
	if mode == LEFT {
		if n.left != nil {
			addTo(n.left, v, mode)
		} else {
			addTo(n.right, v, mode)
		}
	} else {
		if n.right != nil {
			addTo(n.right, v, mode)
		} else {
			addTo(n.left, v, mode)
		}
	}
}

func parseNumber(s string, ptr int) (*Number, int) {
	if ptr >= len(s) {
		return nil, ptr
	}
	if isDigit(s[ptr]) {
		return parseRegular(s, ptr)
	}
	var left, right *Number
	ptr = read(s, ptr, '[')
	left, ptr = parseNumber(s, ptr)
	ptr = read(s, ptr, ',')
	right, ptr = parseNumber(s, ptr)
	ptr = read(s, ptr, ']')
	return &Number{left: left, right: right}, ptr
}

func read(s string, ptr int, ch byte) int {
	if s[ptr] != ch {
		fatalf("failed to read ch %c at pos %d from str %s", ch, ptr, s)
	}
	return ptr + 1
}

func parseRegular(s string, ptr int) (*Number, int) {
	off := ptr
	for off < len(s) && isDigit(s[off]) {
		off++
	}
	num := parseInt(s[ptr:off])
	return &Number{v: num}, off
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func magnitude(num *Number) int {
	if num.isRegular() {
		return num.v
	}
	return 3*magnitude(num.left) + 2*magnitude(num.right)
}

func reduce(num *Number) {
	printf("test num before reduce: %s", num)
	for num.Reduce() {
		printf("number after reduction: %s", num)
	}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	nums := make([]*Number, 0, 1)
	lines := readLines(f)

	for _, line := range lines {
		num, _ := parseNumber(line, 0)
		nums = append(nums, num)
		printf("num: %s", num)
	}

	for len(nums) > 1 {
		reduce(nums[0])
		reduce(nums[1])
		nn := nums[0].Add(nums[1])
		reduce(nn)
		nums = append([]*Number{nn}, nums[2:]...)
	}

	printf("final number: %s", nums[0])
	printf("magnitude: %d", magnitude(nums[0]))

	maxMagn := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			num1, _ := parseNumber(lines[i], 0)
			num2, _ := parseNumber(lines[j], 0)
			reduce(num1)
			reduce(num2)
			nn := num1.Add(num2)
			reduce(nn)
			magn := magnitude(nn)
			if magn > maxMagn {
				maxMagn = magn
			}
		}
	}

	printf("max magnitude: %d", maxMagn)
}
