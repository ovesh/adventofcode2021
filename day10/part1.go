package day10

import (
	"strings"
)

type stack []string

func (s *stack) push(c string) {
	*s = append(*s, c)
}

func (s *stack) pop() string {
	res := []string(*s)[len(*s)-1]
	*s = []string(*s)[0 : len(*s)-1]
	return res
}

var close2open = map[string]string{")": "(", "]": "[", ">": "<", "}": "{"}
var points = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}

func processLine(line string) int {
	s := stack{}
	for _, r := range line {
		c := string(r)
		switch c {
		case "(", "[", "{", "<":
			s.push(c)
		case ")", "]", "}", ">":
			popped := s.pop()
			if popped != close2open[c] {
				return points[c]
			}
		}
	}
	return 0
}

func Part1() int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		sum += processLine(line)
	}
	return sum
}
