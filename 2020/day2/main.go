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
	file, err := os.Open("./2020/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var validPaswwords int
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.SplitAfter(line, ": ")
		policy := createPolicy(split[0])
		passwd := split[1]
		if err != nil {
			log.Fatal(err)
		}
		if policy(passwd) {
			validPaswwords++
		}
	}

	fmt.Println(validPaswwords)

}

func createPolicy(policy string) func(passwd string) bool {
	initialSplit := strings.Split(policy, " ")
	letter := initialSplit[1][0]
	numbers := strings.Split(initialSplit[0], "-")
	lower, _ := strconv.Atoi(numbers[0])
	upper, _ := strconv.Atoi(numbers[1])

	return func(passwd string) bool {
		a := passwd[lower-1]
		b := passwd[upper-1]
		return (a == letter && b != letter) || (a != letter && b == letter)
	}
}
