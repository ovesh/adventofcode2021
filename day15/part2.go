package day15

import (
	"sort"
	"strconv"
	"strings"
)

func buildMatrix2(lines []string) [][]*square {
	origHeight := len(lines)
	origWidth := len(lines[0])
	res := make([][]*square, 5*origHeight)
	for bigY := 0; bigY < 5; bigY++ {
		for y := 0; y < origHeight; y++ {
			line := lines[y]
			res[y+bigY*origHeight] = make([]*square, 5*origWidth)
			for bigX := 0; bigX < 5; bigX++ {
				for x, c := range line {
					num, _ := strconv.Atoi(string(c))
					if bigX > 0 {
						num = res[y+bigY*origHeight][x+(bigX-1)*origWidth].i + 1
					}
					if bigY > 0 {
						num = res[y+(bigY-1)*origHeight][x+bigX*origWidth].i + 1
					}
					num = num % 10
					if num == 0 {
						num++
					}
					res[y+bigY*origHeight][x+bigX*origWidth] = &square{
						i:             num,
						traversed:     false,
						x:             x,
						y:             y,
						localMinScore: MaxInt,
					}
				}
			}
		}
	}
	for y := 0; y < len(res); y++ {
		for x := 0; x < len(res[0]); x++ {
			cur := res[y][x]
			if x > 0 {
				cur.neighbors = append(cur.neighbors, res[y][x-1])
			}
			if x < len(res[0])-1 {
				cur.neighbors = append(cur.neighbors, res[y][x+1])
			}
			if y > 0 {
				cur.neighbors = append(cur.neighbors, res[y-1][x])
			}
			if y < len(res)-1 {
				cur.neighbors = append(cur.neighbors, res[y+1][x])
			}
			sort.Sort(bySquareVal(cur.neighbors))
		}
	}
	return res
}

func Part2() int {
	lines := strings.Split(input, "\n")
	matrix := buildMatrix2(lines)
	//for y := 0; y < len(matrix); y++ {
	//	for x := 0; x < len(matrix); x++ {
	//		fmt.Printf("%d", matrix[y][x].i)
	//	}
	//	fmt.Println()
	//}
	start := matrix[0][0]
	start.i = 0
	end := matrix[len(matrix)-1][len(matrix[0])-1]
	res := start.traverse(0, end)
	return res
}
