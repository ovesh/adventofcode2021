package day01

import (
	"strconv"
	"strings"
)

func Part1() int {
	lines := strings.Split(input, "\n")
	firstDepth, _ := strconv.Atoi(lines[0])
	lastDepth := firstDepth
	res := 0
	for i := 1; i < len(lines); i++ {
		depth, _ := strconv.Atoi(lines[i])
		if depth > lastDepth {
			res++
		}
		lastDepth = depth
	}

	return res
}
