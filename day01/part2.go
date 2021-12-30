package day01

import (
	"strconv"
	"strings"
)

func windowDepth(lines []string, j int) int {
	res := 0
	for i := j; i < j+3; i++ {
		depth, _ := strconv.Atoi(lines[i])
		res += depth
	}
	return res
}

func Part2() int {
	lines := strings.Split(input, "\n")
	firstWindowSum := windowDepth(lines, 0)
	lastDepth := firstWindowSum
	res := 0
	for i := 1; i < len(lines)-2; i++ {
		depth := windowDepth(lines, i)
		if depth > lastDepth {
			res++
		}
		lastDepth = depth
	}
	return res
}
