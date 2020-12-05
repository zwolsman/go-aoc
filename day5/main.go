package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	boardingPasses := readAllLines("/Users/mzwolsman/Developer/go-aoc/day5/input.txt")

	highestSeatID := 0
	for _, boardingPass := range boardingPasses {
		id := SeatID(boardingPass)
		if id > highestSeatID {
			highestSeatID = id
		}
	}

	fmt.Println("highest seat id", highestSeatID)
}

func SeatID(boardingPass string) int {

	rowStr := boardingPass[:7]
	rowStr = strings.ReplaceAll(rowStr, "F", "0")
	rowStr = strings.ReplaceAll(rowStr, "B", "1")

	row, err := strconv.ParseInt(rowStr, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	columnStr := boardingPass[7:]
	columnStr = strings.ReplaceAll(columnStr, "L", "0")
	columnStr = strings.ReplaceAll(columnStr, "R", "1")

	column, err := strconv.ParseInt(columnStr, 2, 32)
	if err != nil {
		log.Fatal(err)
	}

	return int((row * 8) + column)
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
