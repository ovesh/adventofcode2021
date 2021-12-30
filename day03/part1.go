package day03

import (
	"strconv"
	"strings"
)

func Part1() (uint16, uint16) {
	lines := strings.Split(input, "\n")
	colNum := len(lines[0])
	oneCountsPerColumn := make([]int, colNum)
	for _, line := range lines {
		val, _ := strconv.ParseInt(line, 2, 16)
		for i := colNum - 1; i >= 0; i-- {
			if ((val & (1 << i)) >> i) != 0 {
				oneCountsPerColumn[colNum-i-1] = oneCountsPerColumn[colNum-i-1] + 1
			}
		}
	}

	threshold := len(lines) / 2
	var mask uint16
	for i := 0; i < colNum-1; i++ {
		mask += 1
		mask <<= 1
	}
	mask += 1
	var gamma uint16
	for i := range oneCountsPerColumn {
		if oneCountsPerColumn[i] > threshold {
			gamma |= (1 << (colNum - 1 - i))
		}
	}

	var epsilon uint16 = (^gamma) & mask

	return gamma, epsilon
}
