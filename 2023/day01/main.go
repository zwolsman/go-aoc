package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) int {
	var sum int
	for _, str := range strings.Split(string(in), "\n") {
		var digits []rune
		for _, c := range str {
			if unicode.IsDigit(c) {
				digits = append(digits, c)
			}
		}
		sum += int(digits[0]-'0') * 10
		sum += int(digits[len(digits)-1] - '0')
	}

	return sum
}

func part2(in []byte) any {
	lookup := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	options := make([]string, 0, len(lookup)*2)
	for k, v := range lookup {
		options = append(options, k)
		options = append(options, string(v))
	}

	var sum int
	for _, str := range strings.Split(string(in), "\n") {
		var matches []string
		for i := 0; i < len(str); i++ {
			for _, option := range options {
				if strings.HasPrefix(str[i:], option) {
					matches = append(matches, option)
				}
			}
		}

		for i, digit := range matches {
			if n, ok := lookup[digit]; ok {
				matches[i] = string(n)
			}
		}
		fmt.Println(str, matches)
		sum += int(matches[0][0]-'0') * 10
		sum += int(matches[len(matches)-1][0] - '0')
	}

	return sum
}
