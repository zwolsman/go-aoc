package main

import (
	"bufio"
	"os"
)

func main() {
	//part1()
	part2()
}

func part1() {
	game := New()
	game.Run()

	totalBlocks := 0
	for _, row := range game.screen {
		for _, tile := range row {
			if tile == BLOCK {
				totalBlocks++
			}
		}
	}
	println("total blocks on the screen", totalBlocks)
	println("score", game.score)
}

const (
	left  = "\u001B[D"
	right = "\u001B[C"
)

func part2() {

	in := make(chan int)
	game := New()

	go calculateInput(in, game)
	//go parseInput(in)
	game.RunWithInput(in)
}

func calculateInput(in chan<- int, g *Game) {
	scanner := bufio.NewScanner(os.Stdin)
	paddle, ball := -1, -1
	for scanner.Scan() {
		for _, row := range g.screen {
			for x, tile := range row {
				if tile == BALL {
					ball = x
				}
				if tile == HORIZONTAL_PADDLE {
					paddle = x
				}
			}
		}
		if paddle < ball {
			in <- 1
		}
		if paddle > ball {
			in <- -1
		}
		if paddle == ball {
			in <- 0
		}
	}
}

func parseInput(in chan<- int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case left:
			in <- -1
		case right:
			in <- 1
		default:
			in <- 0
		}
	}
}
