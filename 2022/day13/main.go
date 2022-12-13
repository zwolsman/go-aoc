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
	var packets []packet

	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}

		packet, _ := parsePacket(line[1:])
		packets = append(packets, packet)
	}

	var validIndices int
	for i := 0; i < len(packets); i += 2 {
		n := i/2 + 1
		a, b := packets[i], packets[i+1]

		valid := verify(a, b)
		if valid == CORRECT {
			validIndices += n
		}
	}
	return validIndices

}

func part2(in []byte) any {
	return nil
}

const (
	UNKNOWN = iota
	CORRECT
	INCORRECT
)

func verify(a, b packet) int {
	orderState := UNKNOWN
	for i := 0; i < len(a); i++ {
		if i >= len(b) {
			break
		}
		left, right := a[i], b[i]
		ln, leftIsInt := left.(int)
		rn, rightIsInt := right.(int)

		if leftIsInt && rightIsInt { // l & r == int
			switch {
			case ln == rn: // if same, continue
				continue
			case ln < rn: // found correct order
				return CORRECT
			case ln > rn: // found incorrect order
				return INCORRECT
			}
		}

		larr, leftIsList := left.(packet)
		rarr, rightIsList := right.(packet)

		if leftIsList && !rightIsList {
			rarr = packet{rn}
		}
		if !leftIsList && rightIsList {
			larr = packet{ln}
		}
		if result := verify(larr, rarr); result != UNKNOWN {
			return result
		}
	}

	if len(a) < len(b) { // left ran out before right
		return CORRECT
	}
	if len(a) > len(b) { // right ran out before left
		return INCORRECT
	}

	return orderState
}

type packet = []any

func parsePacket(line string) (packet, int) {
	var raw string
	var result packet
	for i := 0; i < len(line); i++ {
		if line[i] == '[' {
			list, count := parsePacket(line[i+1:])
			result = append(result, list)
			i += count
			continue
		}
		if line[i] == ']' {
			n, err := strconv.Atoi(raw)
			if err == nil {
				result = append(result, n)
			}
			return result, i + 1
		}
		if line[i] == ',' {
			n, err := strconv.Atoi(raw)
			if err == nil {
				result = append(result, n)
				raw = ""
			}
			continue
		}

		raw += string(line[i])
	}

	return result, len(line)
}
