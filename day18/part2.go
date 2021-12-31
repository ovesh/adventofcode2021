package day18

import (
	"strings"
)

func Part2() int {
	lines := strings.Split(input, "\n")
	maxMagnitude := 0
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			a := parseRoot(lines[i])
			b := parseRoot(lines[j])
			magnitude := a.add(b).magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
				//fmt.Println("new max:", maxMagnitude)
			}
			a = parseRoot(lines[i])
			b = parseRoot(lines[j])
			magnitude = b.add(a).magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
				//fmt.Println("new max:", maxMagnitude)
			}
		}
	}
	return maxMagnitude
}
