package main

import (
	intprogram "../program"
)

func main() {
	program := intprogram.Read("./2019/day9/input.txt")

	program.Run()
	print("done")
}
