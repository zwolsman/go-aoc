package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	decks := readDecks()

	isFinished := func() bool {
		for _, d := range decks {
			if d.Len() == 0 {
				return true
			}
		}
		return false
	}

	fmt.Printf("playing a game of Combat with %d players\n", len(decks))

	p1, p2 := decks[0], decks[1]
	for round := 1; !isFinished(); round++ {

		fmt.Printf("-- Round %d --\n", round)
		fmt.Printf("Player 1's deck: %v\n", p1.List)
		fmt.Printf("Player 2's deck: %v\n", p2.List)

		c1, c2 := p1.Dequeue(), p2.Dequeue()
		fmt.Printf("Player 1 plays: %d\n", c1)
		fmt.Printf("Player 2 plays: %d\n", c2)

		if c1 > c2 {
			fmt.Println("Player 1 wins the round\n")
			p1.Enqueue(c1)
			p1.Enqueue(c2)
		}
		if c2 > c1 {
			fmt.Println("Player 2 wins the round\n")
			p2.Enqueue(c2)
			p2.Enqueue(c1)
		}
	}

	fmt.Println("\n== Post-game results ==")
	fmt.Printf("Player 1's deck: %v\n", p1.List)
	fmt.Printf("Player 2's deck: %v\n", p2.List)

	fmt.Printf("Player 1's score: %d\n", p1.Score())
	fmt.Printf("Player 2's score: %d\n", p2.Score())

}

func readDecks() (decks []*Deck) {
	data, err := ioutil.ReadFile("./2020/day22/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	players := strings.Split(string(data), "\n\n")

	for _, p := range players {
		d := strings.Split(p, "\n")
		deck := Deck{}
		for i := 1; i < len(d); i++ {
			card, err := strconv.Atoi(d[i])
			if err != nil {
				log.Fatal(err)
			}
			deck.Enqueue(card)
		}
		decks = append(decks, &deck)
	}
	return decks
}

type Deck struct {
	list.List
}

func (d *Deck) Enqueue(v int) {
	d.PushBack(v)
}

func (d *Deck) Dequeue() int {
	element := d.Front()
	val := element.Value.(int)
	d.List.Remove(element)

	return val
}

func (d *Deck) Score() (score int) {
	for i := d.Len(); i > 0; i-- {
		score += i * d.Dequeue()
	}

	return score
}
