package main

import (
	_ "embed"
	"fmt"
	"github.com/zwolsman/go-aoc/common"
	"golang.org/x/exp/maps"
	"math"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

type tile struct {
	id    int
	body  []string
	edges map[location]string
}

func newTile(id int, body []string) (tiles []tile) {
	for _, e := range edges(body) {
		tiles = append(tiles, tile{
			id:    id,
			body:  body,
			edges: e,
		})
	}

	return
}

func part1(in []byte) any {
	rawTiles := strings.Split(string(in), "\n\n")
	var tiles []tile

	for _, rawTile := range rawTiles {
		id, body := parseTile(rawTile)
		tiles = append(tiles, newTile(id, body)...)
	}

	width := int(math.Sqrt(float64(len(tiles) / 12)))
	backtrack(tiles, make(map[common.Vector]tile), common.Vector{}, width)

	//fmt.Println(results)
	//fmt.Println(tiles)
	fmt.Println(common.MaxBy(results, func(x map[common.Vector]tile) int {
		return len(x)
	}))
	return nil
}

var (
	right  = common.Vector{X: 1}
	left   = common.Vector{X: -1}
	top    = common.Vector{Y: -1}
	bottom = common.Vector{Y: 1}
)

var results []map[common.Vector]tile

func backtrack(options []tile, solution map[common.Vector]tile, position common.Vector, max int) {
	if len(options) == 0 {
		results = append(results, solution)
		return
	}

	if position.X == max {
		position = position.Plus(common.Vector{X: -max, Y: 1})
	}
	fmt.Println(len(options))
	fmt.Println(position)
	fmt.Println()

	//if position.Y == max && position.X == max {
	//	return solution
	//}

	knownEdges := make(map[location]string)
	if tile, ok := solution[position.Plus(left)]; ok {
		knownEdges[LEFT] = tile.edges[RIGHT]
	}
	if tile, ok := solution[position.Plus(right)]; ok {
		knownEdges[RIGHT] = tile.edges[LEFT]
	}
	if tile, ok := solution[position.Plus(bottom)]; ok {
		knownEdges[BOTTOM] = tile.edges[TOP]
	}
	if tile, ok := solution[position.Plus(top)]; ok {
		knownEdges[TOP] = tile.edges[BOTTOM]
	}

	possibilities := filter(options, func(t tile) bool {
		for l, e := range knownEdges {
			if v := t.edges[l]; v != e {
				return false
			}
		}

		return true
	})

	//fmt.Println(knownEdges)
	//fmt.Println(possibilities)

	if len(possibilities) == 0 {
		//fmt.Println("KWEEENIE BRO")
	}

	for _, option := range possibilities {
		nextOptions := filter(options, func(o tile) bool {
			return o.id != option.id
		})

		nextSolution := make(map[common.Vector]tile)

		maps.Copy(nextSolution, solution)
		nextSolution[position] = option

		backtrack(nextOptions, nextSolution, position.Plus(right), max)
	}
	results = append(results, solution)
}

type location = int

const (
	TOP location = iota
	RIGHT
	BOTTOM
	LEFT
)

func part2(in []byte) any {
	return nil
}

func parseTile(rawTile string) (int, []string) {
	lines := strings.Split(rawTile, "\n")
	var id int
	_, err := fmt.Fscanf(strings.NewReader(lines[0]), "Tile %d:", &id)
	if err != nil {
		panic(err)
	}

	return id, lines[1:]
}

func edges(body []string) []map[location]string {
	topLeftRight := body[0]
	bottomLeftRight := body[len(body)-1]

	var leftTopBottom, rightTopBottom string

	for i := 0; i < len(body); i++ {
		leftTopBottom += string(body[i][0])
		rightTopBottom += string(body[i][len(body[i])-1])
	}

	original := map[location]string{
		TOP:    topLeftRight,
		RIGHT:  rightTopBottom,
		BOTTOM: bottomLeftRight,
		LEFT:   leftTopBottom,
	}

	mutations := []map[location]string{
		original,
	}

	for i := 1; i < 4; i++ { // 3 ways to turn
		mutation := make(map[location]string)
		for l, e := range original {
			mutation[(l+i)%4] = e
		}
		mutations = append(mutations, mutation)
	}

	for _, mutation := range mutations {
		for i := 0; i < 2; i++ {
			flip := make(map[location]string)
			maps.Copy(flip, mutation)

			flip[i] = mutation[i+2]
			flip[i+2] = mutation[i]
			mutations = append(mutations, flip)
		}

	}

	return mutations
}

func filter[S any](s []S, f func(s S) bool) []S {
	var out []S
	for _, e := range s {
		if f(e) {
			out = append(out, e)
		}
	}

	return out
}
