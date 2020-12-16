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

func part1(rules []validateFunction, nearbyTickets []ticket) {
	ticketScanningErrorRate := 0

	isValid := func(v int) bool {
		any := false
		for _, rule := range rules {
			if rule(v) {
				any = true
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

const ruleRegex = "(\\d+)\\-(\\d+)"

func parseRules(rules string) (validators []validateFunction) {
	regex, err := regexp.Compile(ruleRegex)
	if err != nil {
		log.Fatal(err)
	}
	for _, rule := range strings.Split(rules, "\n") {
		groups := regex.FindAllStringSubmatch(rule, 2)
		for _, validateRange := range groups {
			validators = append(validators, toValidateFunction(validateRange))
		}
	}
	return
}

type validateFunction func(in int) bool
type ticket []int

func toValidateFunction(input []string) validateFunction {
	min, err := strconv.Atoi(input[1])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(input[2])
	if err != nil {
		log.Fatal(err)
	}
	return func(in int) bool {
		return in >= min && in <= max
	}
}
