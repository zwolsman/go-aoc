package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2020/day24/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(data), "\n")

	tiles := part1(instructions)
	part2(tiles)
}

func part1(instructions []string) Tiles {
	tiles := make(Tiles)

	for _, instruction := range instructions {
		tile := followInstruction(instruction)
		if _, ok := tiles[tile]; ok {
			delete(tiles, tile)
		} else {
			tiles[tile] = struct{}{}
		}
	}

	fmt.Printf("flipped %d tiles to black\n", len(tiles))
	return tiles
}

func part2(tiles Tiles) {
	for day := 1; day <= 100; day++ {
		tiles = simulate(tiles)
	}

	println("How many tiles will be black after 100 days?", len(tiles))
}

func simulate(in Tiles) Tiles {
	tiles := make(Tiles)
	out := make(Tiles)

	for t := range in {
		for _, dir := range directions {
			tiles[t.add(dir)] = struct{}{}
		}
	}

	for t := range tiles {
		blackTiles := 0
		for _, dir := range directions {
			if _, ok := in[t.add(dir)]; ok {
				blackTiles++
			}
		}

		if _, isBlack := in[t]; isBlack && (blackTiles == 1 || blackTiles == 2) || !isBlack && blackTiles == 2 {
			out[t] = struct{}{}
		}
	}

	return out
}

type Tiles map[Vector]struct{}

var regex = regexp.MustCompile("((se|sw|nw|ne)|([ew]))")

// directions
var (
	//e, se, sw, w, nw and ne
	east = Vector{
		1, 0, -1,
	}
	west = Vector{
		-1, 0, 1,
	}
	northEast = Vector{
		0, 1, -1,
	}
	northWest = Vector{
		-1, 1, 0,
	}
	southEast = Vector{
		1, -1, 0,
	}
	southWest = Vector{
		0, -1, 1,
	}

	directions = map[string]Vector{
		"e":  east,
		"w":  west,
		"ne": northEast,
		"nw": northWest,
		"se": southEast,
		"sw": southWest,
	}
)

func followInstruction(instruction string) Vector {
	pos := Vector{}
	results := regex.FindAllStringSubmatch(instruction, -1)
	for _, m := range results {
		direction, ok := directions[m[0]]
		if !ok {
			log.Fatal(m[0])
		}
		pos = pos.add(direction)
	}
	return pos
}

type Vector struct {
	x, y, z int
}

func (v *Vector) add(v2 Vector) Vector {
	return Vector{
		v.x + v2.x,
		v.y + v2.y,
		v.z + v2.z,
	}
}

func (v *Vector) String() string {
	return fmt.Sprintf("<x=%d, y=%dm z=%d>", v.x, v.y, v.z)
}
