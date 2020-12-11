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
	fixProgram(program)
	part1(program)
}

func part1(program Program) {
	program = runProgram(program)
	println("The value left at position 0 after the program halts", program[0])
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

func fixProgram(p Program) Program {
	p[1] = 12
	p[2] = 2
	return p
}

func runProgram(program Program) Program {
	p := make(Program, len(program)) //stop mutating the original state
	copy(p, program)

	for pos := 0; p[pos] != 99; {
		opcode := p[pos]

		switch opcode {
		case 1:
			x, y, target := p[pos+1], p[pos+2], p[pos+3]
			p[target] = p[x] + p[y]

			pos += 4
			break
		case 2:
			x, y, target := p[pos+1], p[pos+2], p[pos+3]
			p[target] = p[x] * p[y]

			pos += 4
			break
		default:
			println("weird opcode", pos, opcode)
		}
	}

	return p
}
