package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./2020/day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var adapters []int

	for scanner.Scan() {
		line := scanner.Text()
		adapter, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)

	part1(adapters)
	part2(adapters)
}

func part2(adapters []int) {
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	history := make(map[int]int)
	var combinations func(a int) int

	combinations = func(a int) (count int) {
		if h, ok := history[a]; ok {
			return h
		}

		for _, adapter := range adapters {
			diff := adapter - a
			if diff > 0 && diff <= 3 {
				count += combinations(adapter)
			}
		}

		if adapters[len(adapters)-1] == a { // done
			count++
		}

		history[a] = count
		return
	}

	println("combinations", combinations(0))
}

func part1(adapters []int) {
	oneDiff, threeDiff := 0, 1
	for i, adapter := range adapters {
		var diff int
		if i == 0 {
			diff = adapter
		} else {
			diff = adapter - adapters[i-1]
		}
		switch diff {
		case 1:
			oneDiff++
			break
		case 3:
			threeDiff++
			break
		default:
			println("diff is not 1 or 3", diff)
		}
	}

	println(oneDiff, threeDiff, oneDiff*threeDiff)
}
