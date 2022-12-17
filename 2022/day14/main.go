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

func part1(in []byte) any {

	sandSpawn := common.Vector{X: 500}
	sand := sandSpawn.Copy()

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

	placedSand := make(map[common.Vector]int)

	for sand.Y <= highestY+1 {
		next, found := findNextOption(sand, rocks, placedSand)
		if !found {
			placedSand[sand]++
			sand = sandSpawn.Copy()
			continue
		}

		sand = next
	}

	return len(placedSand)
}

func part2(in []byte) any {
	return nil
}

var options = [3]common.Vector{
	{0, 1},
	{-1, 1},
	{1, 1},
}

func findNextOption(origin common.Vector, paths []common.Path, occupied map[common.Vector]int) (common.Vector, bool) {
	for _, option := range options {
		newPos := origin.Plus(option)
		if _, ok := occupied[newPos]; ok {
			//fmt.Printf("can't place sand at %v, already occupied\n", newPos)
			continue
		}

		isOption := true
		for _, path := range paths {
			if path.Intersects(newPos) {
				//fmt.Printf("Path intersects with new pos(%v)\n", newPos)
				isOption = false
				break
			}
		}

		if isOption {
			return newPos, true
		}
	}

	return common.Vector{}, false
}
