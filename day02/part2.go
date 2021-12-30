package day02

import (
	"strconv"
	"strings"
)

func Part2() int {
	lines := strings.Split(input, "\n")
	depth := 0
	distance := 0
	aim := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		cmd := split[0]
		val, _ := strconv.Atoi(split[1])
		switch cmd {
		case "forward":
			distance += val
			depth += (aim * val)
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return depth * distance
}
