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
	classifications := classifyTiles(tiles)
	width := int(math.Sqrt(float64(len(tiles) / 12)))

	var startingTiles []tile
	for id, classification := range classifications {
		if classification == corner {
			startingTiles = append(startingTiles, filter(tiles, func(t tile) bool {
				return t.id == id
			})...)
		}
	}

	var solution map[common.Vector]tile

	for _, t := range startingTiles {
		start := map[common.Vector]tile{
			common.Vector{}: t,
		}

		btiles := filter(tiles, func(t2 tile) bool {
			return t2.id != t.id
		})

		if solution = backtrack(btiles, start, common.Vector{X: 1}, width); solution != nil {
			break
		}
	}

	var solvedMap []string
	for y := 0; y < width; y++ {

		tile := solution[common.Vector{Y: y}]

		current := strip(tile.body)
		for x := 1; x < width; x++ {
			coord := common.Vector{X: x, Y: y}
			tile := solution[coord]

			for i, l := range strip(tile.body) {
				current[i] += l
			}
		}
		solvedMap = append(solvedMap, current...)
	}

	water := 0
	for _, l := range solvedMap {
		for _, c := range l {
			if c == '#' {
				water++
			}
		}
	}

	pattern := parseSeaMonsterPattern()
	for _, mutation := range mutations(0, solvedMap) {
		if count := countPattern(mutation.body, pattern); count > 0 {
			return water - (len(pattern) * count)
		}
	}

	return -1
}

func countPattern(body []string, pattern []common.Vector) int {
	m := make(map[common.Vector]rune)
	for y, row := range body {
		for x, c := range row {
			m[common.Vector{X: x, Y: y}] = c
		}
	}

	count := 0

	for y := 0; y < len(body); y++ {
		for x := 0; x < len(body); x++ {
			coord := common.Vector{X: x, Y: y}
			transposedPattern := common.Apply(common.Vector.Plus, pattern, coord)

			result := filter(transposedPattern, func(p common.Vector) bool {
				return m[p] == '#'
			})

			if len(result) == len(pattern) {
				count++
			}
		}
	}

	return count
}

type tile struct {
	id    int
	body  []string
	edges map[location]string
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
	inner classification = iota
	corner
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
		tiles = append(tiles, mutations(id, body)...)
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

func edges(body []string) map[location]string {
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

func mutations(id int, body []string) []tile {
	var results []tile

	addMutation := func() {
		results = append(results, tile{
			id:    id,
			body:  body,
			edges: edges(body),
		})
	}

	addMutation()

	for i := 0; i < 3; i++ { // 3 ways to turn
		body = rotate(body)
		addMutation()
	}

	for _, mutation := range results {
		body = flipVertical(mutation.body)
		addMutation()
		body = flipHorizontal(mutation.body)
		addMutation()
	}

	return results
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

func flipHorizontal(b []string) []string {
	result := make([]string, len(b))
	for i, s := range b {
		result[i] = reverse(s)
	}
	return result
}

func flipVertical(b []string) []string {
	result := make([]string, len(b))
	for i, s := range b {
		result[len(b)-i-1] = s
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

const monster = `
                  # 
#    ##    ##    ###
 #  #  #  #  #  #   
`

func parseSeaMonsterPattern() []common.Vector {
	rows := strings.Split(monster, "\n")
	var coords []common.Vector

	for y, row := range rows[1:] {
		for x, c := range row {
			if c == '#' {
				coords = append(coords, common.Vector{X: x, Y: y})
			}
		}
	}

	return coords
}

func strip(body []string) []string {
	var output []string

	for i := 1; i < len(body)-1; i++ {
		output = append(output, body[i][1:len(body[i])-1])
	}
	return output
}
