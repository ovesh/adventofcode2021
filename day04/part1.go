package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type bingoBoard struct {
	rows        [5][5]int
	marked      [5][5]bool
	markedCount int
	// for part 2
	origIdx int
}

func parseBoard(lines []string) *bingoBoard {
	rows := [5][5]int{}
	for rowIdx := 0; rowIdx < 5; rowIdx++ {
		line := lines[rowIdx]
		sVals := strings.Fields(line)
		for colIdx := 0; colIdx < 5; colIdx++ {
			val, _ := strconv.Atoi(sVals[colIdx])
			rows[rowIdx][colIdx] = val
		}
	}
	return &bingoBoard{marked: [5][5]bool{}, rows: rows}
}

func (b *bingoBoard) mark(i int) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if b.rows[row][col] == i {
				b.marked[row][col] = true
				b.markedCount++
			}
		}
	}
}

func (b *bingoBoard) rowDone() bool {
	for row := 0; row < 5; row++ {
		done := true
		for col := 0; col < 5; col++ {
			if !b.marked[row][col] {
				done = false
				break
			}
		}
		if done {
			return true
		}
	}
	return false
}

func (b *bingoBoard) colDone() bool {
	for col := 0; col < 5; col++ {
		done := true
		for row := 0; row < 5; row++ {
			if !b.marked[row][col] {
				done = false
				break
			}
		}
		if done {
			return true
		}
	}
	return false
}

func (b *bingoBoard) unmarkedSum() int {
	sum := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !b.marked[row][col] {
				sum += b.rows[row][col]
			}
		}
	}
	return sum
}

func prepareBoards() ([]int, []*bingoBoard) {
	lines := strings.Split(input, "\n")
	sRandoms := strings.Split(lines[0], ",")
	randoms := make([]int, len(sRandoms))
	for i, sNum := range sRandoms {
		val, _ := strconv.Atoi(sNum)
		randoms[i] = val
	}

	lines = lines[1:]
	boards := []*bingoBoard{}
	for len(lines) > 0 {
		board := parseBoard(lines[1:6])
		if board == nil {
			break
		}
		board.origIdx = len(boards)
		boards = append(boards, board)
		lines = lines[6:]
	}

	return randoms, boards
}

func Part1() int {
	randoms, boards := prepareBoards()

	for _, nextRand := range randoms {
		for i, board := range boards {
			board.mark(nextRand)
			if board.rowDone() || board.colDone() {
				fmt.Printf("winner! board %d: %#v, %#v\n", i, board.rows, board.marked)
				return board.unmarkedSum() * nextRand
			}
		}
	}
	panic("winner not found")
}
