package day07

import (
	"sort"
	"strconv"
	"strings"
)

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func Part1() int {
	sVals := strings.Split(input, ",")
	vals := make([]int, len(sVals))
	for i := 0; i < len(sVals); i++ {
		vals[i], _ = strconv.Atoi(sVals[i])
	}

	sort.Ints(vals)
	median := vals[len(vals)/2]
	sum := 0
	for i := 0; i < len(vals); i++ {
		sum += abs(vals[i] - median)
	}

	return sum
}
