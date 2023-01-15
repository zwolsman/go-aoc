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
	monkeys, _, _ := readMonkeys(in)
	return monkeys["root"]()
}

func part2(in []byte) any {
	monkeys := make(map[string]string)
	for _, line := range strings.Split(string(in), "\n") {
		id := line[:4]
		monkeys[id] = line[6:]
	}

	delete(monkeys, "humn")

	facts := make(map[string]int)
	for id := range monkeys {
		n, ok := calculate(id, monkeys)
		if ok {
			facts[id] = n
		}
	}

	solve("root", monkeys, facts)

	return facts["humn"]
}

func solve(target string, monkeys map[string]string, facts map[string]int) {
	if target == "humn" {
		return
	}

	fields := strings.Fields(monkeys[target])

	l, operator, r := fields[0], fields[1], fields[2]
	options := []string{l, r}
	for i, option := range options {
		if n, ok := facts[option]; ok {
			other := options[(i+1)%2]

			if target == "root" {
				facts[other] = n
			} else {
				switch operator {
				case "/":
					if i == 0 {
						facts[other] = n / facts[target]
					} else {
						facts[other] = n * facts[target]
					}
				case "*":
					facts[other] = facts[target] / n
				case "+":
					facts[other] = facts[target] - n
				case "-":
					if i == 0 {
						facts[other] = (facts[target] - n) * -1
					} else {
						facts[other] = facts[target] + n
					}
				}
			}

			solve(other, monkeys, facts)
			break
		}
	}
}

func calculate(id string, monkeys map[string]string) (int, bool) {
	monkey, ok := monkeys[id]
	if !ok {
		return 0, false
	}

	fields := strings.Fields(monkey)

	if len(fields) == 1 {
		n, _ := strconv.Atoi(fields[0])
		return n, true
	} else {
		l, o, r := fields[0], fields[1], fields[2]

		leftN, ok := calculate(l, monkeys)
		if !ok {
			return 0, false
		}
		rightN, ok := calculate(r, monkeys)
		if !ok {
			return 0, false
		}

		switch o {
		case "*":
			return leftN * rightN, true
		case "/":
			return leftN / rightN, true
		case "+":
			return leftN + rightN, true
		case "-":
			return leftN - rightN, true
		default:
			return 0, false
		}
	}
}

func readMonkeys(in []byte) (map[string]func() int, string, string) {
	monkeys := make(map[string]func() int)
	var rootLeft, rootRight string
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
			if id == "root" {
				rootLeft, rootRight = l, r
			}

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

	return monkeys, rootLeft, rootRight
}
