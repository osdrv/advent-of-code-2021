package main

func fire(vx, vy, xmin, xmax, ymin, ymax int) (int, bool) {
	x, y := 0, 0
	maxY := -ALOT
	for {
		x += vx
		if vx > 0 {
			vx -= 1
		} else if vx < 0 {
			vx += 1
		}
		y += vy
		vy--

		//printf("vx: %d, vy: %d, x: %d y: %d", vx, vy, x, y)

		if y > maxY {
			maxY = y
		}

		if x >= xmin && x <= xmax && y >= ymin && y <= ymax {
			// we hit it
			return maxY, true
		}

		if y < ymin {
			// we missed it
			return -1, false
		}
	}
}

func main() {

	xmin, xmax, ymin, ymax := 143, 177, -106, -71
	//xmin, xmax, ymin, ymax := 20, 30, -10, -5

	maxY := -ALOT

	hits := 0
	for vx0 := 1; vx0 < 200; vx0++ {
		for vy0 := -200; vy0 < 200; vy0++ {
			locY, ok := fire(vx0, vy0, xmin, xmax, ymin, ymax)
			if ok {
				hits++
				if locY > maxY {
					maxY = locY
				}
			}
		}
	}

	printf("max y: %d, hits: %d", maxY, hits)
}
