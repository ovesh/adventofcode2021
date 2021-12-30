package day09

import (
	"strconv"
	"strings"
)

func isLowest(x, y int, matrix [height][width]int) bool {
	cur := matrix[y][x]
	if x > 0 {
		if matrix[y][x-1] <= cur {
			return false
		}
	}
	if x < width-1 {
		if matrix[y][x+1] <= cur {
			return false
		}
	}
	if y > 0 {
		if matrix[y-1][x] <= cur {
			return false
		}
	}
	if y < height-1 {
		if matrix[y+1][x] <= cur {
			return false
		}
	}

	return true
}

func Part1() int {
	matrix := [height][width]int{}
	lines := strings.Split(input, "\n")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			matrix[y][x], _ = strconv.Atoi(string(lines[y][x]))
		}
	}
	sum := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if isLowest(x, y, matrix) {
				sum += matrix[y][x] + 1
				//fmt.Println("found ", x, ", ", y, ": ", matrix[y][x])
			}
		}
	}

	return sum
}
