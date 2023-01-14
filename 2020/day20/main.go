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

func part1(in []byte) any {
	tiles := readAllTiles(in)
	classifications := classifyTiles(tiles)

	result := 1
	for id, classification := range classifications {
		if classification == corner {
			result *= id
		}
	}
	return result
}

func part2(in []byte) any {
	tiles := readAllTiles(in)

	width := int(math.Sqrt(float64(len(tiles) / 12)))
	solution := backtrack(tiles, make(map[common.Vector]tile), common.Vector{}, width)

	corners := []common.Vector{
		{0, 0},
		{width - 1, 0},
		{0, width - 1},
		{width - 1, width - 1},
	}

	sum := 1
	for _, c := range corners {
		tile, ok := solution[c]
		if !ok {
			panic("didn't find corner")
		}

		sum *= tile.id
	}

	return sum

	return nil
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

var (
	right  = common.Vector{X: 1}
	left   = common.Vector{X: -1}
	top    = common.Vector{Y: -1}
	bottom = common.Vector{Y: 1}
)

func backtrack(options []tile, solution map[common.Vector]tile, position common.Vector, max int) map[common.Vector]tile {
	if len(options) == 0 {
		return solution
	}

	if position.X == max {
		position = position.Plus(common.Vector{X: -max, Y: 1})
	}

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

	for _, option := range possibilities {
		nextOptions := filter(options, func(o tile) bool {
			return o.id != option.id
		})

		nextSolution := make(map[common.Vector]tile)

		maps.Copy(nextSolution, solution)
		nextSolution[position] = option

		if result := backtrack(nextOptions, nextSolution, position.Plus(right), max); result != nil {
			return result
		}
	}

	return nil
}

type location = int
type classification = int

const (
	TOP location = iota
	RIGHT
	BOTTOM
	LEFT
)

const (
	corner classification = iota
	edge
)

func classifyTiles(tiles []tile) map[int]classification {
	classifications := make(map[int]classification)
	edges := make(map[string]map[int]struct{})

	for _, tile := range tiles {
		for _, e := range tile.edges {
			if _, ok := edges[e]; !ok {
				edges[e] = make(map[int]struct{})
			}

			edges[e][tile.id] = struct{}{}
		}
	}

	counts := make(map[int]int)
	highest := 0
	for _, ids := range edges {
		if len(ids) == 1 {
			for id := range ids {
				counts[id]++
				if highest < counts[id] {
					highest = counts[id]
				}
			}
		}
	}

	for id, count := range counts {
		if count == highest {
			classifications[id] = corner
		} else {
			classifications[id] = edge
		}
	}

	return classifications
}

func readAllTiles(in []byte) []tile {
	rawTiles := strings.Split(string(in), "\n\n")
	var tiles []tile

	for _, rawTile := range rawTiles {
		id, body := parseTile(rawTile)
		tiles = append(tiles, newTile(id, body)...)
	}

	return tiles
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
	bodyToEdges := func() map[location]string {
		topLeftRight := body[0]
		bottomLeftRight := body[len(body)-1]

		var leftTopBottom, rightTopBottom string

		for i := 0; i < len(body); i++ {
			leftTopBottom += string(body[i][0])
			rightTopBottom += string(body[i][len(body[i])-1])
		}

		return map[location]string{
			TOP:    topLeftRight,
			RIGHT:  rightTopBottom,
			BOTTOM: bottomLeftRight,
			LEFT:   leftTopBottom,
		}
	}

	mutations := []map[location]string{
		bodyToEdges(), // Original
	}

	for i := 0; i < 3; i++ { // 3 ways to turn
		body = rotate(body)
		mutations = append(mutations, bodyToEdges())
	}

	for _, mutation := range mutations {
		for i := 0; i < 2; i++ {
			flip := make(map[location]string)
			maps.Copy(flip, mutation)

			flip[i] = mutation[i+2]
			flip[i+2] = mutation[i]

			flip[i+1] = reverse(flip[i+1])
			flip[(i+3)%4] = reverse(flip[(i+3)%4])

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

func rotate(b []string) []string {
	var body [][]string
	for _, x := range b {
		body = append(body, strings.Split(x, ""))
	}

	// reverse the matrix
	for i, j := 0, len(body)-1; i < j; i, j = i+1, j-1 {
		body[i], body[j] = body[j], body[i]
	}

	// transpose it
	for i := 0; i < len(body); i++ {
		for j := 0; j < i; j++ {
			body[i][j], body[j][i] = body[j][i], body[i][j]
		}
	}

	var result []string

	for _, x := range body {
		result = append(result, strings.Join(x, ""))
	}

	return result
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
