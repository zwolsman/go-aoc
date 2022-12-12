package main

import (
	_ "embed"
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/zwolsman/go-aoc/common"
	"math"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) int {
	m := common.ReadMap(in, 0)
	w := m.Width()
	origin, found := m.Find('S')
	if !found {
		panic("couldn't find starting position")
	}

	target, found := m.Find('E')
	if !found {
		panic("couldn't find starting position")
	}

	// Set correct elevation
	m[origin] = 'a'
	m[target] = 'z'

	graph := createGraph(m)
	path, err := graph.Shortest(id(w, origin), id(w, target))

	if err != nil {
		panic(err)
	}

	return int(path.Distance)
}

func part2(in []byte) int {
	m := common.ReadMap(in, 0)
	w := m.Width()
	origin, found := m.Find('S')
	if !found {
		panic("couldn't find starting position")
	}

	target, found := m.Find('E')
	if !found {
		panic("couldn't find starting position")
	}

	// Set correct elevation
	m[origin] = 'a'
	m[target] = 'z'

	graph := createGraph(m)
	shortest := math.MaxInt
	for _, position := range m.FindAll('a') {
		path, err := graph.Shortest(id(w, position), id(w, target))
		if err != nil {
			continue
		}

		if int(path.Distance) < shortest {
			shortest = int(path.Distance)
		}
	}

	return shortest
}

func createGraph(m common.Map2D) *dijkstra.Graph {
	w := m.Width()
	graph := dijkstra.NewGraph()
	for p := range m {
		graph.AddVertex(id(w, p))
	}

	for position, height := range m {
		for _, d := range common.LRUD {
			newPos := position.Plus(d)
			v, ok := m[newPos]
			if !ok { // out of bounds
				continue
			}

			if height == v-1 || height >= v {
				a, b := id(w, position), id(w, newPos)

				err := graph.AddArc(a, b, 1)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	return graph
}

func id(width int, v common.Vector) int {
	return width*v.X + v.Y
}
