package day22

import (
	"strconv"
	"strings"
)

var matrix [101][101][101]bool

func normalize(a, b int) (bool, int, int) {
	if b < -50 || a > 50 {
		return false, 0, 0
	}
	if a < -50 {
		a = -50
	}
	if b > 50 {
		b = 50
	}
	return true, a, b
}

func Part1() int {
	totalOn := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		onOrOff := strings.HasPrefix(line, "on ")
		split := strings.Split(line, " ")
		split = strings.Split(split[1], ",")
		xInstructions := strings.Split(split[0][2:], "..")
		xStart, _ := strconv.Atoi(xInstructions[0])
		xEnd, _ := strconv.Atoi(xInstructions[1])
		valid, xStart, xEnd := normalize(xStart, xEnd)
		if !valid {
			continue
		}
		yInstructions := strings.Split(split[1][2:], "..")
		yStart, _ := strconv.Atoi(yInstructions[0])
		yEnd, _ := strconv.Atoi(yInstructions[1])
		valid, yStart, yEnd = normalize(yStart, yEnd)
		if !valid {
			continue
		}
		zInstructions := strings.Split(split[2][2:], "..")
		zStart, _ := strconv.Atoi(zInstructions[0])
		zEnd, _ := strconv.Atoi(zInstructions[1])
		valid, zStart, zEnd = normalize(zStart, zEnd)
		if !valid {
			continue
		}
		for z := zStart; z <= zEnd; z++ {
			for y := yStart; y <= yEnd; y++ {
				for x := xStart; x <= xEnd; x++ {
					if onOrOff && !matrix[z+50][y+50][x+50] {
						totalOn++
					} else if !onOrOff && matrix[z+50][y+50][x+50] {
						totalOn--
					}
					matrix[z+50][y+50][x+50] = onOrOff
				}
			}
		}
	}
	return totalOn
}
