package main

import (
	intprogram "../program"
	"log"
	"sync"
)

func main() {
	program := intprogram.Read("./2019/day13/input.txt")
	out := make(chan int)
	program.Out = out

	wg := sync.WaitGroup{}
	wg.Add(1)
	totalBlocks := 0
	go func() {
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
		log.Println("done")
		wg.Done()
	}()

	program.Run()
	close(out)
	wg.Wait()
	println("total blocks on the screen", totalBlocks)
}
