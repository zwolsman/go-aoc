package program

import (
	"fmt"
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
	ptr := 0

	readArgs := func(mask int, num int, modes Modes) []int {
		args := make([]int, num)
		for i := 0; i < num; i++ {
			arg := memory[ptr+i]
			if modes[len(modes)-i-1] == 0 {
				arg = memory[arg]
			}
			args[i] = arg
		}
		ptr += num
		return args
	}
	read := func() int {
		result := memory[ptr]
		ptr++
		return result
	}

	for memory[ptr] != 99 {
		mask := read()
		opcode := readOpcode(mask)

		switch opcode {
		case 1:
			modes := ReadModes(mask, 3)
			args := readArgs(mask, 2, modes)
			x, y, target := args[0], args[1], read()
			memory[target] = x + y
			break
		case 2:
			modes := ReadModes(mask, 3)
			args := readArgs(mask, 2, modes)
			x, y, target := args[0], args[1], read()
			memory[target] = x * y
			break
		case 3: //INPUT
			target := read()
			var input int
			print("input: ")
			_, err := fmt.Scan(&input)
			if err != nil {
				log.Fatal(err)
			}
			memory[target] = input
		case 4: //OUTPUT
			modes := ReadModes(mask, 1)
			args := readArgs(mask, 1, modes)
			println(args[0])
		default:
			println("weird opcode", ptr, opcode)
		}
	}

	return memory
}

func readOpcode(mask int) int {
	if mask < 10 {
		return mask
	}
	str := fmt.Sprint(mask)
	opcode, err := strconv.Atoi(str[len(str)-2:])
	if err != nil {
		log.Fatal(err)
	}
	return opcode
}

type Modes []int32

func ReadModes(mask int, num int) Modes {
	correct := make([]int32, num+2)

	str := fmt.Sprintf("%d", mask)
	for i, c := range str {
		correct[len(correct)-len(str)+i] = c
	}
	correct = correct[0:num]
	for i, v := range correct {
		if v != 0 {
			correct[i] = v - '0'
		}
	}
	return correct
}
