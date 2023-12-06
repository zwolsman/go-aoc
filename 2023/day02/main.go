package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

var cubesRegex = regexp.MustCompile("(\\d+) (\\w+)")

const cubesCountIndex = 1
const cubesColorIndex = 2

func part1(in []byte) any {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var sum int

	for _, line := range strings.Split(string(in), "\n") {
		var gameId int
		gameStr, cubesStr, _ := strings.Cut(line, ":")
		_, err := fmt.Fscanf(strings.NewReader(gameStr), "Game %d", &gameId)
		if err != nil {
			panic(err)
		}

		isPossible := true

		for _, round := range strings.Split(cubesStr, ";") {
			cubes := cubesRegex.FindAllStringSubmatch(round, -1)

			for _, cube := range cubes {
				count, err := strconv.Atoi(cube[cubesCountIndex])
				if err != nil {
					panic(err)
				}
				color := cube[cubesColorIndex]

				limit := limits[color]
				if count > limit {
					isPossible = false
					break
				}
			}
		}
		if isPossible {
			sum += gameId
		}
	}

	return sum
}

func part2(in []byte) any {
	return nil
}
