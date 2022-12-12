package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var in []byte

func main() {
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) any {
	return nil
}

func part2(in []byte) any {
	return nil
}
