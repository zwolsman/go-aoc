package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2020/day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	println("*** part 1 ***")
	part1(lines)
	println()

	println("*** part 2***")
	part2(lines)
}

func part1(lines []string) {
	minutes, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	var buses []int
	for _, b := range strings.Split(lines[1], ",") {
		b, err := strconv.Atoi(b)
		if err != nil {
			continue
		}
		buses = append(buses, b)
	}

	//fmt.Printf("%v, %v\n", minutes, buses)

	longestWait, id := math.MaxInt64, -1
	for _, bus := range buses {
		next := int(math.Ceil(float64(minutes)/float64(bus))) * bus
		wait := next - minutes
		if wait < longestWait {
			longestWait = wait
			id = wait * bus
		}
		//fmt.Printf("bus %d, next dept: %d, wait: %d, answer: %d\n", bus, next, wait, (next-minutes)*bus)
	}
	fmt.Printf("The ID of the earliest bus you can take to the airport multiplied by the number of minutes you'll need to wait for that bus is %d\n", id)
}

func part2(lines []string) {
	var buses []int

	for _, b := range strings.Split(lines[1], ",") {
		var n int
		if b == "x" {
			n = 1
		} else {
			n, _ = strconv.Atoi(b)
		}

		buses = append(buses, n)
	}

	interval, departure := 1, 1

	for i, bus := range buses {
		slice := buses[:i+1]
		for {
			if validate(slice, departure) {
				interval *= bus
				break
			} else {
				departure += interval
			}
		}
	}

	println("The earliest timestamp such that all of the listed bus IDs depart at offsets matching their positions in the list is", departure)
}

func validate(buses []int, departure int) bool {
	for i, bus := range buses {
		if !isValidDeparture(i, bus, departure) {
			return false
		}
	}
	return true
}

func isValidDeparture(index, busId, departure int) bool {
	return (departure+index)%busId == 0
}
