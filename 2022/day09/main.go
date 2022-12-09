package main

import (
	_ "embed"
	"fmt"
	"github.com/zwolsman/go-aoc/common"
	"strconv"
	"strings"
)

//go:embed input.txt
var in []byte

var directions = map[string]common.Vector{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) int {
	return run(in, 2)
}

func part2(in []byte) any {
	return run(in, 10)
}

const (
	PREV    = 0
	CURRENT = 1
)

func run(in []byte, size int) int {
	rope := make([][2]common.Vector, size)
	history := make(map[common.Vector]int)

	history[rope[0][CURRENT]]++

	for _, l := range strings.Split(string(in), "\n") {
		motion := strings.Split(l, " ")
		dir := directions[motion[0]]

		c, _ := strconv.Atoi(motion[1])

		for i := 0; i < c; i++ {
			// Update head
			rope[size-1][PREV] = rope[size-1][CURRENT]
			rope[size-1][CURRENT] = rope[size-1][CURRENT].Plus(dir)

			for body := size - 1; body > 0; body-- {
				head := rope[body][CURRENT]
				tail := rope[body-1][CURRENT]

				if head.Dist(tail) > 1 {
					rope[body-1][PREV] = rope[body-1][CURRENT] // tail prev = tail current
					rope[body-1][CURRENT] = rope[body][PREV]   // tail current = head current (not updated yet)

					if body-1 == 0 {
						history[rope[0][CURRENT]]++
					}
				}
			}
		}
	}

	return len(history)
}

func run2(in []byte, size int) int {
	for _, l := range strings.Split(string(in), "\n") {
		motion := strings.Split(l, " ")
		dir := directions[motion[0]]

		c, _ := strconv.Atoi(motion[1])

		for i := 0; i < c; i++ {

		}

	}

	return 0
}
