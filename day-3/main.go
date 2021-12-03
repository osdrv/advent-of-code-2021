package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func noerr(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Number struct {
	octs []byte
}

func NewNumber(s string) *Number {
	octs := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		octs[i] = s[i] - '0'
	}
	return &Number{
		octs: octs,
	}
}

func (n *Number) ByteAt(ix int) byte {
	return n.octs[ix]
}

func (n *Number) BSize() int {
	return len(n.octs)
}

func (n *Number) ToInt() int {
	res := 0
	for ix := 0; ix < len(n.octs); ix++ {
		if n.octs[len(n.octs)-ix-1] > 0 {
			res += 1 << ix
		}
	}
	return res
}

func (n *Number) Flip() *Number {
	octs := make([]byte, len(n.octs))
	for i := 0; i < len(n.octs); i++ {
		if n.octs[i] == 0 {
			octs[i] = 1
		}
	}
	return &Number{octs: octs}
}

func computeMCB(nums []*Number) *Number {
	cnt := make([]int, nums[0].BSize())
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(cnt); j++ {
			if nums[i].ByteAt(j) > 0 {
				cnt[j]++
			}
		}
	}
	octs := make([]byte, len(cnt))
	total := len(nums)
	for i := 0; i < len(cnt); i++ {
		if cnt[i] >= total-cnt[i] {
			octs[i] = 1
		}
	}
	return &Number{octs: octs}
}

func copyNums(nums []*Number) []*Number {
	cp := make([]*Number, len(nums))
	copy(cp, nums)
	return cp
}

func filterNums(nums []*Number, predicate func(*Number) bool) []*Number {
	res := make([]*Number, 0, 1)
	for _, num := range nums {
		if predicate(num) {
			res = append(res, num)
		}
	}
	return res
}

func computeGammaEpsilon(nums []*Number) (*Number, *Number) {
	mcb := computeMCB(nums)
	lcb := mcb.Flip()

	return mcb, lcb
}

func computeOxygenCo2(nums []*Number) (*Number, *Number) {
	oxynums := copyNums(nums)
	ix := 0
	for len(oxynums) > 1 {
		msb := computeMCB(oxynums)
		oxynums = filterNums(oxynums, func(num *Number) bool {
			return num.ByteAt(ix) == msb.ByteAt(ix)
		})
		ix++
	}
	co2s := copyNums(nums)
	ix = 0
	for len(co2s) > 1 {
		lsb := computeMCB(co2s).Flip()
		co2s = filterNums(co2s, func(num *Number) bool {
			return num.ByteAt(ix) == lsb.ByteAt(ix)
		})
		ix++
	}
	return oxynums[0], co2s[0]
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	nums := make([]*Number, 0, 1)
	for scanner.Scan() {
		nums = append(nums, NewNumber(strings.TrimRight(scanner.Text(), "\n\r\t")))
	}

	gamma, epsilon := computeGammaEpsilon(nums)
	log.Printf("gamma: %d, epsilon: %d, gamma * epsilon: %d", gamma.ToInt(), epsilon.ToInt(), gamma.ToInt()*epsilon.ToInt())

	oxygen, co2 := computeOxygenCo2(nums)
	log.Printf("oxygen: %d, co2: %d, oxygen * co2: %d", oxygen.ToInt(), co2.ToInt(), oxygen.ToInt()*co2.ToInt())
}
