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

func part2(in []byte) string {
	program := run(in)
	var out string
	for i := 0; i < 40*6; i++ {
		x := i % 40

		if x == 0 {
			out += "\n"
		}

		position := read(i+1, program)

		if x >= position-1 && x <= position+1 {
			out += "#"
		} else {
			out += "."
		}
	}

	return out
}

type program = map[int]int

func read(cycle int, p program) int {
	n, ok := p[cycle]
	if !ok {
		n = p[cycle-1]
	}
	return n
}

func run(in []byte) program {
	cycle := 1
	register := 1
	p := make(program)

	p[cycle] = register
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

		p[cycle] = register
	}

	return p
}
