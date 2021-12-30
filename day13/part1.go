package day13

import (
	"strconv"
	"strings"
)

type matrix [][]bool

func newMatrix(height, width int) matrix {
	res := make([][]bool, height)
	for y := 0; y < height; y++ {
		res[y] = make([]bool, width)
	}
	return res
}

func (m matrix) String() string {
	res := ""
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func (m matrix) fold(axis string, where int) matrix {
	oldHeight := len(m)
	oldWidth := len(m[0])
	if axis == "x" {
		res := newMatrix(oldHeight, where)
		firstX := 2*where + 1 - oldWidth
		if firstX < 0 {
			firstX = 0
		}
		newWidth := where
		for y := 0; y < oldHeight; y++ {
			for x := firstX; x < newWidth; x++ {
				res[y][x] = m[y][x] || m[y][oldWidth-x+firstX-1]
			}
		}
		return res
	}
	// axis == "y"
	res := newMatrix(where, oldWidth)
	firstY := 2*where + 1 - oldHeight
	if firstY < 0 {
		firstY = 0
	}
	newHeight := where
	for y := firstY; y < newHeight; y++ {
		for x := 0; x < oldWidth; x++ {
			res[y][x] = m[y][x] || m[oldHeight-y+firstY-1][x]
		}
	}
	return res
}

func (m matrix) countDots() int {
	res := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] {
				res++
			}
		}
	}
	return res
}

func buildMatrix(lines []string) (matrix, int) {
	m := newMatrix(height, width)
	lastCoordLine := 0
	for i, line := range lines {
		if line == "" {
			lastCoordLine = i
			break
		}
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		m[y][x] = true
	}
	return m, lastCoordLine
}

func Part1() int {
	lines := strings.Split(input, "\n")
	m, lastCoordLine := buildMatrix(lines)

	line := lines[lastCoordLine+1][len("fold along "):]
	axis := line[:1]
	where, _ := strconv.Atoi(line[2:])
	m = m.fold(axis, where)
	return m.countDots()
}
