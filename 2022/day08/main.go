package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) int {
	m := createMap(in)
	var rows [][]int
	var cols [][]int

	for _, r := range m {
		rows = append(rows, mask(r))
	}

	for i := 0; i < len(m); i++ {
		var out []int
		for j := 0; j < len(m); j++ {
			out = append(out, m[j][i])
		}
		cols = append(cols, mask(out))
	}

	for i := 0; i < len(rows)*len(cols); i++ {
		x := i % len(rows)
		y := i / len(cols)

		if rows[y][x] > 0 || cols[x][y] > 0 {
			m[y][x] = 1
		} else {
			m[y][x] = 0
		}
	}

	var sum int

	for _, arr := range m {
		for _, v := range arr {
			sum += v
		}
	}
	return sum
}

func mask(arr []int) []int {
	out := make([]int, len(arr))

	left, right := -1, -1
	for i := 0; i < len(arr); i++ {

		if n := arr[i]; n > left {
			out[i] += 1
			left = n
		}

		j := len(arr) - i - 1
		if n := arr[j]; n > right {
			out[j] += 1
			right = n
		}
	}

	return out
}

func part2(in []byte) any {
	return nil
}

func createMap(in []byte) [][]int {
	lines := strings.Split(string(in), "\n")
	grid := make([][]int, len(lines))

	for i, row := range lines {
		r := make([]int, len(row))
		for j, h := range row {
			r[j] = int(h - '0')
		}
		grid[i] = r
	}
	return grid
}
