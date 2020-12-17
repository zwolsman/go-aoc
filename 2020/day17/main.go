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
	board := readBoard()

	board.Print()
	fmt.Println()
	for cycle := 0; cycle < 6; cycle++ {
		fmt.Sprintf("cycle %d\n", cycle+1)
		board = simulate(board)
		board.Print()
		fmt.Println()
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
				board[Vector{x, y, z}] = true
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
				neighbors = append(neighbors, Vector{x, y, z})
			}
		}
	}
	return
}
func expand(board Board) Board {
	newBoard := make(Board)

	for coord, v := range board {
		newCoord := Vector{coord.x + 1, coord.y + 1, coord.z}
		newBoard[newCoord] = v
	}
	return newBoard
}

func simulate(board Board) Board {
	options := neighbors3D()
	newState := make(Board)
	//board = expand(board)

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
	x, y, z int
}

func (v *Vector) String() string {
	return fmt.Sprintf("Vector(%d, %d, %d)", v.x, v.y, v.z)
}

func (v *Vector) Plus(other *Vector) Vector {
	return Vector{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
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
				coord := Vector{x, y, z}
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
