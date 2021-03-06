package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

const (
	maskRegex = "^mask = ([X01]+)$"
	memRegex  = "^mem\\[(\\d+)\\] = (\\d+)$"
	bitSize   = 36
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("./2020/day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	maskRegexp, memRegexp := setupRegexp()

	mem := make(map[int]int)
	var mask []bit

	for scanner.Scan() {
		line := scanner.Text()
		if maskRegexp.MatchString(line) {
			match := maskRegexp.FindAllStringSubmatch(line, 1)
			mask = readMask(match[0][1])
		}
		if memRegexp.MatchString(line) {
			match := memRegexp.FindAllStringSubmatch(line, 2)
			addr, value := toInt(match[0][1]), toInt(match[0][2])
			num := applyMask(mask, readBits(bitSize, value))
			mem[addr] = fromBits(num)
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}
	println(sum)
}

func part2() {
	file, err := os.Open("./2020/day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	maskRegexp, memRegexp := setupRegexp()

	mem := make(map[int]int)
	var mask []bit

	for scanner.Scan() {
		line := scanner.Text()
		if maskRegexp.MatchString(line) {
			match := maskRegexp.FindAllStringSubmatch(line, 1)
			mask = readMask(match[0][1])
		}
		if memRegexp.MatchString(line) {
			match := memRegexp.FindAllStringSubmatch(line, 2)
			addr, value := toInt(match[0][1]), toInt(match[0][2])
			for _, address := range calculateAddresses(mask, addr) {
				mem[fromBits(address)] = value
			}
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}
	println(sum)
}

func calculateAddresses(mask []bit, addr int) [][]bit {
	bits := readBits(bitSize, addr)
	actual := make([]bit, bitSize)
	for i, m := range mask {
		var bit bit
		if m == unset {
			bit = bits[i]
		} else {
			bit = m
		}
		actual[i] = bit
	}

	return getAddresses(actual)
}

func getAddresses(mask []bit) (result [][]bit) {
	isComplete := func() bool {
		for _, b := range mask {
			if b == floating {
				return false
			}
		}
		return true
	}

	if isComplete() {
		result = append(result, mask)
		return
	}

	for i, b := range mask {
		if b == floating {
			result = append(result, getAddresses(copyAndSet(mask, i, 0))...)
			result = append(result, getAddresses(copyAndSet(mask, i, 1))...)
			return
		}
	}

	return
}

func copyAndSet(src []bit, i int, v bit) []bit {
	output := make([]bit, len(src))
	copy(output, src)
	output[i] = v
	return output
}

func applyMask(mask, arr []bit) []bit {
	result := make([]bit, len(arr))

	for i := 0; i < len(arr); i++ {
		var r bit
		if mask[i] == -1 {
			r = arr[i]
		} else {
			r = mask[i]
		}
		result[i] = r
	}
	return result
}

func toInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func setupRegexp() (*regexp.Regexp, *regexp.Regexp) {
	mask, err := regexp.Compile(maskRegex)
	if err != nil {
		log.Fatal(err)
	}

	mem, err := regexp.Compile(memRegex)
	if err != nil {
		log.Fatal(err)
	}

	return mask, mem
}

type bit int

const (
	floating bit = iota - 1
	unset
	set
)

func readMask(mask string) []bit {
	bits := make([]bit, len(mask))
	for i, m := range mask {
		var b bit
		switch m {
		case 'X':
			b = floating
		case '0', '1':
			b = bit(m - '0')
		}
		bits[i] = b
	}
	return bits
}

func readBits(size, val int) []bit {
	bits := make([]bit, size)
	for i := 0; i < size; i++ {
		var b bit
		if 1<<i&val > 0 {
			b = set
		} else {
			b = unset
		}
		bits[size-i-1] = b
	}
	return bits
}

func fromBits(bits []bit) (dec int) {
	for i, b := range bits {
		x := int(math.Pow(2, float64(len(bits)-i-1)))
		dec += int(b) * x
	}
	return
}
