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
	//part1()
	//println("*** part 2 ***")
	part2()
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
func check(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type Player int

const (
	Player1 = iota + 1
	Player2
)

func playGame(p1, p2 *Deck, game int) (winner Player) {
	var p1History, p2History [][]int
	isRepeated := func(p1, p2 []int) bool {

		for _, h1 := range p1History {
			if check(h1, p1) {
				return true
			}
		}

		for _, h2 := range p2History {
			if check(h2, p2) {
				return true
			}
		}
		return false
	}

	isFinished := func() bool {
		return p1.Len() == 0 || p2.Len() == 0
	}

	subDeck := func(d *Deck, n int) *Deck {
		d2 := Deck{}
		e := d.Front()
		for cur := 0; cur < n; cur++ {
			d2.Enqueue(e.Value.(int))
			e = e.Next()
		}
		return &d2
	}

	for round := 1; !isFinished(); round++ {
		//fmt.Printf("-- Round %d (Game %d) --\n", round, game)
		//fmt.Printf("Player 1's deck: %v\n", p1.Cards())
		//fmt.Printf("Player 2's deck: %v\n", p2.Cards())

		h1, h2 := p1.Cards(), p2.Cards()
		if isRepeated(h1, h2) {
			//fmt.Println("Player 1 wins the game because it has been played before.\n")
			return Player1
		}
		p1History = append(p1History, h1)
		p2History = append(p1History, h2)

		c1, c2 := p1.Dequeue(), p2.Dequeue()
		//fmt.Printf("Player 1 plays: %d\n", c1)
		//fmt.Printf("Player 2 plays: %d\n", c2)

		if c1 <= p1.Len() && c2 <= p2.Len() {
			//fmt.Println("Playing a sub-game to determine the winner...")
			d1 := subDeck(p1, c1)
			d2 := subDeck(p2, c2)
			result := playGame(d1, d2, game+1)

			switch result {
			case Player1:
				//fmt.Println("Player 1 wins the round\n")
				p1.Enqueue(c1)
				p1.Enqueue(c2)
			case Player2:
				//fmt.Println("Player 2 wins the round\n")
				p2.Enqueue(c2)
				p2.Enqueue(c1)
			}
			continue
		}

		if c1 > c2 {
			//fmt.Println("Player 1 wins the round\n")
			p1.Enqueue(c1)
			p1.Enqueue(c2)
		}
		if c2 > c1 {
			//fmt.Println("Player 2 wins the round\n")
			p2.Enqueue(c2)
			p2.Enqueue(c1)
		}
	}

	if game == 1 {
		fmt.Println("\n== Post-game results ==")
		fmt.Printf("Player 1's deck: %v\n", p1.Cards())
		fmt.Printf("Player 2's deck: %v\n", p2.Cards())

		fmt.Printf("Player 1's score: %d\n", p1.Score())
		fmt.Printf("Player 2's score: %d\n", p2.Score())
	}

	if p1.Len() == 0 {
		return Player2
	} else if p2.Len() == 0 {
		return Player1
	} else {
		panic("Neither of them won??")
	}
}

func part2() {
	decks := readDecks()
	p1, p2 := decks[0], decks[1]
	playGame(p1, p2, 1)
	// not 31050
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

func (d *Deck) Cards() []int {
	out := make([]int, d.Len())
	i := 0
	for e := d.Front(); e != nil; e = e.Next() {
		out[i] = e.Value.(int)
		i++
	}
	return out
}

func (d *Deck) Score() (score int) {
	for i := d.Len(); i > 0; i-- {
		score += i * d.Dequeue()
	}

	return score
}
