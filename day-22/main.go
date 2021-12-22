package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Cuboid struct {
	v                      int
	x0, x1, y0, y1, z0, z1 int
}

func NewCuboid(v, x0, x1, y0, y1, z0, z1 int) *Cuboid {
	return &Cuboid{v, x0, x1, y0, y1, z0, z1}
}

func (c *Cuboid) String() string {
	if c.v == 0 {
		return "C{empty}"
	}
	return fmt.Sprintf("C{%d, %d..%d, %d..%d, %d..%d}", c.v, c.x0, c.x1, c.y0, c.y1, c.z0, c.z1)
}

func (c *Cuboid) Size() uint64 {
	if c.v == 0 {
		return 0
	}
	return uint64(abs(c.x1-c.x0)+1) * uint64(abs(c.y1-c.y0)+1) * uint64(abs(c.z1-c.z0)+1)
}

func parseCuboid(s string) *Cuboid {
	v := 0
	if strings.HasPrefix(s, "on") {
		v = 1
		s = s[3:]
	} else {
		v = -1
		s = s[4:]
	}
	chs := strings.Split(s, ",")
	x0, x1 := parseRange(chs[0][2:])
	y0, y1 := parseRange(chs[1][2:])
	z0, z1 := parseRange(chs[2][2:])

	return NewCuboid(v, x0, x1, y0, y1, z0, z1)
}

func parseRange(s string) (int, int) {
	chs := strings.Split(s, "..")
	return parseInt(chs[0]), parseInt(chs[1])
}

func solve1(cuboids []*Cuboid) int {
	reactor := make(map[Point3]bool)

	for _, cuboid := range cuboids {
		if cuboid.x0 < -50 || cuboid.x1 > 50 {
			continue
		}
		if cuboid.y0 < -50 || cuboid.y1 > 50 {
			continue
		}
		if cuboid.z0 < -50 || cuboid.z1 > 50 {
			continue
		}
		for x := max(cuboid.x0, -50); x <= min(cuboid.x1, 50); x++ {
			for y := max(cuboid.y0, -50); y <= min(cuboid.y1, 50); y++ {
				for z := max(cuboid.z0, -50); z <= min(cuboid.z1, 50); z++ {
					p := NewPoint3(x, y, z)
					reactor[*p] = cuboid.v == 1
				}
			}
		}
	}

	cnt := 0
	for _, p := range reactor {
		if p {
			cnt++
		}
	}

	return cnt
}

/*

    a0.....a1

	  b0.......b1

	  c0...c1

    a0.........a1
       b0...b1
       c0...c1

    a0.........a1
	               b0....b1

*/

func intersect(a0, a1, b0, b1 int) (int, int, bool) {
	c0 := max(a0, b0)
	c1 := min(a1, b1)
	if c0 > c1 {
		return 0, 0, false
	}
	return c0, c1, true
}

func (c1 *Cuboid) Intersect(c2 *Cuboid) *Cuboid {
	x0, x1, xok := intersect(c1.x0, c1.x1, c2.x0, c2.x1)
	y0, y1, yok := intersect(c1.y0, c1.y1, c2.y0, c2.y1)
	z0, z1, zok := intersect(c1.z0, c1.z1, c2.z0, c2.z1)

	if xok && yok && zok {
		return NewCuboid(1, x0, x1, y0, y1, z0, z1)
	}

	return &Cuboid{v: 0}
}

/*
	on: c1 U c2 = c1 + c1 - c1/\c2

	on: c1, c2
	off: c1/\c2

	turning off:
	for cc in enabled {
		intersect := cc.intersect(off)
		total -= intersect.size

	}

    disabled:
	enabled:

	T = (CC - OFF).size

	disable goes over every cuboid in enabled
	enabled if enabled cuboids overlap, total sum subtraction will subtract it twice

	enabled: cc1, cc2
	disabled: []

	off: ccx

	cc1 intersects cc2 in in12
	cc1 intersects ccx in ccx1
	cc2 intersects ccx in ccx2

	enabled: [], disabled: []
	enabled: [cc1], disabled: []
	enabled: [cc1, cc2], disabled: [in12]
	enabled: [cc1, cc2, ccx1 /\ ccx2], disabled: [ccx1, ccx2]

	plus = [cc1, cc2]
	minus = [in12]

	plus = [cc1, cc2, cc3]
	minus = [in12, in13, in23]

	invariant: plus.size - minus.zise = t

	1. t = 0
	2. t += cc1
	3. t += cc2 - in12
	4. t -= cc1 /\ ccx
	5. t -= cc2 /\ ccx
	6. t += ccx1 /\ccx2


	plus = [on1, on2, on3]
	minus = [in12, in13, in23]

	disable x1
	plus = [on1, on2, on3, (inx1 /\ inx2), (inx1 /\ inx3), (inx2 /\ inx3)]
	minus = [in12, in13, in23, inx1, inx2, inx3]
*/

func solve2(cuboids []*Cuboid) uint64 {
	out := make([]*Cuboid, 0, 1)
	for _, cuboid := range cuboids {
		newout := make([]*Cuboid, len(out))
		copy(newout, out)
		for _, cc := range out {
			if in := cuboid.Intersect(cc); in.Size() > 0 {
				if cc.v == 1 {
					in.v = -1
				}
				newout = append(newout, in)
			}
		}
		if cuboid.v == 1 {
			newout = append(newout, cuboid)
		}
		out = newout
	}

	res := int64(0)
	for _, cuboid := range out {
		res += int64(cuboid.v) * int64(cuboid.Size())
	}
	return uint64(res)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)
	cuboids := make([]*Cuboid, 0, len(lines))

	for _, line := range lines {
		cuboid := parseCuboid(line)
		cuboids = append(cuboids, cuboid)
	}

	log.Printf("cuboids: %+v", cuboids)

	printf("part 1: %d", solve1(cuboids))

	printf("part 2: %d", solve2(cuboids))
}
