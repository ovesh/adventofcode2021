package day21

import (
	"fmt"
)

var memo = map[string][2]int{}

func continueGame(rollCount, p1score, p2score, p1pos, p2pos, dieVal int) (int, int) {
	//fmt.Println("rollCount", rollCount)
	roll1From6Idx := rollCount % 6
	memoKey := fmt.Sprintf("%d:%d:%d:%d:%d:%d", roll1From6Idx, p1pos, p1score, p2pos, p2score, dieVal)
	if memoized, ok := memo[memoKey]; ok {
		return memoized[0], memoized[1]
	}
	if roll1From6Idx < 3 { // player 1
		//fmt.Println("p1 turn, die", dieVal)
		p1pos += dieVal
		if roll1From6Idx == 2 {
			p1pos %= 10
			p1score += (p1pos + 1)
		}
	} else { // player 2
		//fmt.Println("p2 turn, die", dieVal)
		p2pos += dieVal
		if roll1From6Idx == 5 {
			p2pos %= 10
			p2score += (p2pos + 1)
		}
	}
	if roll1From6Idx == 2 {
		if p1score >= 21 {
			//fmt.Println("p1 wins", p1score, ":", p2score)
			return 1, 0
		}
	}
	if roll1From6Idx == 5 {
		if p2score >= 21 {
			//fmt.Println("p2 wins", p2score, ":", p1score)
			return 0, 1
		}
	}

	p1wins, p2wins := continueGame(rollCount+1, p1score, p2score, p1pos, p2pos, 1)
	p1, p2 := continueGame(rollCount+1, p1score, p2score, p1pos, p2pos, 2)
	p1wins += p1
	p2wins += p2
	p1, p2 = continueGame(rollCount+1, p1score, p2score, p1pos, p2pos, 3)
	p1wins += p1
	p2wins += p2
	memo[memoKey] = [2]int{p1wins, p2wins}
	return p1wins, p2wins
}

func Part2() int {
	p1wins, p2wins := continueGame(0, 0, 0, 5, 7, 1)
	p1, p2 := continueGame(0, 0, 0, 5, 7, 2)
	p1wins += p1
	p2wins += p2
	p1, p2 = continueGame(0, 0, 0, 5, 7, 3)
	p1wins += p1
	p2wins += p2
	fmt.Println("p1", p1wins, "p2", p2wins)
	if p1wins > p2wins {
		return p1wins
	}
	return p2wins
}
