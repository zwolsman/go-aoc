package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const PublicSubjectNumber = 7

func main() {
	data, err := ioutil.ReadFile("./2020/day25/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	temp := strings.Split(string(data), "\n")
	p1, err := strconv.Atoi(temp[0])
	if err != nil {
		log.Fatal(err)
	}
	p2, err := strconv.Atoi(temp[1])
	if err != nil {
		log.Fatal(err)
	}

	loopSize := 0
	n := 1
	for n != p1 && n != p2 {
		n = (n * PublicSubjectNumber) % 20201227
		loopSize++
	}

	if n == p1 {
		println(encrypt(p2, loopSize))
	} else {
		println(encrypt(p1, loopSize))
	}

}

func encrypt(sn, loopSize int) int {
	n := 1
	for i := 0; i < loopSize; i++ {
		n *= sn
		n %= 20201227
	}
	return n
}
