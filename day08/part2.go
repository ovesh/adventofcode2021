package day08

import (
	"sort"
	"strings"
)

type byRune []rune

func (r byRune) Len() int           { return len(r) }
func (r byRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byRune) Less(i, j int) bool { return r[i] < r[j] }

func stringToRunes(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// order the string, e.g efgab => abefg
func normalize(s string) string {
	r := stringToRunes(s)
	sort.Sort(byRune(r))
	return string(r)
}

//  aaa
// b   c
// b   c
// b   c
//  ddd
// e   f
// e   f
// e   f
//  ggg

func find(vals []string, length int) string {
	for _, v := range vals {
		if len(v) == length {
			return v
		}
	}
	return ""
}

func contains(s1, s2 string) bool {
	j := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[j] {
			j++
		}
		if j == len(s2) {
			return true
		}
	}
	return false
}

func find3(inputs []string, seven string) string {
	// to find 3, it's the only 5-char that contains 7
	for j := 0; j < len(inputs); j++ {
		if len(inputs[j]) == 5 {
			if contains(inputs[j], seven) {
				return inputs[j]
			}
		}
	}
	panic("no 3")
}

func find9(inputs []string, three string) string {
	// to find 9, it's the only 6-char that contains 3
	for j := 0; j < len(inputs); j++ {
		if len(inputs[j]) == 6 {
			if contains(inputs[j], three) {
				return inputs[j]
			}
		}
	}
	panic("no 9")
}

func find5(inputs []string, three, nine string) string {
	// to find 5, it's the only 5-char that is contained by 9
	for j := 0; j < len(inputs); j++ {
		if len(inputs[j]) == 5 {
			if inputs[j] == three {
				continue
			}
			if contains(nine, inputs[j]) {
				return inputs[j]
			}
		}
	}
	panic("no 5")
}

func find2(inputs []string, three, five, nine string) string {
	// that leaves the 2 as the only 5-char left
	for j := 0; j < len(inputs); j++ {
		if len(inputs[j]) == 5 {
			if inputs[j] == nine {
				continue
			}
			if inputs[j] == three {
				continue
			}
			if inputs[j] == five {
				continue
			}
			return inputs[j]
		}
	}
	panic("no 2")
}

func find0(inputs []string, seven, nine string) string {
	// to find 0, it's the only remaining 6-char that contains 7
	for j := 0; j < len(inputs); j++ {
		if len(inputs[j]) == 6 {
			if inputs[j] == nine {
				continue
			}
			if contains(inputs[j], seven) {
				return inputs[j]
			}
		}
	}
	panic("no 0")
}

func find6(inputs []string, zero, nine string) string {
	// that leaves 6 as the only remaining 6-char
	for j := 0; j < len(inputs); j++ {
		if len(inputs[j]) == 6 {
			if inputs[j] == nine {
				continue
			}
			if inputs[j] == zero {
				continue
			}
			return inputs[j]
		}
	}
	panic("no 6")
}

func Part2() int {
	lines := strings.Split(input, "\n")
	outputs := make([][]string, len(lines))
	inputs := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		split := strings.Split(lines[i], "|")
		inputs[i] = strings.Fields(split[0])
		for j := 0; j < len(inputs[i]); j++ {
			// order the inputs (e.g efgab => abefg) b/c the order of signals doesn't matter
			inputs[i][j] = normalize(inputs[i][j])
		}
		outputs[i] = strings.Fields(split[1])
		for j := 0; j < len(outputs[i]); j++ {
			outputs[i][j] = normalize(outputs[i][j])
		}
	}

	sum := 0
	for lineIdx := 0; lineIdx < len(inputs); lineIdx++ {
		//fmt.Println(inputs[lineIdx])
		// find all the "easy" digits that are the only ones to have exactly N lines turned on
		one := find(inputs[lineIdx], 2)
		if one == "" {
			panic("no 1")
		}
		seven := find(inputs[lineIdx], 3)
		if seven == "" {
			panic("no 7")
		}
		four := find(inputs[lineIdx], 4)
		if four == "" {
			panic("no 4")
		}
		eight := find(inputs[lineIdx], 7)
		if eight == "" {
			panic("no 8")
		}

		// 3 is the only 5-char that contains 7
		three := find3(inputs[lineIdx], seven)
		// 9 is the only 6-char that contains 3
		nine := find9(inputs[lineIdx], three)
		// 5 is the only 5-char that is contained by 9
		five := find5(inputs[lineIdx], three, nine)
		// that leaves the 2 as the only 5-char left
		two := find2(inputs[lineIdx], three, five, nine)
		// 0 is the only remaining 6-char that contains 7
		zero := find0(inputs[lineIdx], seven, nine)
		// that leaves 6 as the only remaining 6-char
		six := find6(inputs[lineIdx], zero, nine)

		m := map[string]int{zero: 0, one: 1, two: 2, three: 3, four: 4, five: 5, six: 6, seven: 7, eight: 8, nine: 9}
		//fmt.Println(m)
		curLineSum := 0
		for _, sNum := range outputs[lineIdx] {
			curLineSum *= 10
			curLineSum += m[sNum]
		}
		//fmt.Println(curLineSum)
		sum += curLineSum
	}

	return sum
}
