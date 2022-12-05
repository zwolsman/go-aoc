package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

var matcher = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func part1(in []byte) any {
	return run(in, func(count int, from, to *stack[rune]) {
		for i := 0; i < count; i++ {
			item := from.pop()
			to.push(item)
		}
	})
}

func part2(in []byte) any {
	return run(in, func(count int, from, to *stack[rune]) {
		items := from.content[:count]
		for i := 0; i < count; i++ {
			from.pop()
		}

		for i := 0; i < count; i++ {
			to.push(items[len(items)-i-1])
		}
	})
}

type stack[T any] struct {
	content []T
}

func (s *stack[T]) pop() T {
	item := s.content[0]
	s.content = s.content[1:]
	return item
}

func (s *stack[T]) peek() T {
	return s.content[0]
}

func (s *stack[T]) push(item T) {
	var c []T
	c = append(c, item)
	c = append(c, s.content...)
	s.content = c
}

func (s *stack[T]) append(item T) {
	var c []T
	c = append(c, s.content...)
	c = append(c, item)
	s.content = c
}

func (s *stack[T]) string(mapFunc func(T) string) string {
	str := "["
	for _, c := range s.content {
		str += mapFunc(c) + " "
	}
	str += "]"
	return str
}

func run(in []byte, move func(count int, from, to *stack[rune])) string {
	stacks := make(map[int]*stack[rune])
	isBuildingStacks := true
	var keys []int
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			isBuildingStacks = false
		}

		if isBuildingStacks {

			for i, c := range line {
				stackId := i/4 + 1
				if unicode.IsLetter(c) {
					if _, ok := stacks[stackId]; !ok {
						stacks[stackId] = &stack[rune]{}
						keys = append(keys, stackId)
					}

					stacks[stackId].append(c)
				}
			}
		}

		if m := matcher.FindStringSubmatch(line); len(m) == 4 {
			count, err := strconv.Atoi(m[1])
			if err != nil {
				panic(err)
			}

			from, err := strconv.Atoi(m[2])
			if err != nil {
				panic(err)
			}

			to, err := strconv.Atoi(m[3])
			if err != nil {
				panic(err)
			}

			move(count, stacks[from], stacks[to])
		}
	}

	sort.Ints(keys)
	var output string
	for _, id := range keys {
		stack := stacks[id]
		output += string(stack.peek())
	}

	return output
}
