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

	part1(instructions)
}

func part1(instructions []string) {
	tiles := make(map[Vector]struct{})

	for _, instruction := range instructions {
		tile := followInstruction(instruction)
		if _, ok := tiles[tile]; ok {
			delete(tiles, tile)
		} else {
			tiles[tile] = struct{}{}
		}
	}

	fmt.Printf("flipped %d tiles to black", len(tiles))
}

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
