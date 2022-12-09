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

func run(in []byte, size int) int {
	rope := make([]common.Vector, size)
	history := make(map[common.Vector]int)
	history[rope[size-1]]++

	for _, l := range strings.Split(string(in), "\n") {
		motion := strings.Split(l, " ")
		dir := directions[motion[0]]
		c, _ := strconv.Atoi(motion[1])

		for i := 0; i < c; i++ {
			rope[0] = rope[0].Plus(dir)
			for j := 1; j < size; j++ {
				head := rope[j-1]
				tail := rope[j]

				if head.Dist(tail) > 1 {
					move := head.Min(tail).Normalize()
					rope[j] = tail.Plus(move)
				}
			}
			history[rope[size-1]]++
		}
	}

	return len(history)
}
