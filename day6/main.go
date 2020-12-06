package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/mzwolsman/Developer/go-aoc/day6/input.txt")
	groups := Group(string(data))
	part1(groups)
	part2(groups)
}

func part1(groups [][]string) {
	sum := CountVotes(groups)
	println(sum, "is the number of questions to which anyone answered \"yes\"")
}

func part2(groups [][]string) {
	sum := CountAllYesVotes(groups)
	println(sum, "is the sum of the number of questions to which everyone answered \"yes\"")
}

func CountVotes(groups [][]string) (sum int) {
	votes := createVoteMap(groups)
	for _, groupVotes := range votes {
		sum += len(groupVotes)
	}
	return
}

func CountAllYesVotes(groups [][]string) (sum int) {
	votes := createVoteMap(groups)

	for i, group := range groups {
		groupVotes := votes[i]
		for _, v := range groupVotes {
			if v == len(group) {
				sum++
			}
		}
	}

	return
}

func Group(input string) (output [][]string) {
	groups := strings.Split(input, "\n\n")
	for _, group := range groups {
		people := strings.Split(group, "\n")
		output = append(output, people)
	}
	return
}

func createVoteMap(groups [][]string) (allVotes []map[int32]int) {
	for _, group := range groups {
		votes := make(map[int32]int)
		for _, person := range group {
			for _, vote := range person {
				votes[vote]++
			}
		}
		allVotes = append(allVotes, votes)
	}
	return
}
