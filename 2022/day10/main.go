package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) int {
	program := run(in)
	var sum int

	for i := 0; i <= 5; i++ {
		c := 20 + (40 * i)

		n, ok := program[c]
		if !ok {
			n = program[c-1]
		}

		sum += n * c
	}

	return sum
}

func part2(in []byte) any {
	return nil
}

func run(in []byte) map[int]int {
	cycle := 1
	register := 1
	program := make(map[int]int)

	for _, instruction := range strings.Split(string(in), "\n") {
		cmd, arg, _ := strings.Cut(instruction, " ")

		switch cmd {
		case "addx":
			n, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			register += n
			cycle += 2
		case "noop":
			cycle++
			break
		}

		program[cycle] = register
	}

	return program
}
