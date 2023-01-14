package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) any {
	monkeys := make(map[string]func() int)

	for _, line := range strings.Split(string(in), "\n") {
		id := line[:4]
		n, err := strconv.Atoi(line[6:])

		if err == nil {
			monkeys[id] = func() int {
				return n
			}
		} else {
			fields := strings.Fields(line[6:])

			l, r := fields[0], fields[2]

			var fn func() int

			switch fields[1] {
			case "+":
				fn = func() int {
					return monkeys[l]() + monkeys[r]()
				}
				break
			case "-":
				fn = func() int {
					return monkeys[l]() - monkeys[r]()
				}
				break
			case "/":
				fn = func() int {
					return monkeys[l]() / monkeys[r]()
				}
				break
			case "*":
				fn = func() int {
					return monkeys[l]() * monkeys[r]()
				}
				break
			}
			monkeys[id] = fn
		}
	}

	return monkeys["root"]()
}

func part2(in []byte) any {
	return nil
}
