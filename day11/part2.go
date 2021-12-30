package day11

func (m matrix) allFlashing() bool {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if m[y][x].val != 0 {
				return false
			}
		}
	}
	return true
}

func Part2() int {
	m := prepareMatrix()

	for t := 0; t < 1000; t++ {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				m[y][x].plus(turn{i: t})
			}
		}
		if m.allFlashing() {
			return t + 1
		}
	}
	return 0
}
