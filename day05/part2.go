package day05

func Part2() int {
	inputLines := getInput()
	coords := [inputWidth][inputWidth]coord{}
	for _, line := range inputLines {
		//fmt.Printf("line %#v\n", line)
		if line.startX == line.endX {
			step := 1
			if line.startY >= line.endY {
				step = -1
			}

			for i := line.startY; i != line.endY; i += step {
				coords[line.startX][i].times++
			}
			coords[line.startX][line.endY].times++

		} else if line.startY == line.endY {
			step := 1
			if line.startX >= line.endX {
				step = -1
			}

			for i := line.startX; i != line.endX; i += step {
				coords[i][line.startY].times++
			}
			coords[line.endX][line.startY].times++

		} else {
			stepx := 1
			stepy := 1
			if line.startX >= line.endX {
				stepx = -1
			}
			if line.startY >= line.endY {
				stepy = -1
			}

			y := line.startY
			for x := line.startX; x != line.endX; x += stepx {
				coords[x][y].times++
				y += stepy
			}
			coords[line.endX][line.endY].times++
		}
		//printCoords(coords)

	}

	sum := 0
	for y := 0; y < inputWidth; y++ {
		for x := 0; x < inputWidth; x++ {
			if coords[x][y].times > 1 {
				sum++
			}
		}
	}
	return sum
}
