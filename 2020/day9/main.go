package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	data := readInput()
	part1(data)
	part2(data)
}

func part1(data NumRange) {
	preamble := 25

	isValid := func(prev NumRange, current int64) bool {
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

type NumRange []int64

func (a NumRange) Len() int           { return len(a) }
func (a NumRange) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NumRange) Less(i, j int) bool { return a[i] < a[j] }

func part2(data NumRange) {
	for i := 2; i < len(data); i++ {
		println("testing window", i)
		if testWindow(data, i) {
			break
		}
	}
}
func sum(d NumRange) (sum int64) {
	for _, num := range d {
		sum += num
	}
	return sum
}

func testWindow(data NumRange, window int) bool {
	const target = 20874512

	for i := window; i < len(data); i++ {
		slidingWindow := data[i-window : i]
		if sum(slidingWindow) == target {
			println("found set")
			sort.Sort(slidingWindow)
			a := slidingWindow[0]
			b := slidingWindow[window-1] // NOT 1931501367179
			println("encryption weakness", a+b)
			return true
		}
	}

	return false
}

func readInput() NumRange {
	var nums NumRange
	file, err := os.Open("./2020/day9/input.txt")
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
