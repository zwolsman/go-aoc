package common

import "fmt"

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
