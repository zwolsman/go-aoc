package main

import (
	intprogram "../program"
	"fmt"
	"log"
	"sync"
)

func main() {
	println("*** part 1 ***")
	part1()
	println("*** part 2 ***")
	part2()
}

func part1() {
	phases := []int{0, 1, 2, 3, 4}
	highest := 0
	for _, sequence := range permutations(phases) {
		fmt.Printf("%v\n", sequence)
		if output := runSequence(sequence); output > highest {
			highest = output
		}
	}
	println("highest output", highest)
}

func part2() {
	phases := []int{9, 8, 7, 6, 5}
	highest := 0
	for _, sequence := range permutations(phases) {
		fmt.Printf("%v\n", sequence)
		if output := runSequenceLoop(sequence); output > highest {
			highest = output
		}
	}
	println("highest output", highest)
}

func runSequence(sequence []int) (output int) {
	if len(sequence) != 5 {
		log.Fatal("sequence is not 5 numbers")
	}
	for _, init := range sequence {
		amp := intprogram.Read("./2019/day7/input.txt")
		channel := make(chan int, 2)
		amp.In = channel
		amp.Out = channel

		amp.In <- init
		amp.In <- output
		amp.Run()
		output = <-amp.Out
	}
	return
}

func runSequenceLoop(sequence []int) int {
	amps := make([]intprogram.Program, len(sequence))
	for i := 0; i < len(amps); i++ {
		amp := intprogram.Read("./2019/day7/input.txt")
		amp.SetPrefix(fmt.Sprintf("amp-%d ", i))
		if i != 0 {
			amp.In = amps[i-1].Out
		}
		amp.Out = make(chan int, 2)
		amps[i] = amp
	}
	amps[0].In = amps[len(amps)-1].Out

	//Set amp 0 with ZERO once
	amps[0].In <- 0
	for i, init := range sequence {
		log.Println("setting up amp", i, "with code", init)
		amps[i].In <- init
	}

	wg := sync.WaitGroup{}
	for i, amp := range amps {
		wg.Add(1)
		go func(i int, amp intprogram.Program) {
			log.Println("starting amp", i)
			amp.Run()
			wg.Done()
		}(i, amp)
	}
	wg.Wait()
	return <-amps[len(amps)-1].Out
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	var res [][]int

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
