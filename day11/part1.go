package day11

import (
	"strconv"
	"strings"
)

type node struct {
	val       int
	neighbors []*node
	didTurn   int
}

type turn struct {
	i int
}

func (n *node) String() string {
	return strconv.Itoa(n.val)
}

var numFlashes = 0

func (n *node) plus(t turn) {
	if n.didTurn == t.i {
		return
	}

	n.val++
	if n.val > 9 {
		n.didTurn = t.i
		n.val = 0
		numFlashes++
		n.broadcast(t)
	}
}

func (n *node) broadcast(t turn) {
	for _, neighbor := range n.neighbors {
		neighbor.plus(t)
	}
}

type matrix [10][10]*node

func (m matrix) String() string {
	res := ""
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			res += m[y][x].String()
		}
		res += "\n"
	}
	return res
}

func prepareMatrix() matrix {
	m := matrix{}
	lines := strings.Split(input, "\n")
	for y := 0; y < 10; y++ {
		line := lines[y]
		for x, c := range line {
			val, _ := strconv.Atoi(string(c))
			m[y][x] = &node{val: val, didTurn: -1}
		}
	}
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if x > 0 {
				m[y][x].neighbors = append(m[y][x].neighbors, m[y][x-1])
			}
			if x < 9 {
				m[y][x].neighbors = append(m[y][x].neighbors, m[y][x+1])
			}
			if y > 0 {
				m[y][x].neighbors = append(m[y][x].neighbors, m[y-1][x])
			}
			if y < 9 {
				m[y][x].neighbors = append(m[y][x].neighbors, m[y+1][x])
			}
			if x > 0 && y > 0 {
				m[y][x].neighbors = append(m[y][x].neighbors, m[y-1][x-1])
			}
			if x > 0 && y < 9 {
				m[y][x].neighbors = append(m[y][x].neighbors, m[y+1][x-1])
			}
			if x < 9 && y > 0 {
				m[y][x].neighbors = append(m[y][x].neighbors, m[y-1][x+1])
			}
			if x < 9 && y < 9 {
				m[y][x].neighbors = append(m[y][x].neighbors, m[y+1][x+1])
			}
		}
	}
	return m
}

func Part1() int {
	m := prepareMatrix()
	for t := 0; t < 100; t++ {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				m[y][x].plus(turn{i: t})
			}
		}
	}
	return numFlashes
}
