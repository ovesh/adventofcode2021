package day04

func Part2() int {
	randoms, boards := prepareBoards()
	lastWonIdx := -1
	lastWonRandom := -1
	boardsCopy := make([]*bingoBoard, len(boards))
	for i := 0; i < len(boards); i++ {
		boardsCopy[i] = boards[i]
	}
	for _, nextRand := range randoms {
		for i, board := range boardsCopy {
			board.mark(nextRand)
			if board.rowDone() || board.colDone() {
				//fmt.Printf("winner! lastRand %d board %d: %#v, %#v\n", nextRand, board.origIdx, board.rows, board.marked)
				lastWonIdx = board.origIdx
				lastWonRandom = nextRand
				boardsCopy = append(boardsCopy[:i], boardsCopy[i+1:]...)
				for j := i; j < len(boardsCopy); j++ {
					boardsCopy[j].mark(nextRand)
					if boardsCopy[j].rowDone() || boardsCopy[j].colDone() {
						boardsCopy = append(boardsCopy[:j], boardsCopy[j+1:]...)
						j--
					}
				}
				break
			}
		}
	}
	return boards[lastWonIdx].unmarkedSum() * lastWonRandom
}
