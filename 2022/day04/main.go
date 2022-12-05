package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

var matcher = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func part1(in []byte) int {
	return run(in, func(l1, l2, r1, r2 int) bool {
		return (l1 >= r1 && l1 <= r2 && l2 >= r1 && l2 <= r2) || (r1 >= l1 && r1 <= l2 && r2 >= l1 && r2 <= l2)
	})
}

func part2(in []byte) any {
	return run(in, func(l1, l2, r1, r2 int) bool {
		return (l1 >= r1 && l1 <= r2) || (l2 >= r1 && l2 <= r2) || (r1 >= l1 && r1 <= l2) || (r2 >= l1 && r2 <= l2)
	})
}

func run(in []byte, overlapFn func(int, int, int, int) bool) int {
	var sum int
	for _, line := range strings.Split(string(in), "\n") {
		n := matcher.FindStringSubmatch(line)
		l1, err := strconv.Atoi(n[1])
		if err != nil {
			panic(err)
		}
		l2, err := strconv.Atoi(n[2])
		if err != nil {
			panic(err)
		}
		r1, err := strconv.Atoi(n[3])
		if err != nil {
			panic(err)
		}
		r2, err := strconv.Atoi(n[4])
		if err != nil {
			panic(err)
		}
		if overlapFn(l1, l2, r1, r2) {
			sum++
		}
	}
	return sum
}
