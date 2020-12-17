package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	active   = '#'
	inactive = '.'
)

func main() {
	part1()
	part2()
}

func part1() {
	options := neighbors3D()
	board := readBoard()

	board.Print()
	fmt.Println()
	for cycle := 0; cycle < 6; cycle++ {
		fmt.Sprintf("cycle %d\n", cycle+1)
		board = simulate(board, options)
		board.Print()
		fmt.Println()
	}
	fmt.Println(len(board))
}

func part2() {
	options := neighbors4D()
	board := readBoard()

	board.Print()
	fmt.Println()
	for cycle := 0; cycle < 6; cycle++ {
		fmt.Sprintf("cycle %d\n", cycle+1)
		board = simulate(board, options)
	}
	fmt.Println(len(board))
}

func readBoard() Board {
	board := make(Board)
	data, err := ioutil.ReadFile("./2020/day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	z := 0
	for y, row := range strings.Split(string(data), "\n") {
		for x, state := range row {
			if state == active {
				board[Vector{x, y, z, 0}] = true
			}
		}
	}
	return board
}

func neighbors3D() (neighbors []Vector) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				neighbors = append(neighbors, Vector{x, y, z, 0})
			}
		}
	}
	return
}

func neighbors4D() (neighbors []Vector) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					neighbors = append(neighbors, Vector{x, y, z, w})
				}
			}
		}
	}
	return
}

func simulate(board Board, options []Vector) Board {
	newState := make(Board)

	coordsToCheck := make(Board)

	for coord, _ := range board {
		for _, toCheck := range getNeighbors(coord, options) {
			coordsToCheck[toCheck] = true
		}
	}

	for coord, _ := range coordsToCheck {
		activeNeighbors := 0
		for _, n := range getNeighbors(coord, options) {
			if board[n] {
				activeNeighbors++
			}
		}

		v := board[coord]
		if v {
			if activeNeighbors == 2 || activeNeighbors == 3 {
				newState[coord] = true
			}
		} else {
			if activeNeighbors == 3 {
				newState[coord] = true
			}
		}
	}

	return newState
}

func getNeighbors(base Vector, options []Vector) (neighbors []Vector) {
	for _, option := range options {
		neighbors = append(neighbors, base.Plus(&option))
	}
	return
}

type Vector struct {
	x, y, z, w int
}

func (v *Vector) String() string {
	return fmt.Sprintf("Vector(%d, %d, %d, %d)", v.x, v.y, v.z, v.w)
}

func (v *Vector) Plus(other *Vector) Vector {
	return Vector{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
		w: v.w + other.w,
	}
}

type Board map[Vector]bool

func (b Board) Dimensions() (x, y, z int) {
	for coord, _ := range b {
		if coord.x > x {
			x = coord.x
		}
		if coord.y > y {
			y = coord.y
		}
		if coord.z > z {
			z = coord.z
		}
	}
	return
}

func (b Board) Print() {
	maxX, maxY, maxZ := b.Dimensions()
	for z := -maxZ; z <= maxZ; z++ {
		fmt.Printf("z=%d\n", z)
		for y := -maxY; y <= maxY; y++ {
			for x := -maxX; x <= maxX; x++ {
				coord := Vector{x, y, z, 0}
				if b[coord] {
					fmt.Print(string(active))
				} else {
					fmt.Print(string(inactive))
				}
			}
			fmt.Println()
		}
	}
}
