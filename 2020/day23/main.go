package main

import (
	"container/ring"
	"fmt"
	"golang.org/x/tools/container/intsets"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	part1()
}

func part1() {
	cups, min, max := readCups()

	cups = playCups(cups, 100, min, max)
	printCups(cups)
}

func playCups(cups *ring.Ring, rounds, min, max int) *ring.Ring {
	for i := 0; i < 100; i++ {
		cup := cups.Value.(int)
		//fmt.Printf("-- move %d --\n", i+1)
		//print("all cups: ")
		//printCupsWithHighlight(cups, cup)
		pickUp := cups.Unlink(3)

		//print("pickup: ")
		//printCups(pickUp)
		target := cup - 1
		for {
			if target < min {
				target = max
			}
			if !containsCup(target, pickUp) {
				break
			}
			//println("target", target, "is already in the pickups. substracting one")
			target--

		}

		//println("target cup -> ", target)
		for {
			if cups.Value.(int) == target {
				break
			}
			cups = cups.Next()
		}

		cups.Link(pickUp)
		for {
			if cups.Value.(int) == cup {
				break
			}
			cups = cups.Next()
		}
		cups = cups.Next()
	}

	//println("-- final --")
	return cups
}

func containsCup(cup int, cups *ring.Ring) bool {
	if cups.Value.(int) == cup {
		return true
	}
	for p := cups.Next(); p != cups; p = p.Next() {
		if p.Value.(int) == cup {
			return true
		}
	}
	return false
}

func printCups(cups *ring.Ring) {
	print(cups.Value.(int), " ")
	for p := cups.Next(); p != cups; p = p.Next() {
		print(p.Value.(int), " ")
	}
	println()
}

func printCupsWithHighlight(cups *ring.Ring, highlight int) {
	printFn := func(cup int) {
		format := "%d "
		if cup == highlight {
			format = "(%d) "
		}

		fmt.Printf(format, cup)
	}
	printFn(cups.Value.(int))
	for p := cups.Next(); p != cups; p = p.Next() {
		printFn(p.Value.(int))
	}
	println()
}

func readCups() (*ring.Ring, int, int) {
	data, err := ioutil.ReadFile("./2020/day23/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := ring.New(len(data))
	min, max := intsets.MaxInt, 0
	for _, str := range strings.Split(string(data), "") {
		cup, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		if cup > max {
			max = cup
		}
		if cup < min {
			min = cup
		}
		r.Value = cup
		r = r.Next()
	}
	return r, min, max
}
