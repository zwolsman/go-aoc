package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./2022/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var sum, max int
	var elves []int

	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		if line == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			sum += n
			if sum > max {
				max = sum
			}
		}
	}
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1])

	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
