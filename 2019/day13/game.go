package main

import (
	intprogram "github.com/zwolsman/go-aoc/2019/program"
	"log"
)

type TileKind int

const (
	EMPTY TileKind = iota
	WALL
	BLOCK
	HORIZONTAL_PADDLE
	BALL
)

const (
	width  = 44
	height = 20
)

type Game struct {
	score   int
	screen  [height][width]TileKind
	program *intprogram.Program
}

func New() *Game {
	program := intprogram.Read("./2019/day13/input.txt")
	program.Out = make(chan int)
	game := Game{program: &program}

	go readOutput(program.Out, &game)
	return &game
}

func (g *Game) Run() {
	g.program.Run()
	close(g.program.Out)
}

func (g *Game) RunWithInput(in chan int) {
	g.program.Memory[0] = 2
	g.program.In = in
	g.Run()
}

func readOutput(output <-chan int, g *Game) {
	for {
		x, ok := <-output
		if !ok {
			break
		}
		y, ok := <-output
		if !ok {
			break
		}
		tileId, ok := <-output
		if !ok {
			break
		}
		if x == -1 && y == 0 {
			g.score = tileId
		} else {
			g.screen[y][x] = TileKind(tileId)
		}
		g.render()
	}
}

func (g *Game) render() {
	println("\033[H\033[2J")
	println("score", g.score)
	for _, row := range g.screen {
		str := make([]rune, len(row))
		for x, tile := range row {
			str[x] = getRune(tile)
		}
		println(string(str))
	}
}

func getRune(kind TileKind) rune {
	switch kind {
	case EMPTY:
		return ' '
	case BLOCK:
		return 'x'
	case BALL:
		return 'o'
	case HORIZONTAL_PADDLE:
		return '-'
	case WALL:
		return '.'
	}

	log.Fatal("no rune for", kind)
	return '-'
}
