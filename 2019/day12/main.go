package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	moons := readInput()
	print := func() {
		for _, m := range moons {
			fmt.Println(m)
		}
		fmt.Println()
	}

	for i := 0; i < 1000; i++ {
		simulate(moons)
		fmt.Printf("After %d steps:\n", i+1)
		print()
	}

	energy := 0
	for _, m := range moons {
		energy += m.Energy()
	}
	fmt.Printf("Total energy in the system: %d", energy)
}

func simulate(moons []*Moon) {
	for _, moon := range moons {
		for _, other := range moons {
			if moon == other {
				continue
			}

			if moon.pos.x > other.pos.x {
				moon.vel.x--
			}
			if moon.pos.x < other.pos.x {
				moon.vel.x++
			}

			if moon.pos.y > other.pos.y {
				moon.vel.y--
			}
			if moon.pos.y < other.pos.y {
				moon.vel.y++
			}

			if moon.pos.z > other.pos.z {
				moon.vel.z--
			}
			if moon.pos.z < other.pos.z {
				moon.vel.z++
			}

		}
	}
	for _, m := range moons {
		m.Invalidate()
	}
}

func readInput() (output []*Moon) {
	data, err := ioutil.ReadFile("./2019/day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	regex, err := regexp.Compile("[xyz]=(-?\\d+)")
	if err != nil {
		log.Fatal(err)
	}

	for _, input := range strings.Split(string(data), "\n") {
		matches := regex.FindAllStringSubmatch(input, -1)
		x, y, z := matches[0][1], matches[1][1], matches[2][1]
		output = append(output, NewMoon(x, y, z))
	}
	return
}

func NewMoon(x, y, z string) *Moon {
	i, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	j, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}
	k, err := strconv.Atoi(z)
	if err != nil {
		log.Fatal(err)
	}
	pos := Vector{i, j, k}
	return &Moon{
		pos: &pos,
		vel: &Vector{},
	}
}

type Vector struct {
	x, y, z int
}

type Moon struct {
	pos *Vector
	vel *Vector
}

func (v *Vector) String() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", v.x, v.y, v.z)
}
func (m *Moon) String() string {
	return fmt.Sprintf("pos=%v, vel=%v", m.pos, m.vel)
}
func (v *Vector) plus(other *Vector) *Vector {
	return &Vector{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v *Vector) Length() int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)) + math.Abs(float64(v.z)))
}

func (m *Moon) Invalidate() {
	m.pos = m.pos.plus(m.vel)
}

func (m *Moon) Energy() int {
	return m.pos.Length() * m.vel.Length()
}
