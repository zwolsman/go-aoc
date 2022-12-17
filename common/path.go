package common

import "math"

type Path struct {
	Lines [][2]Vector
}

func (p Path) Intersects(other Vector) bool {
	for _, line := range p.Lines {
		a, b := line[0], line[1]

		left := int(math.Min(float64(a.X), float64(b.X)))
		right := int(math.Max(float64(a.X), float64(b.X)))
		top := int(math.Min(float64(a.Y), float64(b.Y)))
		bottom := int(math.Max(float64(a.Y), float64(b.Y)))

		if other.X >= left && other.X <= right && other.Y >= top && other.Y <= bottom {
			return true
		}
	}
	return false
}
