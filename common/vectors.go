package common

import (
	"fmt"
	"math"
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

func (v Vector) Dir(o Vector) Vector {
	x := math.Abs(float64(v.X - o.X))
	y := math.Abs(float64(v.Y - o.Y))

	if x > y {
		return Vector{(o.X - v.X) / 2, 0}
	} else {
		return Vector{0, (o.Y - v.Y) / 2}
	}
}
