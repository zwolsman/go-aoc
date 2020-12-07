package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	bags := readBags()

	part1(bags)
	part2(bags)
}

func part1(bags map[string][]*Bag) {
	result := find("shiny gold", bags)
	println(len(result))
}

func part2(bags map[string][]*Bag) {
	var bagsInBag func(string) int

	bagsInBag = func(bag string) (sum int) {
		for _, child := range bags[bag] {
			sum += child.num + (child.num * bagsInBag(child.color))
		}
		return
	}

	answer := bagsInBag("shiny gold")

	println(answer)
}

func readBags() map[string][]*Bag {
	file, err := os.Open("/Users/mzwolsman/Developer/go-aoc/day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	regex, err := regexp.Compile("(no other bags|(?:\\s(?P<num>\\d+)\\s)?(?P<color>\\w+\\s\\w+)\\sbags?)")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	bagMap := make(map[string][]*Bag)

	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindAllStringSubmatch(line, -1)

		base := match[0][3]

		var children []*Bag

		for i := 1; i < len(match); i++ {
			info := match[i]
			if info[0] == "no other bags" {
				continue
			}

			num, _ := strconv.Atoi(info[2])
			color := info[3]

			children = append(children, &Bag{num, color})
		}
		bagMap[base] = children
	}

	return bagMap
}

func unique(input []string) (keys []string) {
	temp := make(map[string]int)
	for _, l := range input {
		temp[l] = 1
	}
	for k, _ := range temp {
		keys = append(keys, k)
	}
	return
}

func find(color string, bagMap map[string][]*Bag) (result []string) {
	for k, v := range bagMap {
		for _, bag := range v {
			if bag.color == color {
				result = append(result, k)
				result = append(result, find(k, bagMap)...)
			}
		}
	}

	return unique(result)
}

type Bag struct {
	num   int
	color string
}
