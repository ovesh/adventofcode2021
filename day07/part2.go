package day07

import (
	"strconv"
	"strings"
)

func increasingSum(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func uniques(vals []int) map[int]bool {
	res := map[int]bool{}
	for i := 0; i < len(vals); i++ {
		res[vals[i]] = true
	}
	return res
}

func sumForI(n int, vals []int) int {
	sum := 0
	for i := 0; i < len(vals); i++ {
		diff := vals[i] - n
		sum += increasingSum(abs(diff))
	}
	return sum
}

func Part2() int {
	sVals := strings.Split(input, ",")
	vals := make([]int, len(sVals))
	for i := 0; i < len(sVals); i++ {
		vals[i], _ = strconv.Atoi(sVals[i])
	}

	distribution := uniques(vals)
	uniqueVals := make([]int, len(distribution))
	i := 0
	for val := range distribution {
		uniqueVals[i] = val
		i++
	}

	min := 99999999
	for n := 0; n < len(uniqueVals); n++ {
		curSum := sumForI(n, vals)
		if curSum < min {
			min = curSum
		}
	}

	return min
}
