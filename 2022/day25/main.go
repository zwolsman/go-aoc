package main

import (
	_ "embed"
	"fmt"
	"github.com/zwolsman/go-aoc/common"
	"math"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
}

func part1(in []byte) string {
	var sum int
	for _, n := range common.Map(strings.Split(string(in), "\n"), toDec) {
		sum += n
	}

	return toSNAFU(sum)
}

var symbols = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

func toDec(SNAFU string) int {
	var sum int
	for i, c := range SNAFU {
		n := symbols[c]

		sum += n * multiplier(len(SNAFU)-1-i)
	}

	return sum
}

func multiplier(pos int) int {
	return int(math.Pow(5, float64(pos)))
}

func best(index, target int) rune {
	best := math.MaxInt
	var r rune
	for s, v := range symbols {
		n := multiplier(index) * v
		diff := target - n
		if diff < 0 {
			diff *= -1
		}
		if diff <= best {
			best = diff
			r = s
		}
	}

	return r
}

func toSNAFU(n int) string {
	var index int
	for ; multiplier(index)*2 < n; index++ {

	}

	var result string
	for ; index >= 0; index-- {
		s := best(index, n)
		n -= symbols[s] * multiplier(index)
		result += string(s)
	}

	return strings.TrimLeft(result, "0")
}
