package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Slope struct {
	right, down int
}

func main() {
	treeMap := readTreeMap()

	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1

	for _, slope := range slopes {
		result *= treesEncountered(slope, treeMap)
	}

	fmt.Println("Encoutered", result, "trees")
}

func treesEncountered(slope Slope, treeMap [][]bool) (treesEncountered int) {
	x := 0
	for y := 0; y < len(treeMap); y += slope.down {
		width := len(treeMap[y])
		if treeMap[y][x] {
			treesEncountered++
		}

		x = (x + slope.right) % width
	}
	return
}

func readTreeMap() (treeMap [][]bool) {
	file, err := os.Open("./2020/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var treeRow []bool

		for _, c := range line {
			treeRow = append(treeRow, c == '#')
		}
		treeMap = append(treeMap, treeRow)
	}

	return
}
