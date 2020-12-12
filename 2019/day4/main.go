package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2019/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(data), "-")
	start, stop := parseRange(input)
	p1, p2 := 0, 0

	for i := start; i <= stop; i++ {
		str := fmt.Sprintf("%d", i)
		if part1(str) {
			p1++
		}
		if part2(str) {
			p2++
		}
	}

	println("part 1", p1)
	println("part 2", p2)
}

func part1(password string) bool {
	if len(password) != 6 {
		return false
	}

	hasDouble := false
	for j := 0; j < len(password)-1; j++ {
		if password[j] > password[j+1] {
			return false
		}
		if !hasDouble && password[j] == password[j+1] {
			hasDouble = true
		}
	}
	return hasDouble
}

func part2(password string) bool {
	if len(password) != 6 {
		return false
	}

	groups := make(map[int32]int)
	for i := 0; i < len(password)-1; i++ {
		if password[i] > password[i+1] {
			return false
		}
	}

	for _, c := range password {
		groups[c]++
	}

	for _, v := range groups {
		if v == 2 {
			return true
		}
	}
	return false
}

func parseRange(input []string) (int, int) {
	start, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal(err)
	}
	stop, err := strconv.Atoi(input[1])
	if err != nil {
		log.Fatal(err)
	}
	return start, stop
}
