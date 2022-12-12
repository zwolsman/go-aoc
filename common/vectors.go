package common

import (
	"fmt"
	"math"
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

func (v Vector) Dist(o Vector) float64 {
	x := math.Abs(float64(v.X - o.X))
	y := math.Abs(float64(v.Y - o.Y))
	return math.Max(x, y)
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
