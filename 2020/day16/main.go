package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2020/day16/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(data), "\n\n")

	rules := parseRules(parts[0])
	//yourTicket := parseTickets(parts[1])
	nearbyTickets := parseTickets(parts[2])

	part1(rules, nearbyTickets)
}

func part1(rules map[string][]validateFunction, nearbyTickets []ticket) {
	ticketScanningErrorRate := 0

	isValid := func(v int) bool {
		any := false
		for _, validators := range rules {
			for _, rule := range validators {
				if rule(v) {
					any = true
				}
			}
		}
		return any
	}
	for _, t := range nearbyTickets {
		for _, v := range t {
			if !isValid(v) {
				ticketScanningErrorRate += v
				break
			}
		}
	}

	println("Ticket scanning error rate", ticketScanningErrorRate)
}

func parseTickets(input string) (output []ticket) {
	ticketLines := strings.Split(input, "\n")

	for i := 1; i < len(ticketLines); i++ {
		var t ticket
		for _, raw := range strings.Split(ticketLines[i], ",") {
			num, err := strconv.Atoi(raw)
			if err != nil {
				log.Fatal(err)
			}
			t = append(t, num)
		}
		output = append(output, t)
	}
	return
}

const ruleRegex = "(\\w+): (\\d+)\\-(\\d+) or (\\d+)\\-(\\d+)"

func parseRules(data string) map[string][]validateFunction {
	validators := make(map[string][]validateFunction)
	regex, err := regexp.Compile(ruleRegex)
	if err != nil {
		log.Fatal(err)
	}
	rules := regex.FindAllStringSubmatch(data, -1)
	for _, rule := range rules {
		name := rule[1]
		for i := 2; i < len(rule); i += 2 {
			validators[name] = append(validators[name], toValidateFunction(rule[i:i+2]))
		}
	}
	return validators
}

type validateFunction func(in int) bool
type ticket []int

func toValidateFunction(input []string) validateFunction {
	min, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(input[1])
	if err != nil {
		log.Fatal(err)
	}
	return func(in int) bool {
		return in >= min && in <= max
	}
}
