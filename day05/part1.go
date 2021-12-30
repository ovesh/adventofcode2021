package day05

import (
	"fmt"
	"strconv"
	"strings"
)

//const inputWidth = 10
const inputWidth = 991

type coord struct {
	times int
}

type inputLine struct {
	startX int
	startY int
	endX   int
	endY   int
}

func printCoords(coords [inputWidth][inputWidth]coord) {
	for y := 0; y < inputWidth; y++ {
		for x := 0; x < inputWidth; x++ {
			fmt.Printf("%d ", coords[x][y].times)
		}
		fmt.Println()
	}
	fmt.Println()
}

func getInput() []inputLine {
	lines := strings.Split(input, "\n")
	inputLines := []inputLine{}
	for _, line := range lines {
		split1 := strings.Split(line, " -> ")
		split2 := strings.Split(split1[0], ",")
		startX, _ := strconv.Atoi(split2[0])
		startY, _ := strconv.Atoi(split2[1])
		split3 := strings.Split(split1[1], ",")
		endX, _ := strconv.Atoi(split3[0])
		endY, _ := strconv.Atoi(split3[1])
		inputLines = append(inputLines, inputLine{startX: startX, startY: startY, endX: endX, endY: endY})
	}
	return inputLines
}

func Part1() int {
	inputLines := getInput()
	coords := [inputWidth][inputWidth]coord{}
	for _, line := range inputLines {
		//fmt.Printf("line %#v\n", line)
		if line.startX == line.endX {
			step := 1
			if line.startY >= line.endY {
				step = -1
			}

			for i := line.startY; i != line.endY; i += step {
				coords[line.startX][i].times++
			}
			coords[line.startX][line.endY].times++

		} else if line.startY == line.endY {
			step := 1
			if line.startX >= line.endX {
				step = -1
			}

			for i := line.startX; i != line.endX; i += step {
				coords[i][line.startY].times++
			}
			coords[line.endX][line.startY].times++

		}
		//printCoords(coords)

	}

	sum := 0
	for y := 0; y < inputWidth; y++ {
		for x := 0; x < inputWidth; x++ {
			if coords[x][y].times > 1 {
				sum++
			}
		}
	}
	return sum
}
