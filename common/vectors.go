package common

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	Left  = Vector{1, 0, 0}
	Right = Vector{-1, 0, 0}
	Up    = Vector{0, -1, 0}
	Down  = Vector{0, 1, 0}

	LRUD = [4]Vector{Left, Right, Up, Down}
)

type Vector struct {
	X, Y, Z int
}

func NewVector(in string) (Vector, error) {
	s := strings.Split(in, ",")
	if len(s) < 2 {
		return Vector{}, errors.New("could not parse x,y or x,y,z")
	}

	var x, y, z int
	x, _ = strconv.Atoi(s[0])

	y, _ = strconv.Atoi(s[1])

	if len(s) == 3 {
		z, _ = strconv.Atoi(s[2])
	}

	return Vector{x, y, z}, nil
}

func (v Vector) Plus(o Vector) Vector {
	return Vector{
		v.X + o.X,
		v.Y + o.Y,
		v.Z + o.Z,
	}
}

func (v Vector) Min(o Vector) Vector {
	return Vector{
		v.X - o.X,
		v.Y - o.Y,
		v.Z - o.Z,
	}
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector{X: %d, Y: %d, Z: %d}", v.X, v.Y, v.Z)
}

func (v Vector) Times(n int) Vector {
	return Vector{
		v.X * n,
		v.Y * n,
		v.Z * n,
	}
}

// Dist returns the Manhattan distance from v to o
// |X1 – X2| + |Y1 – Y2|
func (v Vector) Dist(o Vector) int {
	return int(math.Abs(float64(o.X-v.X)) + math.Abs(float64(o.Y-v.Y)))
}

func (v Vector) Span(radius, y int) (Vector, int) {

	target := Vector{
		X: v.X,
		Y: y,
	}

	dist := v.Dist(target)

	span := 0
	if dist == 0 {
		span = radius*2 + 1
	}

	if dist <= radius {
		span = (radius-dist)*2 + 1
	}

	if span == 0 {
		return Vector{}, 0
	}

	start := Vector{
		X: v.X - (span-1)/2,
		Y: y,
	}

	return start, span
}

func (v Vector) Normalize() Vector {
	o := Vector{v.X, v.Y, v.Z}
	if o.X < -1 {
		o.X = -1
	}
	if o.X > 1 {
		o.X = 1
	}
	if o.Y < -1 {
		o.Y = -1
	}
	if o.Y > 1 {
		o.Y = 1
	}
	if o.Z > 1 {
		o.Z = 1
	}
	return o
}

func (v Vector) Copy() Vector {
	return Vector{v.X, v.Y, v.Z}
}
