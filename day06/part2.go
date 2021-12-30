package day06

import (
	"strconv"
	"strings"
)

func Part2() int64 {
	sVals := strings.Split(input, ",")
	vals := make([]int, len(sVals))
	for i := 0; i < len(sVals); i++ {
		vals[i], _ = strconv.Atoi(sVals[i])
	}

	counts := [9]int64{}
	for i := 0; i < len(vals); i++ {
		counts[vals[i]]++
	}

	for day := 0; day < 256; day++ {
		toRestart := counts[0]
		counts[0] = 0
		for i := 1; i < 9; i++ {
			counts[i-1] = counts[i]
		}
		counts[6] += toRestart
		counts[8] = toRestart
	}

	sum := int64(0)
	for i := 0; i < 9; i++ {
		sum += counts[i]
	}

	return sum
}
