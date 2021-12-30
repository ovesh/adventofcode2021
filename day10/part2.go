package day10

import (
	"sort"
	"strings"
)

var points2 = map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}

func processLine2(line string) int {
	s := stack{}
	score := 0
	for _, r := range line {
		c := string(r)
		switch c {
		case "(", "[", "{", "<":
			s.push(c)
		case ")", "]", "}", ">":
			popped := s.pop()
			if popped != close2open[c] {
				return 0
			}
		}
	}
	if len(s) > 0 {
		for len(s) > 0 {
			popped := s.pop()
			score *= 5
			score += points2[popped]
		}
	}
	return score
}

func Part2() int {
	lines := strings.Split(input, "\n")
	scores := []int{}
	for _, line := range lines {
		score := processLine2(line)
		if score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
