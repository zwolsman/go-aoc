package main

import (
	_ "embed"
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/zwolsman/go-aoc/common"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var in []byte

type pipe struct {
	id       string
	flowRate int
	valves   []string
	open     bool
}

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

var pipeRegex = regexp.MustCompile(`^Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)

func part1(in []byte) int {
	pipes := readPipes(in)
	referenceGraph := createGraph(pipes)
	return open(referenceGraph, "AA", 30, 0, pipes, 0)
}

func part2(in []byte) int {
	pipes := readPipes(in)
	referenceGraph := createGraph(pipes)

	var ids []string
	for id, pipe := range pipes {
		if pipe.flowRate > 0 {
			ids = append(ids, id)
		}
	}

	//sort.Strings(ids)
	scores := make(map[string]int)

	c := common.Combinations(ids)

	for _, combination := range c {
		fmt.Printf("Point 4 before sort combination: %v\n", ids)
		//sort.Strings(combination)

		fmt.Println(combination)
		fmt.Printf("Point 4: %v\n", ids)

		key := strings.Join(combination, ".")
		test := common.Copy(pipes)

		for id := range test {
			if _, ok := common.Index(combination, id); !ok {
				delete(test, id)
			}
		}

		score := open(referenceGraph, "AA", 26, 0, test, 0)
		scores[key] = score
	}

	var max int

	for _, combination := range c {
		var other []string

		for _, id := range ids {
			if _, ok := common.Index(combination, id); !ok {
				other = append(other, id)
			}
		}

		sort.Strings(other)

		leftKey := strings.Join(combination, ".")
		rightKey := strings.Join(other, ".")
		x, ok := scores[leftKey]
		if !ok {
			log.Fatal("couldn't find left key")
		}
		y, ok := scores[rightKey]
		if !ok {
			log.Fatal("couldn't find right key")
		}
		max = common.Max(max, x+y)
	}

	return max
}

func createGraph(pipes map[string]pipe) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	for id := range pipes {
		graph.AddMappedVertex(id)
	}

	for id, p := range pipes {
		for _, dst := range p.valves {
			err := graph.AddMappedArc(id, dst, 1)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return graph
}

var cache = make(map[cacheKey]int)

type cacheKey struct {
	location string
	time     int
	options  string
}

func open(graph *dijkstra.Graph, location string, limit int, time int, pipes map[string]pipe, score int) int {
	if time > limit {
		return score
	}

	key := cacheKey{
		location: location,
		time:     time,
	}

	for _, p := range pipes {
		if !p.open {
			key.options += p.id + "-"
		}
	}

	if v, ok := cache[key]; ok {
		return score + v
	}

	p := pipes[location]

	var tto int
	if p.flowRate > 0 {
		tto = 1
	}

	pipeScore := (30 - time - tto) * p.flowRate

	pipesLeft := make(map[string]pipe)
	for id, pipe := range pipes {
		if !pipe.open && pipe.flowRate > 0 {
			pipesLeft[id] = pipe
		}
	}

	if len(pipesLeft) == 0 {
		return score + pipeScore
	}

	src, err := graph.GetMapping(location)
	if err != nil {
		log.Fatal(err)
	}

	var max int
	for id := range pipesLeft {
		dst, err := graph.GetMapping(id)
		if err != nil {
			log.Fatal(err)
		}

		path, err := graph.Shortest(src, dst)
		if err != nil {
			log.Fatal(err)
		}

		next := common.Copy(pipes)

		p := next[id]
		p.open = true
		next[id] = p

		max = common.Max(max, open(graph, id, limit, time+tto+int(path.Distance), next, score+pipeScore))
	}

	cache[key] = max - score
	return max
}

func readPipes(in []byte) map[string]pipe {
	pipes := make(map[string]pipe)

	for _, line := range strings.Split(string(in), "\n") {
		match := pipeRegex.FindStringSubmatch(line)

		if len(match) == 0 {
			log.Fatal("didn't find match for input", line)
		}

		id, flowRate := match[1], match[2]
		valves := strings.Split(match[3], ", ")

		n, err := strconv.Atoi(flowRate)
		if err != nil {
			panic(err)
		}

		pipes[id] = pipe{
			id:       id,
			flowRate: n,
			valves:   valves,
			open:     false,
		}
	}

	return pipes
}
