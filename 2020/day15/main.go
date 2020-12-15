package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2020/day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var numbers []int
	for _, str := range strings.Split(string(data), ",") {
		num, err := strconv.Atoi(str)

		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	play(numbers, 2020)
	play(numbers, 30000000)
}

func play(init []int, limit int) {

	history := make(map[int][]int)
	prev := make([]int, 2)
	for turn := 0; turn < limit; turn++ {
		if turn < len(init) {
			next := init[turn]
			history[next] = append(history[next], turn)
			prev[0], prev[1] = next, prev[0]
			continue
		}

		next := 0
		turns, spoken := history[prev[0]]
		if !spoken || len(turns) == 1 {
			next = 0
		} else if prev[0] == prev[1] {
			next = 1
		} else {
			a, b := turns[len(turns)-1], turns[len(turns)-2]
			next = a - b
		}
		history[next] = append(history[next], turn)
		prev[0], prev[1] = next, prev[0]
	}

	println(prev[0])
}
