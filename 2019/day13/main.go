package main

import (
	intprogram "../program"
	"log"
)

func main() {
	program := intprogram.Read("./2019/day13/input.txt")
	out := make(chan int)
	program.Out = out

	go func() {
		program.Run()
		close(program.Out)
	}()

	totalBlocks := 0
	for {
		x, ok := <-out
		if !ok {
			break
		}
		y, ok := <-out
		if !ok {
			break
		}
		tileId, ok := <-out
		if !ok {
			break
		}

		if tileId == 2 {
			totalBlocks++
		}
		log.Printf("x: %d, y: %d, tileId: %d\n", x, y, tileId)
	}

	println("total blocks on the screen", totalBlocks)
}
