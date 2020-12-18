package main

import (
	intprogram "../program"
	"math"
)

const (
	black = iota
	white
)

func main() {
	hull := runProgram()
	part1(hull)
	part2(hull)
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
			color, ok := hull[pos]
			if !ok {
				in <- white
			} else {
				in <- color
			}

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

func part2(hull map[Vector]int) {
	mx, my := 0, 0
	for v, _ := range hull {
		if v.x > mx {
			mx = v.x
		}
		if v.y > my {
			my = v.y
		}
	}

	for y := 0; y <= my; y++ {
		for x := 0; x <= mx; x++ {
			v := Vector{x, y}
			color, ok := hull[v]
			if !ok {
				print(" ")
				continue
			}
			if color == 1 {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}
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
