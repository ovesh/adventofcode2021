package day13

import (
	"strconv"
	"strings"
)

func Part2() string {
	lines := strings.Split(input, "\n")
	m, lastCoordLine := buildMatrix(lines)

	for i := lastCoordLine + 1; i < len(lines); i++ {
		line := lines[i][len("fold along "):]
		//fmt.Println(line)
		axis := line[:1]
		where, _ := strconv.Atoi(line[2:])

		m = m.fold(axis, where)
		//fmt.Println(m)
		//fmt.Println(m.countDots())
	}
	return m.String()
}
