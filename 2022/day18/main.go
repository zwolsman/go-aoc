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
	cubes, err := readCubes(in)
	if err != nil {
		panic(err)
	}

	exposures := make(map[common.Vector]int)
	for _, cube := range cubes {
		exposed := 6

		for _, dir := range directions {
			other := cube.Plus(dir)
			if n, ok := exposures[other]; ok {
				exposed--
				exposures[other] = n - 1
			}
		}
		exposures[cube] = exposed
	}

	var sum int
	for _, exposure := range exposures {
		sum += exposure
	}

	return sum
}

func part2(in []byte) any {
	return nil
}

func readCubes(in []byte) ([]common.Vector, error) {
	var cubes []common.Vector
	for _, line := range strings.Split(string(in), "\n") {
		cube, err := common.NewVector(line)
		if err != nil {
			return nil, err
		}
		cubes = append(cubes, cube)
	}
	return cubes, nil
}
