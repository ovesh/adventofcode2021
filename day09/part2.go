package day09

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	children []*node
	val      int
	x        int
	y        int
	marked   bool
	size     int
}

func (n *node) String() string {
	return fmt.Sprintf("x: %d, y: %d, val: %d, size: %d, children: %s", n.x, n.y, n.val, n.size, n.children)
}

func fillTree(x, y int, matrix [height][width]*node) *node {
	src := matrix[y][x]
	if src.val == 9 {
		return nil
	}
	if src.marked {
		return nil
	}

	src.marked = true

	var left, right, bottom, top *node
	if x > 0 {
		left = fillTree(x-1, y, matrix)
		if left != nil {
			src.children = append(src.children, left)
			src.size += left.size
		}
	}
	if x < width-1 {
		right = fillTree(x+1, y, matrix)
		if right != nil {
			src.children = append(src.children, right)
			src.size += right.size
		}
	}
	if y > 0 {
		bottom = fillTree(x, y-1, matrix)
		if bottom != nil {
			src.children = append(src.children, bottom)
			src.size += bottom.size
		}
	}
	if y < height-1 {
		top = fillTree(x, y+1, matrix)
		if top != nil {
			src.children = append(src.children, top)
			src.size += top.size
		}
	}

	return src
}

func Part2() int {
	matrix := [height][width]*node{}
	lines := strings.Split(input, "\n")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			val, _ := strconv.Atoi(string(lines[y][x]))
			matrix[y][x] = &node{children: []*node{}, val: val, x: x, y: y, size: 1}
		}
	}

	//fmt.Println(fillTree(0, 0, matrix))
	sizes := []int{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			node := fillTree(x, y, matrix)
			if node != nil {
				sizes = append(sizes, node.size)
				//fmt.Println(node)
			}
		}
	}

	sort.Ints(sizes)

	product := 1
	for i := 3; i > 0; i-- {
		product *= sizes[len(sizes)-i]
	}

	return product
}
