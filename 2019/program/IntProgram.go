package program

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Program []int

func Read(path string) (program Program) {
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

func Fix(noun, verb int, p Program) Program {
	p[1] = noun
	p[2] = verb
	return p
}

func Run(program Program) Program {
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
