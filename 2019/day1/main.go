package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./2019/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var modules []int

	for scanner.Scan() {
		module, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}
		modules = append(modules, module)
	}

	part1(modules)
}

func part1(modules []int) {
	sum := 0
	for _, module := range modules {
		fuel := math.Floor(float64(module)/3.0) - 2
		sum += int(fuel)
	}

	println("The sum of the fuel requirements for all of the modules on the spacecraft", sum)
}
