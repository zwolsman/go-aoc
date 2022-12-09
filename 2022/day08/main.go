package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) int {
	m := readMap(in)

	var walk func(base, dir vector, h int) bool
	walk = func(base, dir vector, h int) bool {
		cur := base.plus(dir)

		v, ok := m[cur]
		if !ok { // got to the edge
			return true
		}

		if v >= h {
			return false
		}

		return walk(cur, dir, h)
	}

	bound := int(math.Sqrt(float64(len(m)))) - 1
	var sum int
	for base, h := range m {
		if base.y == 0 || base.y == bound || base.x == 0 || base.x == bound {
			sum++
			continue
		}

		for _, op := range ops {
			if walk(base, op, h) {
				sum++
				break
			}
		}
	}

	return sum
}

func mask(arr []int) []int {
	out := make([]int, len(arr))

	high := -1
	for i := 0; i < len(arr); i++ {
		if n := arr[i]; n > high {
			out[i] = i
			high = n
		} else {
			for j := i - 1; j > 0; j-- {
				if arr[j] >= n {
					out[i] = i - j - 1
					break
				}
			}
		}
	}
	return out
}

func part2(in []byte) int {
	m := readMap(in)

	var walk func(base, dir vector, h int) int
	walk = func(base, dir vector, h int) int {
		cur := base.plus(dir)

		v, ok := m[cur]
		if !ok { // got to the edge
			return 0
		}

		if v >= h {
			return 1
		}

		return 1 + walk(cur, dir, h)
	}

	top := -1
	for base, h := range m {
		score := 1
		for _, op := range ops {
			score *= walk(base, op, h)
		}

		if score > top {
			top = score
		}
	}

	return top
}

func readMap(in []byte) map[vector]int {
	lines := strings.Split(string(in), "\n")
	out := make(map[vector]int)

	for y, row := range lines {
		for x, h := range row {
			out[vector{x, y}] = int(h - '0')
		}
	}
	return out
}

var ops = []vector{
	{0, 1},  // down
	{0, -1}, // up
	{1, 0},  // right
	{-1, 0}, //left
}

type vector struct {
	x, y int
}

func (v vector) plus(o vector) vector {
	return vector{
		v.x + o.x,
		v.y + o.y,
	}
}

func (v vector) String() string {
	return fmt.Sprintf("vector{x: %d, y: %d}", v.x, v.y)
}
