package main

import (
	intprogram "../program"
)

func main() {
	program := intprogram.Read("./2019/day5/input.txt")

	intprogram.Run(program)
	print("done")
}
