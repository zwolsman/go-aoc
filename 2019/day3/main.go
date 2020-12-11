package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) Add(a Point) Point {
	return Point{p.x + a.x, p.y + a.y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}
func (p Point) Length() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

type PointArray []Point

func (p PointArray) Len() int           { return len(p) }
func (p PointArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PointArray) Less(i, j int) bool { return p[i].Length() < p[j].Length() }

func test(a1, a2, b1, b2 Point) (Point, error) {

	x1, x2 := a1.x, a2.x
	x3, x4 := b1.x, b2.x

	y1, y2 := a1.y, a2.y
	y3, y4 := b1.y, b2.y

	check := func(x, y int) bool {
		if x == 0 && y == 0 {
			return false
		}
		l1 := int(math.Min(float64(x1), float64(x2)))
		l2 := int(math.Max(float64(x1), float64(x2)))

		l3 := int(math.Min(float64(x3), float64(x4)))
		l4 := int(math.Max(float64(x3), float64(x4)))
		if !(x >= l1 &&
			x <= l2 &&
			x >= l3 &&
			x <= l4) {
			return false
		}

		m1 := int(math.Min(float64(y1), float64(y2)))
		m2 := int(math.Max(float64(y1), float64(y1)))
		m3 := int(math.Min(float64(y3), float64(y4)))
		m4 := int(math.Max(float64(y3), float64(y4)))

		if !(y >= m1 &&
			y <= m2 &&
			y >= m3 &&
			y <= m4) {
			return false
		}
		return true
	}

	x12 := x1 - x2
	x34 := x3 - x4
	y12 := y1 - y2
	y34 := y3 - y4
	c := x12*y34 - y12*x34
	a := x1*y2 - y1*x2
	b := x3*y4 - y3*x4
	if c != 0 {
		x := (a*x34 - b*x12) / c
		y := (a*y34 - b*y12) / c
		if check(x, y) {
			return Point{x, y}, nil
		} else {
			return Point{}, errors.New("point out of bound")
		}
	} else {
		return Point{}, errors.New("lines are parallel")
	}
}

func main() {
	data, err := ioutil.ReadFile("./2019/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var wires []PointArray

	for _, wire := range strings.Split(string(data), "\n") {
		v := Point{0, 0}
		current := []Point{v}
		for _, piece := range strings.Split(wire, ",") {
			length, err := strconv.Atoi(piece[1:])
			if err != nil {
				log.Fatal(err)
			}
			var a Point
			switch piece[0] {
			case 'R':
				a = Point{x: length}
				break
			case 'L':
				a = Point{x: -length}
				break
			case 'D':
				a = Point{y: length}
				break
			case 'U':
				a = Point{y: -length}
				break
			}
			v = v.Add(a)
			current = append(current, v)
		}

		wires = append(wires, current)
	}

	part1(wires)
}

func part1(wires []PointArray) {
	intersections := Intersections(wires...)

	sort.Sort(intersections)
	println(intersections[0].Length())
}

func Intersections(wires ...PointArray) (intersections PointArray) {
	wireA, wireB := wires[0], wires[1]
	for a := 0; a < len(wireA)-1; a++ {
		for b := 0; b < len(wireB)-1; b++ {
			a1, a2, b1, b2 := wireA[a], wireA[a+1], wireB[b], wireB[b+1]
			i1, err := test(a1, a2, b1, b2)
			if err == nil {
				intersections = append(intersections, i1)
			}
			i2, err := test(b1, b2, a1, a2)
			if err == nil && i1 != i2 {
				intersections = append(intersections, i2)
			}
		}
	}
	return
}
