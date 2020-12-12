package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Vector struct {
	x, y float64
}

func (v1 Vector) Mult(amt float64) Vector {
	return Vector{v1.x * amt, v1.y * amt}
}
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}
func (v1 Vector) Min(v2 Vector) Vector {
	return Vector{v1.x - v2.x, v1.y - v2.y}
}
func (v Vector) Rotate(radians float64) Vector {
	ca := math.Cos(radians)
	sa := math.Sin(radians)
	return Vector{ca*v.x - sa*v.y, sa*v.x + ca*v.y}
}
func (v Vector) Length() float64 {
	return math.Abs(v.x) + math.Abs(v.y)
}
func (v Vector) Angle() float64 {
	return math.Atan2(v.y, v.x)
}
func (v Vector) String() string {
	return fmt.Sprintf("Vector(%.f, %.f)", v.x, v.y)
}

func main() {
	data, err := ioutil.ReadFile("./2020/day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(data), "\n")
	println("*** part 1 ***")
	part1(instructions)

	println("*** part 2 ***")
	part2(instructions)
}

func part1(instructions []string) {
	direction := Vector{x: 1} //start facing east
	position := Vector{}

	for _, instruction := range instructions {
		num, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(err)
		}
		arg := float64(num)

		var v Vector
		switch instruction[0] {
		case 'N':
			v = Vector{y: -arg}
			break
		case 'S':
			v = Vector{y: arg}
			break
		case 'E':
			v = Vector{x: arg}
			break
		case 'W':
			v = Vector{x: -arg}
			break
		case 'L': //left degrees
			radians := arg * (math.Pi / 180)
			result := direction.Angle() - radians
			direction = Vector{math.Cos(result), math.Sin(result)}
			break
		case 'R': //right degrees
			radians := arg * (math.Pi / 180)
			result := direction.Angle() + radians
			direction = Vector{math.Cos(result), math.Sin(result)}
			break
		case 'F':
			v = direction.Mult(arg)
			break
		}

		position = position.Add(v)
		println("instruction", instruction, "position:", position.String())
	}

	fmt.Printf("%v, %.f\n\n", position, position.Length())
}

func part2(instructions []string) {
	waypoint := Vector{10, -1}
	ship := Vector{}

	for _, instruction := range instructions {
		num, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(err)
		}
		arg := float64(num)

		var v Vector
		switch instruction[0] {
		case 'N':
			v = Vector{y: -arg}
			break
		case 'S':
			v = Vector{y: arg}
			break
		case 'E':
			v = Vector{x: arg}
			break
		case 'W':
			v = Vector{x: -arg}
			break
		case 'F':
			v = waypoint.Min(ship).Mult(arg)
			ship = ship.Add(v)
			break
		case 'R':
			radians := arg * (math.Pi / 180)
			line := waypoint.Min(ship)
			waypoint = ship.Add(line.Rotate(radians))
			break
		case 'L':
			radians := arg * (math.Pi / 180)
			line := waypoint.Min(ship)
			waypoint = ship.Add(line.Rotate(-radians))
			break
		}

		waypoint = waypoint.Add(v)
		println(instruction, "-> ship", ship.String(), "waypoint", waypoint.String())
	}
	fmt.Printf("%v, %.f\n\n", ship, ship.Length())
}
