package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2020/day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(data), "\n")
	part1(instructions)
	part2(instructions)
}

func part1(instructions []string) {
	acc, _ := runProgram(instructions)
	println(acc)
}

func part2(instructions []string) {
	for i := 0; i < len(instructions); i++ {
		newInstructions := mutateInstructions(instructions, i)
		acc, err := runProgram(newInstructions)
		if err == nil {
			println(acc)
			break
		}
	}
}

var replacer = strings.NewReplacer("jmp", "nop", "nop", "jmp")

func mutateInstructions(instructions []string, offset int) []string {
	mutatedInstructions := make([]string, len(instructions))
	copy(mutatedInstructions, instructions)

	for i := len(mutatedInstructions) - 1; i != -1; i-- {
		instruction := mutatedInstructions[i]

		if instruction[:3] == "jmp" {
			offset--
		}
		if instruction[:3] == "nop" {
			offset--
		}

		if offset == -1 { //do mutation
			mutatedInstructions[i] = replacer.Replace(instruction)
			return mutatedInstructions
		}
	}
	panic("NOTHING HAS BEEN REPLACED!")
}

func runProgram(instructions []string) (acc int, err error) {
	ptr := 0
	history := make(map[int]bool)
	for true {
		//println(ptr, acc, instructions[ptr])
		if _, ok := history[ptr]; ok {
			return acc, errors.New("infinite loop")
		}
		if ptr > len(instructions)-1 {
			return acc, errors.New("index out of range [626] with length 626")
		}
		history[ptr] = true
		instruction := instructions[ptr][:3]
		args := instructions[ptr][4:]

		intArg := func() int {
			num, err := strconv.ParseInt(args, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			return int(num)
		}

		switch instruction {
		case "nop":
			ptr++
			break
		case "acc":
			acc += intArg()
			ptr++
			break
		case "jmp":
			ptr += intArg()
			break
		}

		if ptr > len(instructions)-1 {
			break
		}
	}
	return acc, nil
}
