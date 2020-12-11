package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/mzwolsman/Developer/go-aoc/day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var layout [][]string

	for scanner.Scan() {
		row := scanner.Text()
		layout = append(layout, strings.Split(row, ""))
	}

	println("original state")
	println(stringify(layout))
	part1(layout)
}

func stringify(layout [][]string) (out string) {
	for _, row := range layout {
		out += strings.Join(row, "") + "\n"
	}
	return out
}

func part1(layout [][]string) {
	i := 0

	adjacent := func(x, y int) string {
		var output []string
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				newx := x + i
				newy := y + j

				if newx < 0 || newx >= len(layout[0]) {
					continue
				}
				if newy < 0 || newy >= len(layout) {
					continue
				}

				output = append(output, layout[newy][newx])
			}
		}
		return strings.Join(output, "")
	}

	for true {
		println("simulation", i)
		newLayout := simulate(layout, 4, adjacent)
		i++

		if str := stringify(newLayout); str == stringify(layout) {
			println("stabalized!")
			println("seats", strings.Count(str, "#"))
			break
		}
		layout = newLayout
	}
}

func simulate(in [][]string, n int, adjacent func(int, int) string) [][]string {
	out := make([][]string, len(in))

	for y, row := range in {
		newRow := make([]string, len(row))
		for x, val := range row {
			next := val
			if val == "#" {
				if strings.Count(adjacent(x, y), "#") >= n {
					next = "L"
				}
			}

			if val == "L" {
				if !strings.Contains(adjacent(x, y), "#") {
					next = "#"
				}
			}

			newRow[x] = next
		}
		out[y] = newRow
	}
	return out
}
