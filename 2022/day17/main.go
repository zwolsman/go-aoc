package main

import (
	_ "embed"
	"fmt"
	"github.com/zwolsman/go-aoc/common"
	"reflect"
	"sort"
)

//go:embed input.txt
var in []byte

/*
0 1 2 3 4
0 # # # #
*/
var horizontalLineShape = []common.Vector{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 3, Y: 0},
}

/*
0 1 2 3 4
2 . # .
0 # # #
1 . # .
*/
var plusShape = []common.Vector{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 2, Y: 1},
	{X: 1, Y: 2},
}

/*
		0 1 2 3 4
	  2 . . #
	  1 . . #
	  0 # # #
*/
var cornerShape = []common.Vector{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 2, Y: 1},
	{X: 2, Y: 2},
}

/*
	0 1 2 3 4

0 #
1 #
2 #
3 #
*/
var verticalLineShape = []common.Vector{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: 2},
	{X: 0, Y: 3},
}

/*
	0 1 2 3 4

0 # #
1 # #
*/
var squareShape = []common.Vector{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
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
	fmt.Println(run(in, 2022))              // 3177 -> too low
	fmt.Println(run(in, 1_000_000_000_000)) // 1209677419384 --> too low
}

func run(in []byte, rocks int) int {

	jets := string(in)
	shapeIndex, jetIndex := 0, 0

	// initial floor design
	hitboxes := map[common.Vector]any{
		{X: 0, Y: 0}: common.PLACEHOLDER,
		{X: 1, Y: 0}: common.PLACEHOLDER,
		{X: 2, Y: 0}: common.PLACEHOLDER,
		{X: 3, Y: 0}: common.PLACEHOLDER,
		{X: 4, Y: 0}: common.PLACEHOLDER,
		{X: 5, Y: 0}: common.PLACEHOLDER,
		{X: 6, Y: 0}: common.PLACEHOLDER,
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

	patternLocations := make(map[int]int)
	offset := 0

	for i := 1; i <= rocks; i++ {
		fmt.Printf("Placing rock %d\n", i)
		for {
			// being pushed by a jet of hot gas one unit
			next := common.Apply(common.Vector.Plus, shape, nextJet())

			// if within bounds
			if common.MinBy(next, vectorX) >= 0 && common.MaxBy(next, vectorX) <= 6 && !hits(next, hitboxes) {
				shape = next
			}

			//falling one unit down.
			next = common.Apply(common.Vector.Plus, shape, down)

			// oh no, hit into something
			if hits(next, hitboxes) {

				// update hitboxes
				for _, s := range shape {
					highestY = common.Max(highestY, s.Y)

					hitboxes[s] = common.PLACEHOLDER
				}

				for _, s := range shape {
					if hasPatternStart(s.Y, hitboxes) {
						patternLocations[s.Y] = i

						if len(patternLocations) == 5 {
							fmt.Println("Full pattern detected!")
							printMap(hitboxes)
							fmt.Println(patternLocations)

							ys := common.Keys(patternLocations)
							rs := common.Values(patternLocations)

							startY := common.MinArr(ys)
							startRocks := common.MinArr(rs)

							jumpY := s.Y - startY
							jumpRocks := i - startRocks
							left := rocks - i
							times := left / jumpRocks

							offset = jumpY * times

							fmt.Printf("to make a jump of %d rocks you will increase Y by %d (can do this %d times starting at %d)\n", jumpRocks, jumpY, times, s.Y)

							i += jumpRocks * times
							break
						}
					}
				}

				break
			} else {
				shape = next
			}
		}

		shape = nextShape()
	}

	printMap(hitboxes)
	println(offset)

	return highestY + offset
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

func hits(test []common.Vector, hitboxes map[common.Vector]any) bool {
	for _, t := range test {
		if _, ok := hitboxes[t]; ok {
			return true
		}
	}

	return false
}

var patternStart = []int{1, 2, 3, 4, 5}

func hasPatternStart(y int, hitboxes map[common.Vector]any) bool {
	var xs []int
	for c := range hitboxes {
		if c.Y != y {
			continue
		}

		xs = append(xs, c.X)
	}
	sort.Ints(xs)
	return reflect.DeepEqual(xs, patternStart)
}
func vectorX(vec common.Vector) int {
	return vec.X
}

func vectorY(vec common.Vector) int {
	return vec.Y
}
