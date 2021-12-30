package day15

import (
	"sort"
	"strconv"
	"strings"
)

type square struct {
	i             int
	traversed     bool
	neighbors     []*square
	localMinScore int
	// only for debugging
	x int
	y int
}

type bySquareVal []*square

func (r bySquareVal) Len() int           { return len(r) }
func (r bySquareVal) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r bySquareVal) Less(i, j int) bool { return r[i].i < r[j].i }

func buildMatrix(lines []string) [][]*square {
	res := make([][]*square, len(lines))
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		res[y] = make([]*square, len(lines[0]))
		for x, c := range line {
			num, _ := strconv.Atoi(string(c))
			res[y][x] = &square{
				i:             num,
				traversed:     false,
				x:             x,
				y:             y,
				localMinScore: MaxInt,
			}
		}
	}
	for y := 0; y < len(res); y++ {
		for x := 0; x < len(res[0]); x++ {
			cur := res[y][x]
			if x > 0 {
				cur.neighbors = append(cur.neighbors, res[y][x-1])
			}
			if x < len(res[0])-1 {
				cur.neighbors = append(cur.neighbors, res[y][x+1])
			}
			if y > 0 {
				cur.neighbors = append(cur.neighbors, res[y-1][x])
			}
			if y < len(res)-1 {
				cur.neighbors = append(cur.neighbors, res[y+1][x])
			}
			sort.Sort(bySquareVal(cur.neighbors))
		}
	}
	return res
}

const MaxInt = int(^uint(0) >> 1)

var globalMin = MaxInt

var times = 0

func (s *square) traverse(curScore int, end *square) int {
	if curScore > globalMin {
		return -1
	}
	if s.traversed {
		return -1
	}
	if curScore+s.i > s.localMinScore {
		return -1
	}
	s.localMinScore = curScore + s.i
	s.traversed = true
	if s == end {
		times++
		s.traversed = false
		if curScore+s.i < globalMin {
			globalMin = curScore + s.i
		}
		return s.i
	}
	min := -1
	for _, neighbor := range s.neighbors {
		cur := neighbor.traverse(curScore+s.i, end)
		if cur == -1 {
			continue
		}
		if min == -1 || cur < min {
			min = cur
		}
	}
	s.traversed = false
	if min == -1 { // all neighbors traversed
		return -1
	}
	return s.i + min
}

func Part1() int {
	lines := strings.Split(input, "\n")
	matrix := buildMatrix(lines)
	start := matrix[0][0]
	start.i = 0
	end := matrix[len(matrix)-1][len(matrix[0])-1]
	res := start.traverse(0, end)
	return res
}
