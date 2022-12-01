package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	in, err := os.ReadFile("./2022/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(in []byte) any {
	panic("implement me")
}

func part2(in []byte) any {
	panic("implement me")
}
