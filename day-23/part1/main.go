package main

import (
	"bytes"
)

var (
	CAP = [11]int{
		1, 1, 3, 1, 3, 1, 3, 1, 3, 1, 1,
	}

	COST = map[byte]int{
		'A': 1,
		'B': 10,
		'C': 100,
		'D': 1000,
	}

	HOME = map[byte]int{
		'A': 2,
		'B': 4,
		'C': 6,
		'D': 8,
	}
)

type Game struct {
	pegs [11][]byte
}

func NewGame(init ...string) *Game {
	var pegs [11][]byte
	for ix, cc := range CAP {
		pegs[ix] = make([]byte, cc)
	}
	for i, in := range init {
		for j := 0; j < len(in); j++ {
			pegs[2+2*i][j] = in[j]
		}
	}
	return &Game{
		pegs: pegs,
	}
}

func copyPegs(pegs [11][]byte) [11][]byte {
	var res [11][]byte
	for ix, peg := range pegs {
		res[ix] = make([]byte, len(peg))
		copy(res[ix], peg)
	}
	return res
}

type state struct {
	pegs  [11][]byte
	score int
	dist  int
}

/*

01
10
11

A 100
B 101
C 110
D 111

xxx xxx xxxxxx xxx xxxxxx xxx xxxxxx xxx xxxxxx xxx xxx

*/

func hashCodeStr(pegs [11][]byte) string {
	var buf bytes.Buffer
	for _, peg := range pegs {
		for _, v := range peg {
			if v == 0 {
				buf.WriteByte(' ')
			} else {
				buf.WriteByte(v)
			}
		}
	}
	return buf.String()
}

func hashCode(pegs [11][]byte) uint64 {
	var res uint64
	for i := 0; i < len(pegs); i++ {
		res <<= 3
		if pegs[i][0] > 0 {
			res |= uint64(0b100 | (pegs[i][0] - 'A'))
		}
		if len(pegs[i]) > 1 {
			res <<= 3
			if pegs[i][1] > 0 {
				res |= uint64(0b100 | (pegs[i][1] - 'A'))
			}
		}
	}
	return res
}

// just kidding, simply bring the next best option to the top in O(n)
func heapify(states []state) {
	minIx := 0
	minDist := states[0].dist
	// limit this to a 1024 steps to make sure we're not spending too much time here
	// any "decent" option is ok
	for i := 1; i < min(len(states), 1024); i++ {
		if states[i].dist < minDist {
			minDist = states[i].dist
			minIx = i
		}
	}
	if minIx > 0 {
		states[0], states[minIx] = states[minIx], states[0]
	}
}

var (
	stopAt = map[string]int{
		//"  AB BDC  C   AD   ": 40,                                          // v
		//"  AB BD   CC  AD   ": 40 + 400,                                    // v
		//"  AB  B  DCC  AD   ": 40 + 400 + 3000 + 30,                        // v
		//"  A   BB DCC  AD   ": 40 + 400 + 3000 + 30 + 40,                   // v
		//"  A   BB DCC D   A ": 40 + 400 + 3000 + 30 + 40 + 2003,            // v
		//"  A   BB  CC  DD A ": 40 + 400 + 3000 + 30 + 40 + 2003 + 7000,     // v
		//"  AA  BB  CC  DD   ": 40 + 400 + 3000 + 30 + 40 + 2003 + 7000 + 8, // v
	}
)

func (g *Game) Play() int {
	score := ALOT
	states := []state{
		{pegs: g.pegs, score: 0, dist: completionDist(g.pegs)},
	}
	var ss state
	memo := make(map[string]int)
	for len(states) > 0 {
		// boost it up!
		heapify(states)
		ss, states = states[0], states[1:]
		//printf("step score: %d", ss.score)
		hc := hashCodeStr(ss.pegs)

		if v, ok := memo[hc]; ok {
			// fuck me and my life, it was >= here
			if v <= ss.score {
				//printf("we've been there with a better score")
				continue
			}
		}

		memo[hc] = ss.score

		//printf("completion distance: %d", ss.dist)
		//printf("signature: %064b", hc)
		//printf("signature: %q", hc)
		//print(printPegs(ss.pegs))

		//if stopAt[hc] > 0 {
		//	printf("expected score: %d", stopAt[hc])
		//	runtime.Breakpoint()
		//}

		//runtime.Breakpoint()
		if isComplete(ss.pegs) {
			//runtime.Breakpoint()
			printf("==================Complete===================")
			if ss.score < score {
				score = ss.score
			}
			//break
		}
		if ss.score > score {
			//printf("***** skipping as there is a better score *****")
			// no need to check a more expensive option if there is already a solution
			continue
		}
		for _, next := range computeSteps(ss.pegs) {
			next.score += ss.score // score so far
			hhc := hashCodeStr(next.pegs)
			if v, ok := memo[hhc]; ok {
				if v <= next.score {
					// we've tried this combination with a better score
					continue
				}
			}
			next.dist = completionDist(next.pegs)
			states = append(states, next)
		}
	}
	return score
}

func checkIntegrity(pegs [11][]byte) bool {
	var cnt [4]int
	for _, peg := range pegs {
		for _, pp := range peg {
			if pp > 0 {
				cnt[pp-'A']++
			}
		}
	}
	for _, cc := range cnt {
		if cc != 2 {
			return false
		}
	}
	return true
}

func correctPad(pegs [11][]byte) bool {
	for _, peg := range pegs {
		zero := false
		for j := 0; j < len(peg); j++ {
			if peg[j] == 0 {
				zero = true
				continue
			}
			if zero {
				// wrong padding
				return false
			}
		}
	}
	return true
}

func computeSteps(pegs [11][]byte) []state {
	res := make([]state, 0, 1)
	for i := 0; i < len(pegs); i++ {
		head, off := getHead(pegs[i])
		if head == 0 {
			continue
		}
		if HOME[head] == i {
			// check if we should move it
			if off == 2 || off == 1 && pegs[i][0] == pegs[i][1] {
				// if the element is in its room at the bottom
				// or the room contains the same amphipods, no need to move
				continue
			}
		}
		//printf("peg %d, head: %c(ix:%d)", i, head, off)
		//if off == 2 {
		//	runtime.Breakpoint()
		//}

		minj, maxj := i, i
		for j := i - 1; j >= 0; j-- {
			_, oo := getHead(pegs[j])
			if oo == 0 {
				minj = j + 1
				break
			}
			minj = j
		}
		for j := i + 1; j < len(pegs); j++ {
			_, oo := getHead(pegs[j])
			if oo == 0 {
				maxj = j - 1
				break
			}
			maxj = j
		}
		//printf("minj: %d, maxj: %d", minj, maxj)

		for j := minj; j <= maxj; j++ {
			if j == i {
				continue
			}
			hh, oo := getHead(pegs[j])
			switch oo {
			case -1:
				// this space is free
				if CAP[j] > 1 {
					// this is an empty room
					// need to check if this amphipod belongs here
					if HOME[head] != j {
						continue
					}
				}
				cp := copyPegs(pegs)
				cp[j][0] = head
				cp[i][CAP[i]-1-off] = 0
				score := COST[head] * (off + abs(i-j) + CAP[j] - 1)
				//if !correctPad(cp) {
				//	printf("wrong padding")
				//	runtime.Breakpoint()
				//}
				//if !checkIntegrity(cp) {
				//	printf("!!!!! incorrect pegs")
				//	runtime.Breakpoint()
				//}
				res = append(res, state{pegs: cp, score: score})
			case 0:
				// it is a hallway and it is blocked
				// there is no way we can move any further
				break
			case 1:
				continue // the peg is full
			case 2:
				// this is a room with 1 amphipod in it
				if HOME[head] != j {
					// this amphipod does not belong here
					continue
				}
				if hh != head {
					continue
				}
				// this is the home and there is already a similar amphipod
				cp := copyPegs(pegs)
				cp[j][1] = head
				cp[i][CAP[i]-1-off] = 0
				score := COST[head] * (off + abs(i-j) + 1)
				//if !correctPad(cp) {
				//	printf("wrong padding")
				//	runtime.Breakpoint()
				//}
				//if !checkIntegrity(cp) {
				//	printf("!!!!! incorrect pegs")
				//	runtime.Breakpoint()
				//}
				res = append(res, state{pegs: cp, score: score})
			}
		}
	}
	return res
}

func getHead(peg []byte) (byte, int) {
	for i := len(peg) - 1; i >= 0; i-- {
		if peg[i] > 0 {
			return peg[i], len(peg) - 1 - i
		}
	}
	return 0, -1
}

func completionDist(pegs [11][]byte) int {
	res := 0
	for i := 0; i < len(pegs); i++ {
		for j := 0; j < len(pegs[i]); j++ {
			if pegs[i][j] > 0 {
				res += abs(HOME[pegs[i][j]] - i)
			}
		}
	}
	return res
}

func isComplete(pegs [11][]byte) bool {
	return completionDist(pegs) == 0
}

func printPegs(pegs [11][]byte) string {
	bb := makeByteField(3, 11)
	for j, peg := range pegs {
		for i := 0; i < len(peg); i++ {
			if peg[i] > 0 {
				bb[CAP[j]-i-1][j] = peg[i] - '0'
			} else {
				bb[CAP[j]-i-1][j] = 1
			}
		}
	}
	return printByteFieldWithSubs(bb, "", map[byte]string{
		0: " ",
		1: ".",
	})
}

func main() {
	//game := NewGame("AB", "DC", "CB", "AD")
	game := NewGame("CD", "CB", "AD", "BA")
	res := game.Play()
	printf("part 1 result is: %d", res)
}
