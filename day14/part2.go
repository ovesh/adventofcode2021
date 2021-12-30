package day14

import (
	"strings"
)

// The trick here is that you don't actually need the string at all.
// All you need to know is that every time you match on a pair AB
// with rule AB->C, you get 2 new pairs AC and CB and you don't
// care about the pairs from the previous iteration.
func Part2() int {
	lines := strings.Split(input, "\n")
	template := lines[0]
	countsPerChar := map[string]int{}
	countsPerPair := map[string]int{}
	for i := 0; i < len(template); i++ {
		c := template[i : i+1]
		countsPerChar[c] = countsPerChar[c] + 1
		if i < len(template)-1 {
			pair := template[i : i+2]
			countsPerPair[pair] = countsPerPair[pair] + 1
		}
	}

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

	for i := 0; i < 40; i++ {
		allPairCountChanges := []map[string]int{}
		//fmt.Println("iteration ", i+1)
		for pair, toInsert := range rules {
			pairA := pair[:1] + toInsert
			pairB := toInsert + pair[1:]
			countsOrig := countsPerPair[pair]
			allPairCountChanges = append(allPairCountChanges, map[string]int{pairA: countsOrig, pairB: countsOrig})
			countsPerChar[toInsert] = countsPerChar[toInsert] + countsOrig
		}
		countsPerPair = map[string]int{}
		for _, changes := range allPairCountChanges {
			for pair, change := range changes {
				countsPerPair[pair] = countsPerPair[pair] + change
			}
		}
		//fmt.Printf("%#v\n", countsPerPair)
	}
	//fmt.Printf("%#v\n", countsPerPair)
	//fmt.Printf("%#v\n", countsPerChar)
	min, max := maxMinOccurrences(countsPerChar)
	return max - min
}
