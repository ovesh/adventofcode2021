package day06

import (
	"strconv"
	"strings"
)

const numDays = 80

func Part1() int {
	sVals := strings.Split(input, ",")
	vals := make([]int, len(sVals))
	for i := 0; i < len(sVals); i++ {
		vals[i], _ = strconv.Atoi(sVals[i])
	}

	for day := 0; day < numDays; day++ {
		lenVals := len(vals)
		for i := 0; i < lenVals; i++ {
			if vals[i] == 0 {
				vals = append(vals, 8)
				vals[i] = 6
			} else {
				vals[i]--
			}
		}
	}
	return len(vals)
}
