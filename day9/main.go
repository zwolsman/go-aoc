package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	data := readInput()
	preamble := 25

	isValid := func(prev []int64, current int64) bool {
		for i, x := range prev {
			for j, y := range prev {
				if i == j {
					continue
				}

				if x+y == current {
					return true
				}
			}
		}
		return false
	}
	for i := preamble; i < len(data); i++ {
		prevNumbers := data[i-preamble : i]
		current := data[i]

		if !isValid(prevNumbers, current) {
			println(current)
		}
	}
}

func readInput() []int64 {
	var nums []int64
	file, err := os.Open("/Users/mzwolsman/Developer/go-aoc/day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	return nums
}
