package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("/Users/mzwolsman/Developer/go-aoc/day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(data), "\n")
	ptr := 0
	acc := 0
	history := make(map[int]bool)
	for true {
		println(ptr, acc, instructions[ptr])
		if _, ok := history[ptr]; ok {
			break
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

		if ptr > len(instructions) {
			println("no instructions left")
			break
		}
	}
	println(acc)
}
