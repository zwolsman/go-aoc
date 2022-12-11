package main

import (
	_ "embed"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(part1(createMonkeys()))
	fmt.Println(part2(createMonkeys()))
}

func part1(monkeys []*monkey) int {
	return run(monkeys, 20, func(stress int) int {
		return stress / 3
	})
}

func part2(monkeys []*monkey) int {
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.test
	}

	return run(monkeys, 10_000, func(stress int) int {
		return stress % lcm
	})
}

func run(monkeys []*monkey, rounds int, reduce func(int) int) int {
	inspections := make([]int, len(monkeys))

	for i := 0; i < rounds; i++ {
		for id, m := range monkeys {
			for _, item := range m.items {
				inspections[id]++
				item = reduce(m.operation(item))

				var receiver *monkey
				if item%m.test == 0 {
					receiver = m.chain[0]
				} else {
					receiver = m.chain[1]
				}

				receiver.items = append(receiver.items, item)
			}

			m.items = []item{}
		}
	}

	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func createMonkeys() []*monkey {
	var (
		m0 = monkey{
			items: []int{98, 89, 52},
			operation: func(i item) item {
				return i * 2
			},
			test: 5,
		}
		m1 = monkey{
			items: []int{57, 95, 80, 92, 57, 78},
			operation: func(i item) item {
				return i * 13
			},
			test: 2,
		}
		m2 = monkey{
			items: []int{82, 74, 97, 75, 51, 92, 83},
			operation: func(i item) item {
				return i + 5
			},
			test: 19,
		}
		m3 = monkey{
			items: []int{97, 88, 51, 68, 76},
			operation: func(i item) item {
				return i + 6
			},
			test: 7,
		}
		m4 = monkey{
			items: []int{63},
			operation: func(i item) item {
				return i + 1
			},
			test: 17,
		}
		m5 = monkey{
			items: []int{94, 91, 51, 63},
			operation: func(i item) item {
				return i + 4
			},
			test: 13,
		}
		m6 = monkey{
			items: []int{61, 54, 94, 71, 74, 68, 98, 83},
			operation: func(i item) item {
				return i + 2
			},
			test: 3,
		}
		m7 = monkey{
			items: []int{90, 56},
			operation: func(i item) item {
				return i * i
			},
			test: 11,
		}
	)

	m0.chain = [2]*monkey{
		&m6,
		&m1,
	}
	m1.chain = [2]*monkey{
		&m2,
		&m6,
	}
	m2.chain = [2]*monkey{
		&m7,
		&m5,
	}
	m3.chain = [2]*monkey{
		&m0,
		&m4,
	}
	m4.chain = [2]*monkey{
		&m0,
		&m1,
	}
	m5.chain = [2]*monkey{
		&m4,
		&m3,
	}
	m6.chain = [2]*monkey{
		&m2,
		&m7,
	}
	m7.chain = [2]*monkey{
		&m3,
		&m5,
	}

	return []*monkey{
		&m0,
		&m1,
		&m2,
		&m3,
		&m4,
		&m5,
		&m6,
		&m7,
	}
}

type item = int
type monkey struct {
	items     []item
	operation func(item) item
	test      int
	chain     [2]*monkey
}
