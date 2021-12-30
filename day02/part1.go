package day02

import (
	"strconv"
	"strings"
)

func Part1() int {
	lines := strings.Split(input, "\n")
	depth := 0
	distance := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		cmd := split[0]
		val, _ := strconv.Atoi(split[1])
		switch cmd {
		case "forward":
			distance += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return depth * distance
}
