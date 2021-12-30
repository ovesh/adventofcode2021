package day03

import (
	"fmt"
	"strconv"
	"strings"
)

func pad(s string, to int) string {
	for len(s) < to {
		s = "0" + s
	}
	return s
}

func print(s []uint16, l int) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += pad(strconv.FormatUint(uint64(s[i]), 2), l) + " "
	}
	return res
}

func elimination(iLines []uint16, colNum int, chooseMax bool) uint16 {
	//fmt.Println(print(iLines, colNum))

	for colIdx := 0; colIdx < colNum; colIdx++ {
		//fmt.Println("col ", colIdx)
		oneCounts := 0
		for _, val := range iLines {
			isOne := (((val & (1 << (colNum - colIdx - 1))) >> (colNum - colIdx - 1)) != 0)
			if isOne {
				oneCounts++
			}
		}
		var keepOnes bool
		if chooseMax {
			keepOnes = oneCounts >= (len(iLines) - oneCounts)
		} else {
			keepOnes = oneCounts < (len(iLines) - oneCounts)
		}
		//fmt.Println("oneCounts: ", oneCounts, " keepOnes: ", keepOnes)
		newILines := []uint16{}
		for _, val := range iLines {
			isOne := (((val & (1 << (colNum - colIdx - 1))) >> (colNum - colIdx - 1)) != 0)
			if (isOne && keepOnes) || (!isOne && !keepOnes) {
				newILines = append(newILines, val)
			}
		}
		iLines = newILines
		//fmt.Println(print(iLines, colNum))
		if len(iLines) == 1 {
			return iLines[0]
		}
	}
	panic("not found")
}

func Part2() uint64 {
	lines := strings.Split(input, "\n")
	colNum := len(lines[0])
	iLines := make([]uint16, len(lines))
	for i, line := range lines {
		val, _ := strconv.ParseInt(line, 2, 16)
		iLines[i] = uint16(val)
	}
	oxygen := elimination(iLines, colNum, true)
	fmt.Println(oxygen)
	co2 := elimination(iLines, colNum, false)
	fmt.Println(co2)

	return uint64(oxygen) * uint64(co2)
}
