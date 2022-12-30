package main

import (
	_ "embed"
	"fmt"
	"github.com/zwolsman/go-aoc/common"
)

//go:embed input.txt
var in []byte

/*
0 1 2 3 4
0 # # # #
*/
var horizontalLineShape = []common.Vector{
	{0, 0},
	{1, 0},
	{2, 0},
	{3, 0},
}

/*
0 1 2 3 4
2 . # .
0 # # #
1 . # .
*/
var plusShape = []common.Vector{
	{1, 0},
	{0, 1},
	{1, 1},
	{2, 1},
	{1, 2},
}

/*
		0 1 2 3 4
	  2 . . #
	  1 . . #
	  0 # # #
*/
var cornerShape = []common.Vector{
	{0, 0},
	{1, 0},
	{2, 0},
	{2, 1},
	{2, 2},
}

/*
	0 1 2 3 4

0 #
1 #
2 #
3 #
*/
var verticalLineShape = []common.Vector{
	{0, 0},
	{0, 1},
	{0, 2},
	{0, 3},
}

/*
	0 1 2 3 4

0 # #
1 # #
*/
var squareShape = []common.Vector{
	{0, 0},
	{1, 0},
	{0, 1},
	{1, 1},
}

var shapes = [][]common.Vector{
	horizontalLineShape,
	plusShape,
	cornerShape,
	verticalLineShape,
	squareShape,
}

var jetMapping = map[uint8]common.Vector{
	'<': {X: -1},
	'>': {X: 1},
}

var down = common.Vector{Y: -1}

func main() {
	fmt.Println(run(in, 2022)) // 3177 -> too low
	fmt.Println(run(in, 1000000000000))
}

func run(in []byte, rocks int) int {

	jets := string(in)
	shapeIndex, jetIndex := 0, 0

	// initial floor design
	hitboxes := map[common.Vector]any{
		{0, 0}: common.PLACEHOLDER,
		{1, 0}: common.PLACEHOLDER,
		{2, 0}: common.PLACEHOLDER,
		{3, 0}: common.PLACEHOLDER,
		{4, 0}: common.PLACEHOLDER,
		{5, 0}: common.PLACEHOLDER,
		{6, 0}: common.PLACEHOLDER,
	}
	highestY := 0

	nextShape := func() []common.Vector {
		target := shapes[shapeIndex]
		shape := make([]common.Vector, len(target))
		copy(shape, target)

		for i, vec := range shape {
			// Each rock appears so that its left edge is two units away from the left wall and its bottom edge is three units above the highest rock in the room
			shape[i] = vec.Plus(common.Vector{X: 2, Y: highestY + 4})
		}

		shapeIndex = (shapeIndex + 1) % len(shapes)

		return shape
	}

	nextJet := func() common.Vector {
		jet := jetMapping[jets[jetIndex]]
		jetIndex = (jetIndex + 1) % len(jets)

		return jet
	}

	// Initial shape setup
	shape := nextShape()

	for i := 0; i < rocks; i++ {
		for {
			// being pushed by a jet of hot gas one unit
			next := apply(common.Vector.Plus, shape, nextJet())

			// if within bounds
			if common.MinBy(next, vectorX) >= 0 && common.MaxBy(next, vectorX) <= 6 && !hits(next, hitboxes) {
				shape = next
			}

			//falling one unit down.
			next = apply(common.Vector.Plus, shape, down)

			// oh no, hit into something
			if hits(next, hitboxes) {

				// update hitboxes
				for _, s := range shape {
					highestY = common.Max(highestY, s.Y)

					hitboxes[s] = common.PLACEHOLDER
				}

				break
			} else {
				shape = next
			}
		}

		shape = nextShape()
	}

	return highestY
}

func printMap(hitboxes map[common.Vector]any) {

	height := common.MaxBy(common.Keys(hitboxes), vectorY)

	for y := height + 1; y >= 0; y-- {
		for x := -1; x <= 7; x++ {
			coord := common.Vector{X: x, Y: y}

			if y == 0 {
				if x == -1 || x == 7 {
					fmt.Print("+")
				} else {
					fmt.Print("-")
				}
			} else if x == -1 || x == 7 {
				fmt.Print("|")
			} else if _, ok := hitboxes[coord]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func apply[T any](fn func(T, T) T, items []T, arg T) []T {
	result := make([]T, len(items))
	for i, v := range items {
		result[i] = fn(v, arg)
	}
	return result
}

func hits(test []common.Vector, hitboxes map[common.Vector]any) bool {
	for _, t := range test {
		if _, ok := hitboxes[t]; ok {
			return true
		}
	}

	return false
}

func vectorX(vec common.Vector) int {
	return vec.X
}

func vectorY(vec common.Vector) int {
	return vec.Y
}
