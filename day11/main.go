package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const HEIGHT = 98
const WIDTH = 90

//const HEIGHT = 10
//const WIDTH = 10

func main() {
	file, err := os.Open("/Users/mzwolsman/Developer/go-aoc/day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var layout []string

	for scanner.Scan() {
		row := scanner.Text()
		layout = append(layout, strings.Split(row, "")...)
	}

	prettyPrint(layout)
	seats := strings.Count(strings.Join(layout, ""), "#")
	i := 0
	for true {
		println("simulation", i+1)
		layout = simulate(layout)
		newSeats := strings.Count(strings.Join(layout, ""), "#")
		println("seats occupied", newSeats)

		//prettyPrint(layout)
		if newSeats == seats {
			break
		}
		seats = newSeats
		i++
	}
}

func prettyPrint(layout []string) {
	for i := 0; i < HEIGHT; i++ {
		slice := layout[i*WIDTH : i*WIDTH+WIDTH]
		println(strings.Join(slice, ""))
	}
	println()
}

func simulate(in []string) []string {
	out := make([]string, len(in))

	adjacent := func(pos int) []string {
		baseX, baseY := pos%HEIGHT, pos/HEIGHT
		var output []string
		//println(pos, baseX, baseY)
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				x, y := baseX+i, baseY+j

				if x < 0 || y < 0 {
					//println("TOO DAMN LOW", x, y)
					continue
				}
				if x >= HEIGHT || y >= WIDTH {
					//println("TOO DAMN HIGH", x, y)
					continue
				}
				newpos := y*HEIGHT + x
				if newpos == 8820 {
					println(pos, newpos, x, y)
				}
				output = append(output, in[newpos])
			}
		}
		return output
	}

	for pos, val := range in {
		next := val
		switch val {
		case ".":
			break
		case "#":
			adj := strings.Join(adjacent(pos), "")
			if strings.Count(adj, "#") >= 4 {
				next = "L"
			}
			break
		case "L":
			adj := strings.Join(adjacent(pos), "")
			if !strings.Contains(adj, "#") {
				next = "#"
			}
			break
		default:
			log.Fatal("can't handle", val)
		}

		out[pos] = next
	}
	return out
}
