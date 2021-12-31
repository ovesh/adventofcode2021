package day17

func Part2() int {
	inpt := inputReal
	minXVelocity := minPossibleXV(inpt)
	candidateCount := 0
	for xV := minXVelocity; xV <= inpt.maxX; xV++ {
		for yV := inpt.minY; yV < 1000; yV++ {
			if isCandidate(xV, yV, inpt) {
				candidateCount++
			}
		}
	}
	return candidateCount
}
