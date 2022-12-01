package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./2022/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(file))
	fmt.Println(part2(file))
}

func part1(in []byte) int {
	var sum, max int
	lines := strings.Split(string(in), "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		n, _ := strconv.Atoi(line)
		if line == "" {
			sum = 0
		} else {
			sum += n
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

func part2(in []byte) int {
	var elves []int
	var sum int

	lines := strings.Split(string(in), "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		n, _ := strconv.Atoi(line)
		if line == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			sum += n
		}
	}
	elves = append(elves, sum)

	sort.Ints(elves)
	fmt.Println(elves)
	return elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
}
