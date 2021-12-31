package day25

import (
	"fmt"
	"strings"
)

func cp(m [][]string) [][]string {
	res := make([][]string, len(m))
	for y := 0; y < len(m); y++ {
		res[y] = make([]string, len(m[0]))
		for x := 0; x < len(m[0]); x++ {
			res[y][x] = m[y][x]
		}
	}
	return res
}

func prnt(matrix [][]string) {
	for y := 0; y < len(matrix); y++ {
		fmt.Println(strings.Join(matrix[y], ""))
	}
}

func step(matrix [][]string) ([][]string, int) {
	width := len(matrix[0])
	height := len(matrix)
	cpd := cp(matrix)
	moves := 0
	// east
	for y := 0; y < height; y++ {
		for x := width - 1; x >= 0; x-- {
			if matrix[y][x] != ">" {
				continue
			}
			oneEast := matrix[y][(x+1)%width]
			if oneEast == "." {
				cpd[y][x] = "."
				cpd[y][(x+1)%width] = ">"
				moves++
			}
		}
	}
	matrix = cp(cpd)
	// south
	for y := 0; y < height; y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] != "v" {
				continue
			}
			oneSouth := matrix[(y+1)%height][x]
			if oneSouth == "." {
				cpd[y][x] = "."
				cpd[(y+1)%height][x] = "v"
				moves++
			}
		}
	}
	return cpd, moves
}

func Part1() int {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)
	matrix := make([][]string, height)
	for y, line := range lines {
		matrix[y] = make([]string, width)
		for x := 0; x < width; x++ {
			matrix[y][x] = line[x : x+1]
		}
	}
	moves := 0
	for i := 0; true; i++ {
		matrix, moves = step(matrix)
		//fmt.Println("i=", i, "moves=", moves)
		//prnt(matrix)
		if moves == 0 {
			return i + 1
		}
	}
	panic("not found")
}
