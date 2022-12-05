package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

type rucksack = map[uint8]int

func part1(in []byte) int {
	var sum int
	for _, line := range strings.Split(string(in), "\n") {
		h := len(line) / 2
		l, r := line[:h], line[h:]
		lm, rm := make(rucksack), make(rucksack)

		for i := 0; i < h; i++ {
			kl, kr := l[i], r[i]
			lm[kl] = 1
			rm[kr] = 1
		}

		for kl, _ := range lm {
			if _, ok := rm[kl]; ok {
				sum += int(priority(kl))
			}
		}
	}

	return sum
}

func part2(in []byte) int {
	elves := []rucksack{
		make(rucksack),
		make(rucksack),
		make(rucksack),
	}

	var sum int
	lines := strings.Split(string(in), "\n")
	for i := 0; i < len(lines); i++ {
		for _, v := range lines[i] {
			elves[i%3][uint8(v)] = 1
		}

		if i%3 == 2 { // validate all 3
			temp := make(rucksack)
			for _, elf := range elves {
				for item, _ := range elf {
					temp[item] += 1
					delete(elf, item)
				}
			}

			for k, v := range temp {
				if v == 3 {
					sum += int(priority(k))
				}
			}

		}
	}

	return sum
}

func priority(p uint8) uint8 {
	var offset uint8
	if p >= 'A' && p <= 'Z' {
		offset = 'A' - 26
	} else {
		offset = 'a'
	}
	return p - offset + 1
}
