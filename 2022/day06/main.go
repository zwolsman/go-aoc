package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) any {
	return run(in, 4)
}

func part2(in []byte) any {
	return run(in, 14)
}

func run[T comparable](in []T, window int) int {
	left, right := 0, 0
	occurrences := make(map[T]int)

	for right < len(in) {
		if len(occurrences) == window {
			found := true
			for _, v := range occurrences {
				if v > 1 {
					found = false
					break
				}
			}
			if found {
				return right
			}
		}

		occurrences[in[right]]++

		if right-left == window {
			occurrences[in[left]]--
			if occurrences[in[left]] == 0 {
				delete(occurrences, in[left])
			}
			left++
		}
		right++
	}

	return -1
}
