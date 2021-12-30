package day08

import (
	"strings"
)

func Part1() int {
	lines := strings.Split(input, "\n")
	outputs := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		split := strings.Split(lines[i], "|")
		split = strings.Fields(split[1])
		outputs[i] = split
	}

	sum := 0
	for _, o := range outputs {
		for _, s := range o {
			lens := len(s)
			switch lens {
			case 2, 3, 4, 7:
				sum++
			}
		}
	}

	return sum
}
