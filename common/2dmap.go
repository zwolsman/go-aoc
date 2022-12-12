package common

import "strings"

func ReadMap(in []byte, offset rune) map[Vector]int {
	out := make(map[Vector]int)

	for y, row := range strings.Split(string(in), "\n") {
		for x, h := range row {
			out[Vector{X: x, Y: y}] = int(h - offset)
		}
	}
	return out
}
