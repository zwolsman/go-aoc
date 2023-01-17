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

var directions = []common.Vector{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func part1(in []byte) int {
	cubes := make(map[common.Vector]int)
	for _, line := range strings.Split(string(in), "\n") {
		cube, err := common.NewVector(line)
		if err != nil {
			panic(err)
		}
		exposed := 6

		for _, dir := range directions {
			other := cube.Plus(dir)
			if n, ok := cubes[other]; ok {
				exposed--
				cubes[other] = n - 1
			}
		}
		cubes[cube] = exposed
	}

	var sum int
	for _, exposure := range cubes {
		sum += exposure
	}

	return sum
}

func part2(in []byte) any {
	return nil
}
