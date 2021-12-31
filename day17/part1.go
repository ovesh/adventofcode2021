package day17

var memo = map[int]int{}

func sumUpTo(n int) int {
	if res, ok := memo[n]; ok {
		return res
	}
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func isCandidate(xV, yV int, inpt input) bool {
	curX := 0
	curY := 0
	for curX < inpt.maxX && curY > inpt.minY {
		curX += xV
		curY += yV
		if curX >= inpt.minX && curX <= inpt.maxX && curY >= inpt.minY && curY <= inpt.maxY {
			return true
		}
		if xV > 0 {
			xV--
		}
		yV--
	}
	return false
}

func minPossibleXV(inpt input) int {
	for i := 1; true; i++ {
		cur := sumUpTo(i)
		if cur > inpt.minX {
			return i
		}
	}
	return 0
}

func Part1() int {
	inpt := inputReal
	minXVelocity := minPossibleXV(inpt)
	maxYV := -1
	for xV := minXVelocity; xV < inpt.maxX; xV++ {
		for yV := 0; yV < 1000; yV++ {
			if isCandidate(xV, yV, inpt) {
				if yV > maxYV {
					maxYV = yV
				}
			}
		}
	}
	return sumUpTo(maxYV)
}
