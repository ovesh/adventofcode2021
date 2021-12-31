package day18

import (
	"regexp"
	"strconv"
	"strings"
)

type pair struct {
	a             *pair
	b             *pair
	val           int
	parent        *pair
	depth         int
	leftNeighbor  *pair
	rightNeighbor *pair
}

func (p *pair) String() string {
	if p.a == nil {
		return strconv.Itoa(p.val)
	}
	return "[" + p.a.String() + "," + p.b.String() + "]"
}

var num = regexp.MustCompile("^\\d$")

const noVal = -1

var orderedLeaves = []*pair{}

func parseRoot(s string) *pair {
	orderedLeaves = []*pair{}
	return parse(s, nil)
}

func parse(s string, parent *pair) *pair {
	if s[:1] != "[" { // has to be a number
		if !num.MatchString(s[:1]) {
			panic("expected number")
		}
		val, _ := strconv.Atoi(s[:1])
		res := &pair{val: val, parent: parent, depth: parent.depth + 1}
		return res
	}
	// find middle comma
	inner := s[1 : len(s)-1]
	depth := 0
	for i, c := range inner {
		if string(c) == "[" {
			depth++
		}
		if string(c) == "]" {
			depth--
		}
		if depth == 0 && string(c) == "," {
			res := &pair{parent: parent, val: noVal}
			if parent != nil {
				res.depth = parent.depth + 1
			} else {
				res.depth = 0
			}
			res.a = parse(inner[:i], res)
			if res.a.val != noVal {
				orderedLeaves = append(orderedLeaves, res.a)
				if len(orderedLeaves) > 1 {
					res.a.leftNeighbor = orderedLeaves[len(orderedLeaves)-2]
					orderedLeaves[len(orderedLeaves)-2].rightNeighbor = res.a
				}
			}
			res.b = parse(inner[i+1:], res)
			if res.b.val != noVal {
				orderedLeaves = append(orderedLeaves, res.b)
				res.b.leftNeighbor = orderedLeaves[len(orderedLeaves)-2]
				orderedLeaves[len(orderedLeaves)-2].rightNeighbor = res.b
			}
			return res
		}
	}
	return nil
}

func (p *pair) rightMostLeaf() *pair {
	if p.b.val != noVal {
		return p.b
	}
	return p.b.rightMostLeaf()
}

func (p *pair) leftMostLeaf() *pair {
	if p.a.val != noVal {
		return p.a
	}
	return p.a.leftMostLeaf()
}

func (a *pair) add(b *pair) *pair {
	res := &pair{a: a, b: b, val: noVal, depth: -1}
	a.parent = res
	b.parent = res
	aRightMost := a.rightMostLeaf()
	bLeftMost := b.leftMostLeaf()
	aRightMost.rightNeighbor = bLeftMost
	bLeftMost.leftNeighbor = aRightMost
	res.plusDepth()
	//fmt.Println("before reductions:", res)
	for true {
		exploder := res.findExploder()
		if exploder != nil {
			exploder.explode()
			//fmt.Println("after explode", res)
			continue
		}
		splitter := res.findSplitter()
		if splitter != nil {
			splitter.split()
			//fmt.Println("after split", res)
			continue
		}
		break
	}
	return res
}

func (p *pair) findExploder() *pair {
	if p.a == nil || p.b == nil {
		return nil
	}
	if p.depth >= 4 && p.a.val != noVal && p.b.val != noVal {
		return p
	}
	if p.a != nil {
		res := p.a.findExploder()
		if res != nil {
			return res
		}
	}
	if p.b != nil {
		res := p.b.findExploder()
		if res != nil {
			return res
		}
	}
	return nil
}

func (p *pair) plusDepth() {
	p.depth++
	if p.a != nil {
		p.a.plusDepth()
	}
	if p.b != nil {
		p.b.plusDepth()
	}
}

func (p *pair) explode() {
	//fmt.Println("exploding!", p)
	leftNeighbor := p.a.leftNeighbor
	if leftNeighbor != nil {
		leftNeighbor.val += p.a.val
		leftNeighbor.rightNeighbor = p
		p.leftNeighbor = leftNeighbor
	}
	rightNeighbor := p.b.rightNeighbor
	if rightNeighbor != nil {
		rightNeighbor.val += p.b.val
		rightNeighbor.leftNeighbor = p
		p.rightNeighbor = rightNeighbor
	}
	p.a = nil
	p.b = nil
	p.val = 0
}

func (p *pair) findSplitter() *pair {
	if p.val >= 10 {
		return p
	}
	if p.a != nil {
		res := p.a.findSplitter()
		if res != nil {
			return res
		}
	}
	if p.b != nil {
		res := p.b.findSplitter()
		if res != nil {
			return res
		}
	}
	return nil
}

func (p *pair) split() {
	//fmt.Println("splitting!", p)
	if p.val == noVal {
		panic("split() only expected on value nodes")
	}
	p.a = &pair{val: p.val / 2, parent: p, depth: p.depth + 1}
	if p.leftNeighbor != nil {
		p.leftNeighbor.rightNeighbor = p.a
		p.a.leftNeighbor = p.leftNeighbor
		p.leftNeighbor = nil
	}
	bVal := p.a.val
	if p.val%2 != 0 {
		bVal++
	}
	p.b = &pair{val: bVal, parent: p, depth: p.depth + 1, leftNeighbor: p.a}
	if p.rightNeighbor != nil {
		p.rightNeighbor.leftNeighbor = p.b
		p.b.rightNeighbor = p.rightNeighbor
		p.rightNeighbor = nil
	}
	p.a.rightNeighbor = p.b
	p.val = noVal
}

func (p *pair) magnitude() int {
	if p.val != noVal {
		return p.val
	}
	return 3*p.a.magnitude() + 2*p.b.magnitude()
}

func Part1() int {
	lines := strings.Split(input, "\n")
	var sum *pair
	for i, line := range lines {
		curPair := parseRoot(line)
		//fmt.Println("parsed line:", curPair)
		if i == 0 {
			sum = curPair
			continue
		}
		sum = sum.add(curPair)
		//fmt.Println("after adding", sum)
	}
	return sum.magnitude()
}
