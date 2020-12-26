package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	cups := readCups()

	cups = playCups(cups, 100)
	for _, cup := range cups {
		print(cup+1, " ")
	}
	println()
}

func part2() {
	initial := readCups()
	cups := make([]int, 1_000_000)

	copy(cups, initial)

	for i := len(initial); i < len(cups); i++ {
		cups[i] = i
	}
	cups = playCups(cups, 10_000_000)

	c1 := cups[0]
	c2 := cups[c1]
	println((c1 + 1) * (c2 + 1))
}

func playCups(cups []int, rounds int) []int {

	N := len(cups)
	next := make([]int, N)
	for i := 0; i < len(next); i++ {
		next[cups[i]] = cups[mod(i+1, N)]
	}

	current := cups[0]
	for i := 0; i < rounds; i++ {
		c1 := next[current]
		c2 := next[c1]
		c3 := next[c2]

		next[current] = next[c3]

		dest := mod(current-1, N)
		for c1 == dest || c2 == dest || c3 == dest {
			dest = mod(dest-1, N)
		}

		next[c3] = next[dest]
		next[dest] = c1

		current = next[current]
	}

	return next
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func readCups() []int {
	data, err := ioutil.ReadFile("./2020/day23/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var cups []int
	for _, str := range strings.Split(string(data), "") {
		cup, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		cups = append(cups, cup-1)
	}
	return cups
}
