package main

func play(p1, p2 int) (int, int, int) {
	players := [2]int{p1 - 1, p2 - 1}
	scores := [2]int{0, 0}
	rolls := 0

	d := 0
	rollDice := func() int {
		d++
		return d
	}

Game:
	for {
		for ix, pp := range players {
			sc := 0
			for i := 0; i < 3; i++ {
				sc += rollDice()
				rolls++
			}
			pp += sc
			if pp >= 10 {
				pp = pp % 10
			}
			players[ix] = pp
			scores[ix] += (pp + 1)
			if scores[ix] >= 1000 {
				break Game
			}
		}

	}

	return scores[0], scores[1], rolls
}

var (
	CNTS = map[int]uint64{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
)

func play2(p1, p2 int, s1, s2 int, first bool) (uint64, uint64) {
	// player rolls 3x1, 4x3, 5x6, 6x7, 7x6, 8x3, 9x1
	if s1 >= 21 {
		return 1, 0
	} else if s2 >= 21 {
		return 0, 1
	}

	r1, r2 := uint64(0), uint64(0)
	for plus, times := range CNTS {
		ss1, ss2 := s1, s2
		pp1, pp2 := p1, p2
		if first {
			pp1 += plus
			if pp1 >= 10 {
				pp1 %= 10
			}
			ss1 += pp1 + 1
		} else {
			pp2 += plus
			if pp2 >= 10 {
				pp2 %= 10
			}
			ss2 += pp2 + 1
		}
		rr1, rr2 := play2(pp1, pp2, ss1, ss2, !first)
		r1 += times * rr1
		r2 += times * rr2
	}

	return r1, r2
}

func main() {
	//p1, p2, rolls := play(4, 8)
	p1, p2, rolls := play(6, 10)
	printf("dice rolls: %d, p1: %d, p2: %d", p1, p2, rolls)
	printf("the answer is: %d", min(p1, p2)*rolls)

	//r1, r2 := play2(4-1, 8-1, 0, 0, true)
	r1, r2 := play2(6-1, 10-1, 0, 0, true)
	printf("r1: %d, r2: %d", r1, r2)
}
