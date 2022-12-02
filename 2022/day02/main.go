package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) any {
	lines := strings.Split(string(in), "\n")
	var score int
	for i := 0; i < len(lines); i++ {
		opponent := lines[i][0] - 'A'
		you := lines[i][2] - 'X'

		score += int(you + 1)
		// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
		// rock == 0
		// paper = 1
		// scisors == 2
		switch {
		case you == 0 && opponent == 2:
			score += 6
		case you == 2 && opponent == 1:
			score += 6
		case you == 1 && opponent == 0:
			score += 6
		case you == opponent:
			score += 3
		}
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
		case 0: // lose
			switch opponent {
			case 0:
				you = 2
			case 1:
				you = 0
			case 2:
				you = 1
			}
		case 1: // draw
			you = int(opponent)
		case 2: // win
			switch opponent {
			case 0:
				you = 1
			case 1:
				you = 2
			case 2:
				you = 0
			}
		}
		score += int(you + 1)
		// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
		// rock == 0
		// paper = 1
		// scisors == 2
		switch {
		case you == 0 && opponent == 2:
			score += 6
		case you == 2 && opponent == 1:
			score += 6
		case you == 1 && opponent == 0:
			score += 6
		case you == opponent:
			score += 3
		}
	}
	return score
}
