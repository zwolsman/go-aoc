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

	fmt.Printf("%v, %v\n", minutes, buses)

	println("*** part 1 ***")
	part1(minutes, buses)

}

func part1(minutes int, buses []int) {

	longestWait, id := math.MaxInt64, -1
	for _, bus := range buses {
		next := int(math.Ceil(float64(minutes)/float64(bus))) * bus
		wait := next - minutes
		if wait < longestWait {
			longestWait = wait
			id = wait * bus
		}
		fmt.Printf("bus %d, next dept: %d, wait: %d, answer: %d\n", bus, next, wait, (next-minutes)*bus)
	}
	fmt.Printf("The ID of the earliest bus you can take to the airport multiplied by the number of minutes you'll need to wait for that bus is %d\n", id)
}
