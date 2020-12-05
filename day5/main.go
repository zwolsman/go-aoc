package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	boardingPasses := readAllLines("/Users/mzwolsman/Developer/go-aoc/day5/input.txt")

	highestSeatID := 0
	for _, boardingPass := range boardingPasses {
		id := SeatID(boardingPass)
		fmt.Println(boardingPass, id)
		if id > highestSeatID {
			highestSeatID = id
		}
	}

	fmt.Println("highest seat id", highestSeatID)
}

func SeatID(boardingPass string) int {
	row := calculate(127, 'F', 'B', boardingPass[0:7])
	column := calculate(7, 'L', 'R', boardingPass[7:])

	println("row", row, "column", column)

	return (row * 8) + column
}

func calculate(high int, lowerHalf, upperHalf rune, chars string) int {
	low := 0
	for _, code := range chars {
		diff := float64(high - low)
		switch code {
		case lowerHalf:
			high = high - int(math.Ceil(diff/2))
			break
		case upperHalf:
			low = low + int(math.Ceil(diff/2))
			break
		}
	}

	last := chars[len(chars)-1:]
	if last == string(lowerHalf) {
		return low
	}
	if last == string(upperHalf) {
		return high
	}

	println(chars, low, high)
	log.Fatal("error, could not solve ", chars)
	return 0
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
