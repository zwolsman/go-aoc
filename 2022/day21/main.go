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
	monkeys := readMonkeys(in)
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
		n, ok := evaluate(id, monkeys)
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
				facts[other] = substitute(facts[target], n, i, operator)
			}

			solve(other, monkeys, facts)
			break
		}
	}
}

func substitute(x, constant, pos int, operator string) int {
	switch operator {
	case "/":
		if pos == 0 {
			return constant / x
		} else {
			return constant * x
		}
	case "*":
		return x / constant
	case "+":
		return x - constant
	case "-":
		if pos == 0 {
			return (x - constant) * -1
		} else {
			return x + constant
		}
	default:
		panic("unknown operator")
	}
}

func evaluate(id string, monkeys map[string]string) (int, bool) {
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

		leftN, ok := evaluate(l, monkeys)
		if !ok {
			return 0, false
		}
		rightN, ok := evaluate(r, monkeys)
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

func readMonkeys(in []byte) map[string]func() int {
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

	return monkeys
}
