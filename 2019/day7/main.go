package main

import (
	intprogram "../program"
	"fmt"
)

func main() {
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

func runSequence(sequence []int) (output int) {
	for _, init := range sequence {
		io := make(chan int, 2)
		io <- init
		io <- output
		amp := intprogram.Read("./2019/day7/input.txt")
		amp.SetChannel(io)
		amp.Run()
		output = <-io
	}
	return
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

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
