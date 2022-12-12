package main

import (
	_ "embed"
	"fmt"
	"github.com/zwolsman/go-aoc/common"
	"math"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) int {
	m := common.ReadMap(in, '0')

	var walk func(base, dir common.Vector, h int) bool
	walk = func(base, dir common.Vector, h int) bool {
		cur := base.Plus(dir)

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
		if base.Y == 0 || base.Y == bound || base.X == 0 || base.X == bound {
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
	m := common.ReadMap(in, '0')

	var walk func(base, dir common.Vector, h int) int
	walk = func(base, dir common.Vector, h int) int {
		cur := base.Plus(dir)

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

var ops = []common.Vector{
	{0, 1},  // down
	{0, -1}, // up
	{1, 0},  // right
	{-1, 0}, //left
}
