package common

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	Left  = Vector{1, 0}
	Right = Vector{-1, 0}
	Up    = Vector{0, -1}
	Down  = Vector{0, 1}

	LRUD = [4]Vector{Left, Right, Up, Down}
)

type Vector struct {
	X, Y int
}

func NewVector(in string) (Vector, error) {
	xstr, ystr, correct := strings.Cut(in, ",")
	if !correct {
		return Vector{}, errors.New("could not parse x,y")
	}

	x, err := strconv.Atoi(xstr)
	if err != nil {
		return Vector{}, err
	}

	y, err := strconv.Atoi(ystr)
	if err != nil {
		return Vector{}, err
	}

	return Vector{x, y}, nil
}

func (v Vector) Plus(o Vector) Vector {
	return Vector{
		v.X + o.X,
		v.Y + o.Y,
	}
}

func (v Vector) Min(o Vector) Vector {
	return Vector{
		v.X - o.X,
		v.Y - o.Y,
	}
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector{X: %d, Y: %d}", v.X, v.Y)
}

func (v Vector) Times(n int) Vector {
	return Vector{
		v.X * n,
		v.Y * n,
	}
}

// Dist returns the Manhattan distance from v to o
// |X1 – X2| + |Y1 – Y2|
func (v Vector) Dist(o Vector) int {
	return int(math.Abs(float64(o.X-v.X)) + math.Abs(float64(o.Y-v.Y)))
}

func (v Vector) Normalize() Vector {
	o := Vector{v.X, v.Y}
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
	return o
}

func (v Vector) Copy() Vector {
	return Vector{v.X, v.Y}
}
