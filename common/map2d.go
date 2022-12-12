package common

import (
	"strings"
)

type Map2D map[Vector]int

func ReadMap(in []byte, offset int) Map2D {
	out := make(Map2D)

	for y, row := range strings.Split(string(in), "\n") {
		for x, h := range row {
			out[Vector{X: x, Y: y}] = int(h) - offset
		}
	}
	return out
}

func (m Map2D) Find(item int) (Vector, bool) {
	for pos, v := range m {
		if v == item {
			return pos, true
		}
	}

	return Vector{}, false
}

func (m Map2D) Copy() Map2D {
	cp := make(Map2D, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func (m Map2D) String(width, height int) string {
	str := ""
	y := 0
	for i := 0; i <= width*height; i++ {
		x := i % width
		if i%width == 0 && i != 0 {
			y++
			str += "\n"
		}

		//fmt.Printf("x: %d, y: %d\n", x, y)
		v, ok := m[Vector{x, y}]
		if !ok {
			str += "."
		} else {
			str += string(rune(v))
		}

	}

	return strings.TrimSpace(str)
}

func (m Map2D) Width() int {
	width := 0

	for v, _ := range m {
		if v.X > width {
			width = v.X
		}
	}

	return width
}
