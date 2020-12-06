package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	part1()
}

func part1() {
	lines := readAllLines("/Users/mzwolsman/Developer/go-aoc/day6/input.txt")
	groups := Group(lines)
	sum := CountVotes(groups)
	println(sum, "is the number of questions to which anyone answered \"yes\"")
}

func CountVotes(groups []string) (sum int) {
	for _, group := range groups {
		votes := make(map[int32]bool)
		for _, vote := range group {
			votes[vote] = true
		}
		sum += len(votes)
	}
	return
}

func Group(lines []string) (groups []string) {
	group := ""
	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = ""
			continue
		}
		group += line
	}
	groups = append(groups, group)
	return
}

func readAllLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
