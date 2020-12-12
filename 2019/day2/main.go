package main

import (
	intprogram "../program"
)

func main() {
	program := intprogram.Read("./2019/day2/input.txt")
	part1(program)
	part2(program)
}

func part1(program intprogram.Program) {
	intprogram.Fix(12, 1, program)
	memory := intprogram.Run(program)
	println("The value left at position 0 after the program halts", memory[0])
}

func part2(program intprogram.Program) {
	const target = 19690720

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			intprogram.Fix(n, v, program)
			memory := intprogram.Run(program)
			if memory[0] == target {
				println("the input noun and verb that cause the program to produce the output", target, "are", n, v)
				println("100 *", n, "+", v, "is", 100*n+v)
				return
			}
		}
	}
}
