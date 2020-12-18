package main

import (
	intprogram "../program"
	"math"
)

func main() {
	hull := runProgram()
	part1(hull)
}

func runProgram() map[Vector]int {
	program := intprogram.Read("./2019/day11/input.txt")

	in := make(chan int, 1)
	out := make(chan int)
	program.In = in
	program.Out = out

	pos := Vector{}
	hull := make(map[Vector]int)
	dir := Vector{x: 0, y: -1}
	handleIO := func() {
		for {
			color := hull[pos]
			in <- color

			value, direction := <-out, <-out

			hull[pos] = value

			switch direction {
			case 0:
				dir = dir.Rotate(-90)
			case 1:
				dir = dir.Rotate(90)
			}
			pos = pos.plus(dir)
		}
	}

	go handleIO()
	program.Run()
	return hull
}

func part1(hull map[Vector]int) {
	println("Amount of panels painted at least once", len(hull))
}

type Vector struct {
	x, y int
}

func (v Vector) plus(o Vector) Vector {
	return Vector{
		x: v.x + o.x,
		y: v.y + o.y,
	}
}
func (v Vector) Rotate(degrees float64) Vector {
	radians := degrees * (math.Pi / 180)
	ca := math.Cos(radians)
	sa := math.Sin(radians)
	return Vector{int(ca*float64(v.x) - sa*float64(v.y)), int(sa*float64(v.x) + ca*float64(v.y))}
}
