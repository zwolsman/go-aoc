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
	part1(start, stop)
}

func part1(start, stop int) {
	correctPasswords := 0

	for i := start; i <= stop; i++ {
		str := fmt.Sprintf("%d", i)
		if isValidPassword(str) {
			correctPasswords++
		}
	}

	println(correctPasswords)
}

func isValidPassword(password string) bool {
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
