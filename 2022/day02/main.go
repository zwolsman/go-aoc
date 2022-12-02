package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var in []byte

const (
	WIN  = 0
	DRAW = 1
	LOSS = 2
)

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) any {
	lines := strings.Split(string(in), "\n")
	var score int
	for i := 0; i < len(lines); i++ {
		opponent := int(lines[i][0] - 'A')
		you := int(lines[i][2] - 'X')

		score += calculateScore(opponent, you)
	}
	return score
}

func part2(in []byte) any {
	lines := strings.Split(string(in), "\n")
	var score int
	for i := 0; i < len(lines); i++ {
		opponent := int(lines[i][0] - 'A')
		outcome := lines[i][2] - 'X'

		var you int

		switch outcome {
		case LOSS:
			you = (opponent + 1) % 3
		case DRAW:
			you = opponent
		case WIN:
			you = (opponent + 2) % 3
		}
		score += calculateScore(opponent, you)
	}
	return score
}

func calculateScore(opponent, you int) (score int) {
	score += you + 1

	// Draw
	if opponent == you {
		score += 3
	}

	if you == (opponent+1)%3 {
		score += 6
	}
	return
}
