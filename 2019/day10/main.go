package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2019/day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var asteroids []Vector
	for y, row := range strings.Split(string(data), "\n") {
		for x, c := range row {
			if c != '#' {
				continue
			}

			asteroids = append(asteroids, Vector{float64(x), float64(y)})
		}
	}

	part1(asteroids)
}

func part1(asteroids []Vector) {
	best, pos := 0, Vector{}
	for i := 0; i < len(asteroids); i++ {
		base := asteroids[i]
		angles := make(map[float64]bool)
		for j := 0; j < len(asteroids); j++ {
			if j == i {
				continue
			}

			asteroid := asteroids[j]
			dist := asteroid.min(base)
			angles[dist.Angle()] = true
		}

		if val := len(angles); val > best {
			pos = base
			best = val
		}
	}
	fmt.Printf("Best position is %v where we can see %d other astroids in direct line of sight.\n", pos, best)
}

type Vector struct {
	x, y float64
}

func (v Vector) min(o Vector) Vector {
	return Vector{
		x: v.x - o.x,
		y: v.y - o.y,
	}
}

func (v Vector) Angle() float64 {
	return math.Atan2(v.y, v.x)
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector(%.f, %.f)", v.x, v.y)
}
