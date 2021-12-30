package day14

import (
	"sort"
	"strings"
)

type insertion struct {
	where int
	what  string
}

type byInsertion []insertion

func (r byInsertion) Len() int           { return len(r) }
func (r byInsertion) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byInsertion) Less(i, j int) bool { return r[i].where > r[j].where }

func buildLocations(template string) map[string][]int {
	res := map[string][]int{}
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		curLocs, ok := res[pair]
		if !ok {
			res[pair] = []int{}
		}
		res[pair] = append(curLocs, i)
	}
	return res
}

func occurrences(template string) map[string]int {
	res := map[string]int{}
	for i := range template {
		c := string(template[i])
		count, ok := res[c]
		if !ok {
			count = 0
		}
		count++
		res[c] = count
	}
	return res
}

func maxMinOccurrences(m map[string]int) (int, int) {
	min := 99999999999999999
	max := -1
	for _, v := range m {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}

func Part1() int {
	lines := strings.Split(input, "\n")
	template := lines[0]
	locations := buildLocations(template)

	rules := map[string]string{}
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if len(line) != 7 {
			panic("unexpected line " + line)
		}
		pair := line[0:2]
		insert := line[6:]
		rules[pair] = insert
		//fmt.Printf("%s -> %s\n", pair, insert)
	}

	for i := 0; i < 10; i++ {
		//fmt.Println("iteration ", i+1)
		locations = buildLocations(template)
		insertions := []insertion{}
		for pair, toInsert := range rules {
			curLocs, ok := locations[pair]
			if ok {
				for _, curLoc := range curLocs {
					insertions = append(insertions, insertion{where: curLoc + 1, what: toInsert})
				}
			}
		}
		sort.Sort(byInsertion(insertions)) // reverse order by insertion location (last to first), to avoid the need
		// for shifting the other insertions
		for _, insertion := range insertions {
			//fmt.Printf("%#v\n", insertion)
			template = template[0:insertion.where] + insertion.what + template[insertion.where:]
			//fmt.Println(template)
		}
		//fmt.Println(template)
	}
	occs := occurrences(template)
	min, max := maxMinOccurrences(occs)
	return max - min
}
