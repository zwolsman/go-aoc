package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./2020/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if j == i {
				continue
			}
			for k := 0; k < len(numbers); k++ {
				if k == j {
					continue
				}

				a := numbers[i]
				b := numbers[j]
				c := numbers[k]

				if a+b+c == 2020 {
					fmt.Println(a * b * c)
					return
				}
			}
		}
	}
}
