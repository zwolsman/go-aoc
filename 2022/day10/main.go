package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

var cycles = []int{
	20,
	60,
	100,
	140,
	180,
	220,
	math.MaxInt,
}

func part1(in []byte) int {
	cycle, register, sum := 1, 1, 0

	nextCycle := func() {
		cycle++
		if cycle == cycles[0] {
			multiplier := cycles[0]
			cycles = cycles[1:]
			sum += register * multiplier
		}
	}

	for _, instruction := range strings.Split(string(in), "\n") {
		cmd, arg, _ := strings.Cut(instruction, " ")

		switch cmd {
		case "addx":
			n, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			nextCycle()
			register += n
		case "noop":
			break
		}

		nextCycle()
	}
	return sum
}

func part2(in []byte) any {
	return nil
}
