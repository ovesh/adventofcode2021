package day12

import (
	"strings"
)

func buildTree2(lines []string) (*vertex, []string) {
	allNodes := map[string]*vertex{}
	root := vertex{name: "start", neighbors: []*vertex{}}
	end := vertex{name: "end", neighbors: []*vertex{}}
	allNodes["start"] = &root
	allNodes["end"] = &end
	for _, line := range lines {
		split := strings.Split(line, "-")
		nameLeft := split[0]
		nameRight := split[1]

		var left *vertex
		var ok bool
		if left, ok = allNodes[nameLeft]; !ok {
			left = &vertex{name: nameLeft, neighbors: []*vertex{}}
			allNodes[nameLeft] = left
		}
		var right *vertex
		if right, ok = allNodes[nameRight]; !ok {
			right = &vertex{name: nameRight, neighbors: []*vertex{}}
			allNodes[nameRight] = right
		}
		if !left.isNeighbor(nameRight) {
			left.neighbors = append(left.neighbors, right)
		}
		if !right.isNeighbor(nameLeft) {
			right.neighbors = append(right.neighbors, left)
		}
	}

	allSmallNodes := []string{}
	for v := range allNodes {
		if v == "start" || v == "end" {
			continue
		}
		if allNodes[v].traversable() {
			continue
		}
		allSmallNodes = append(allSmallNodes, v)
	}

	return &root, allSmallNodes
}

func (v *vertex) traverse2(trailSoFar []string, allowedTwice string) [][]string {
	if v.name == "end" {
		return [][]string{{"end"}}
	}
	times := 0
	for _, name := range trailSoFar {
		if name == v.name && !v.traversable() { // already traversed
			times++
			if v.name != allowedTwice {
				return [][]string{}
			}
			if times == 2 {
				return [][]string{}
			}
		}
	}
	res := [][]string{}
	for _, neighbor := range v.neighbors {
		childTrails := neighbor.traverse2(append(trailSoFar, v.name), allowedTwice)
		for _, childTrail := range childTrails {
			res = append(res, append([]string{v.name}, childTrail...))
		}
	}
	//fmt.Println(res)
	return res
}

func Part2() int {
	lines := strings.Split(input, "\n")
	root, allSmallNodes := buildTree2(lines)
	allPaths := map[string]bool{}
	for _, smallNode := range allSmallNodes {
		paths := root.traverse2([]string{}, smallNode)
		for _, path := range paths {
			// dedupe paths
			allPaths[strings.Join(path, ",")] = true
		}
	}
	return len(allPaths)
}
