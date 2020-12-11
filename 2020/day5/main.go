package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	boardingPasses := readAllLines("./2020/day5/input.txt")
	part1(boardingPasses)
	part2(boardingPasses)
}

func part1(boardingPasses []string) {
	highestSeatID := 0
	for _, boardingPass := range boardingPasses {
		id := SeatID(boardingPass)
		if id > highestSeatID {
			highestSeatID = id
		}
	}

	println("highest seat id", highestSeatID)
}

func part2(boardingPasses []string) {
	seatIDs := make([]int, len(boardingPasses))
	for i, boardingPass := range boardingPasses {
		seatIDs[i] = SeatID(boardingPass)
	}

	sort.Ints(seatIDs)
	for i := 0; i < len(seatIDs)-1; i++ {
		diff := seatIDs[i] - seatIDs[i+1]
		if diff != -1 {
			missingSeatID := seatIDs[i] + 1
			println("my seat id", missingSeatID)
		}
	}
}

var rowReplacer = strings.NewReplacer("F", "0", "B", "1")
var columnReplacer = strings.NewReplacer("L", "0", "R", "1")

func SeatID(boardingPass string) int {
	row, err := strconv.ParseInt(rowReplacer.Replace(boardingPass[:7]), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	column, err := strconv.ParseInt(columnReplacer.Replace(boardingPass[7:]), 2, 32)
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
