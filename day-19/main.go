package main

import (
	"os"
)

const (
	MATCH = 12
)

type Scan struct {
	offset *Point3
	points []*Point3

	lastoff int // a hack to speed up scan re-processing
}

func NewScan(lines []string) *Scan {
	points := make([]*Point3, 0, 1)
	for _, line := range lines {
		nn := parseInts(line)
		points = append(points, NewPoint3(nn[0], nn[1], nn[2]))
	}

	//_, points = norm(points)

	printf("points: %+v", points)
	return &Scan{
		offset: NewPoint3(0, 0, 0), //offset,
		points: points,
	}
}

func norm(points []*Point3) (*Point3, []*Point3) {
	minx, miny, minz := points[0].x, points[0].y, points[0].z
	for _, point := range points {
		if point.x < minx {
			minx = point.x
		}
		if point.y < miny {
			miny = point.y
		}
		if point.z < minz {
			minz = point.z
		}
	}
	upd := make([]*Point3, 0, len(points))
	for _, point := range points {
		upd = append(upd, NewPoint3(point.x-minx, point.y-miny, point.z-minz))
	}
	return NewPoint3(minx, miny, minz), upd
}

func (s *Scan) RotateZ() {
	for _, point := range s.points {
		point.x, point.y = -point.y, point.x
	}
	s.offset.x, s.offset.y = -s.offset.y, s.offset.x
}

func (s *Scan) RotateX() {
	for _, point := range s.points {
		point.y, point.z = -point.z, point.y
	}
	s.offset.y, s.offset.z = -s.offset.z, s.offset.y
}

func (s *Scan) RotateY() {
	for _, point := range s.points {
		point.z, point.x = -point.x, point.z
	}
	s.offset.z, s.offset.x = -s.offset.x, s.offset.z
	//for _, point := range s.points {
	//	point.x, point.z = -point.z, point.x
	//}
	//s.offset.x, s.offset.z = -s.offset.z, s.offset.x
}

func pointsMatch(p1, p2 *Point3) bool {
	return p1.x == p2.x && p1.y == p2.y && p1.z == p2.z
}

func intersect(s1, s2 *Scan) (*Point3, []*Point3) {
	var maxMatched []*Point3
	var matchOff *Point3
Off:
	for _, b1 := range s1.points {
		for _, b2 := range s2.points {
			matched := make([]*Point3, 0, 1)
			for _, p1 := range s1.points {
				for _, p2 := range s2.points {
					p01 := NewPoint3(p1.x-b1.x, p1.y-b1.y, p1.z-b1.z)
					p02 := NewPoint3(p2.x-b2.x, p2.y-b2.y, p2.z-b2.z)
					if pointsMatch(p01, p02) {
						matched = append(matched, p1)
					}
				}
			}
			if len(matched) > len(maxMatched) {
				maxMatched = matched
				matchOff = NewPoint3(
					s1.offset.x+b1.x-b2.x,
					s1.offset.y+b1.y-b2.y,
					s1.offset.z+b1.z-b2.z,
				)
				if len(maxMatched) >= MATCH {
					break Off
				}
			}
		}
	}

	// s2 is shifted from s1 by b1-b2

	return matchOff, maxMatched
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	scans := make([]*Scan, 0, 1)
	ptr, off := 1, 1
	for ptr < len(lines) {
		if lines[ptr] == "" {
			scans = append(scans, NewScan(lines[off:ptr]))
			ptr++ // empty line
			ptr++ // header
			off = ptr
		} else {
			ptr++
		}
	}
	scans = append(scans, NewScan(lines[off:ptr]))

	matched := make([]*Scan, 0, 1)
	matched = append(matched, scans[0])

	for len(scans) > 0 {
		newscans := make([]*Scan, 0, 1)
	Scan:
		for _, scan := range scans {
			for i := 0; i < 4; i++ {
				for j := 0; j < 4; j++ {
					for k := 0; k < 4; k++ {
						for _, mm := range matched[scan.lastoff:] {
							off, mtch := intersect(mm, scan)
							if len(mtch) >= MATCH {
								scan.offset = off
								matched = append(matched, scan)
								continue Scan
							}
						}
						scan.RotateX()
					}
					scan.RotateY()
				}
				scan.RotateZ()
			}
			printf("re-processing scan: %+v", scan)
			newscans = append(newscans, scan)
			scan.lastoff = len(matched)
		}
		scans = newscans
		printf("newscan length: %d", len(newscans))
	}

	points := make(map[Point3]struct{})

	for _, scan := range matched {
		for _, pp := range scan.points {
			np := NewPoint3(
				scan.offset.x+pp.x,
				scan.offset.y+pp.y,
				scan.offset.z+pp.z,
			)
			points[*np] = struct{}{}
		}
	}

	printf("points(%d)", len(points))

	maxdist := 0
	for i := 0; i < len(matched); i++ {
		for j := i + 1; j < len(matched); j++ {
			dist := manhattan(matched[i].offset, matched[j].offset)
			if dist > maxdist {
				maxdist = dist
				printf("%d, %d, %d",
					abs(matched[i].offset.x-matched[j].offset.x),
					abs(matched[i].offset.y-matched[j].offset.y),
					abs(matched[i].offset.z-matched[j].offset.z),
				)
			}
		}
	}

	printf("max distance: %d", maxdist)
}

func manhattan(p1, p2 *Point3) int {
	printf("p1: %s", p1)
	printf("p2: %s", p2)
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}
