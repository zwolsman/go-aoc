package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Program []int

func main() {
	program := readProgram("./2019/day2/input.txt")
	part1(program)
	part2(program)
}

func part1(program Program) {
	fixProgram(12, 1, program)
	memory := runProgram(program)
	println("The value left at position 0 after the program halts", memory[0])
}

func part2(program Program) {
	const target = 19690720

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			fixProgram(n, v, program)
			memory := runProgram(program)
			if memory[0] == target {
				println("the input noun and verb that cause the program to produce the output", target, "are", n, v)
				println("100 *", n, "+", v, "is", 100*n+v)
				return
			}
		}
	}
}

func readProgram(path string) (program Program) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, code := range strings.Split(string(data), ",") {
		opCode, err := strconv.Atoi(code)
		if err != nil {
			log.Fatal("error reading opcode", err, code)
		}
		program = append(program, opCode)
	}

	return
}

func fixProgram(noun, verb int, p Program) Program {
	p[1] = noun
	p[2] = verb
	return p
}

func runProgram(program Program) Program {
	memory := make(Program, len(program))
	copy(memory, program)

	for ptr := 0; program[ptr] != 99; {
		opcode := program[ptr]

		switch opcode {
		case 1:
			x, y, target := memory[ptr+1], memory[ptr+2], memory[ptr+3]
			memory[target] = memory[x] + memory[y]

			ptr += 4
			break
		case 2:
			x, y, target := memory[ptr+1], memory[ptr+2], memory[ptr+3]
			memory[target] = memory[x] * memory[y]

			ptr += 4
			break
		default:
			println("weird opcode", ptr, opcode)
		}
	}

	return memory
}
