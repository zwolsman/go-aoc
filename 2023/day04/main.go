package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

var cardRegex = regexp.MustCompile("(\\d+|\\|)")

func part1(in []byte) any {
	cards := strings.Split(string(in), "\n")

	var sum int
	for _, card := range cards {
		numbers := cardRegex.FindAllString(card, -1)
		sep := slices.Index(numbers, "|")

		winningNumbers, cardNumbers := associate(numbers[1:sep]), associate(numbers[sep+1:])

		score := 0
		for c := range cardNumbers {
			if _, ok := winningNumbers[c]; ok {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		sum += score
	}
	return sum
}

func part2(in []byte) any {
	rawCards := strings.Split(string(in), "\n")

	cards := make([][]int, len(rawCards))

	// links cards
	for i, card := range rawCards {
		numbers := cardRegex.FindAllString(card, -1)
		sep := slices.Index(numbers, "|")

		winningNumbers, cardNumbers := associate(numbers[1:sep]), associate(numbers[sep+1:])

		n := 1
		for c := range cardNumbers {
			if _, ok := winningNumbers[c]; ok {
				cards[i] = append(cards[i], i+n)
				n++
			}
		}
	}

	// count cards
	counts := make(map[int]int) //card -> count
	for card, winners := range cards {
		counts[card]++
		count := 1
		if m, ok := counts[card]; ok {
			count = m
		}

		for _, w := range winners {
			counts[w] += count
		}
	}

	var s int
	for _, c := range counts {
		s += c
	}
	return s
}

func associate(in []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, s := range in {
		m[s] = struct{}{}
	}
	return m
}
