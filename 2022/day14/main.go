package main

import (
	_ "embed"
	"fmt"
	"github.com/zwolsman/go-aoc/common"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

var sandSpawn = common.Vector{X: 500}

func part1(in []byte) any {
	sand := sandSpawn.Copy()

	rocks, highestY := readMap(in)
	placedSand := make(map[common.Vector]int)

	for sand.Y <= highestY+1 {
		next, found := findNextOption(sand, placedSand, func(pos common.Vector) bool {
			for _, path := range rocks {
				if path.Intersects(pos) {
					return false
				}
			}
			return true
		})

		if !found {
			placedSand[sand]++
			sand = sandSpawn.Copy()
			continue
		}

		sand = next
	}

	return len(placedSand)
}

func part2(in []byte) int {
	sand := sandSpawn.Copy()

	rocks, highestY := readMap(in)
	placedSand := make(map[common.Vector]int)

	for {
		next, found := findNextOption(sand, placedSand, func(pos common.Vector) bool {
			if pos.Y == highestY+2 {
				return false
			}

			for _, path := range rocks {
				if path.Intersects(pos) {
					return false
				}
			}
			return true
		})

		if !found {
			placedSand[sand]++
			if sand == sandSpawn {
				break
			}
			sand = sandSpawn.Copy()
			continue
		}

		sand = next
	}

	return len(placedSand)
}

func readMap(in []byte) ([]common.Path, int) {
	var rocks []common.Path
	highestY := 0
	for _, line := range strings.Split(string(in), "\n") {
		spots := strings.Split(line, " -> ")

		rock := common.Path{}
		for i := 0; i < len(spots)-1; i++ {
			a, b := spots[i], spots[i+1]

			vecA, err := common.NewVector(a)
			if err != nil {
				panic(err)
			}

			vecB, err := common.NewVector(b)

			if vecA.Y > highestY {
				highestY = vecA.Y
			}

			if vecB.Y > highestY {
				highestY = vecB.Y
			}

			if err != nil {
				panic(err)
			}

			rock.Lines = append(rock.Lines, [2]common.Vector{vecA, vecB})
		}

		rocks = append(rocks, rock)
	}
	return rocks, highestY
}

var options = [3]common.Vector{
	{0, 1},
	{-1, 1},
	{1, 1},
}

func findNextOption(origin common.Vector, occupied map[common.Vector]int, isOption func(pos common.Vector) bool) (common.Vector, bool) {
	for _, option := range options {
		newPos := origin.Plus(option)
		if _, ok := occupied[newPos]; ok {
			continue
		}

		if isOption(newPos) {
			return newPos, true
		}
	}

	return common.Vector{}, false
}
