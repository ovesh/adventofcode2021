package day12

import (
	"fmt"
	"strings"
)

type vertex struct {
	name      string
	neighbors []*vertex
}

func (v *vertex) traversable() bool {
	return v.name[0] >= 'A' && v.name[0] <= 'Z'
}

func (v *vertex) isNeighbor(name string) bool {
	for _, neighbor := range v.neighbors {
		if neighbor.name == name {
			return true
		}
	}
	return false
}

func buildTree(lines []string) *vertex {
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

	return &root
}

func (v *vertex) String() string {
	res := "["
	res += fmt.Sprintf("(%s) ", v.name)
	for _, n := range v.neighbors {
		res += n.name + " "
	}
	res += "]"
	return res
}

func (v *vertex) traverse(trailSoFar []string) [][]string {
	if v.name == "end" {
		return [][]string{{"end"}}
	}
	for _, name := range trailSoFar {
		if name == v.name && !v.traversable() { // already traversed
			return [][]string{}
		}
	}
	res := [][]string{}
	for _, neighbor := range v.neighbors {
		childTrails := neighbor.traverse(append(trailSoFar, v.name))
		for _, childTrail := range childTrails {
			res = append(res, append([]string{v.name}, childTrail...))
		}
	}
	//fmt.Println(res)
	return res
}

func Part1() int {
	lines := strings.Split(input, "\n")
	root := buildTree(lines)
	return len(root.traverse([]string{}))
}
