package day21

type die struct {
	last int
}

func (d *die) roll() int {
	d.last = (d.last % 100) + 1
	return d.last
}

func Part1() int {
	//player1Pos := 3
	//player2Pos := 7
	player1Pos := 5
	player2Pos := 7
	player1Score := 0
	player2Score := 0
	rolls := 0
	d := &die{}
	for i := 0; true; i++ {
		for j := 0; j < 3; j++ {
			player1Pos += d.roll()
			player1Pos %= 10
			rolls++
		}
		player1Score += (player1Pos + 1)
		//fmt.Println("p1", player1Pos+1, player1Score)
		if player1Score >= 1000 {
			//fmt.Println("p1", player2Score, rolls)
			return player2Score * rolls
		}
		for j := 0; j < 3; j++ {
			player2Pos += d.roll()
			player2Pos %= 10
			rolls++
		}
		player2Score += (player2Pos + 1)
		//fmt.Println("p2", player2Pos+1, player2Score)
		if player2Score >= 1000 {
			return player1Score * rolls
		}
	}
	return 0
}
