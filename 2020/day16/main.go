package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
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
	yourTicket := parseTickets(parts[1])[0]
	nearbyTickets := parseTickets(parts[2])

	part1(rules, nearbyTickets)
	part2(rules, nearbyTickets, yourTicket)
}

func part1(rules validateMap, nearbyTickets []ticket) {
	ticketScanningErrorRate := 0

	isValid := func(v int) bool {
		for _, validators := range rules {
			a, b := validators[0], validators[1]
			if a(v) || b(v) {
				return true
			}
		}
		return false
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

func part2(rules validateMap, nearbyTickets []ticket, yourTicket ticket) {
	var validTickets []ticket

	isValid := func(v int) bool {
		for _, validators := range rules {
			a, b := validators[0], validators[1]
			if a(v) || b(v) {
				return true
			}
		}
		return false
	}

	for _, t := range nearbyTickets {
		isValidTicket := true
		for _, v := range t {
			if !isValid(v) {
				isValidTicket = false
				break
			}
		}
		if isValidTicket {
			validTickets = append(validTickets, t)
		}
	}

	worksForAll := func(index int, validators [2]validateFunction) bool {
		a, b := validators[0], validators[1]
		for _, t := range validTickets {
			field := t[index]
			if !a(field) && !b(field) {
				return false
			}
		}
		return true
	}

	possibilities := make([][]string, len(yourTicket))
	for i := 0; i < len(yourTicket); i++ {
		var options []string
		for rule, validators := range rules {
			if worksForAll(i, validators) {
				options = append(options, rule)
			}
		}
		possibilities[i] = options
	}

	ticketDef := make(map[string]int)

	for i := 0; i < len(yourTicket); i++ {
		for field, options := range possibilities {
			if len(options) == i-1 {
				sort.Strings(options)
				for _, o := range options {
					if _, ok := ticketDef[o]; !ok {
						ticketDef[o] = field
						break
					}
				}
			}
		}
	}

	sum := 1
	for key, field := range ticketDef {
		if strings.HasPrefix(key, "departure") {
			sum *= yourTicket[field]
		}

	}
	fmt.Println(sum)
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

const ruleRegex = "(.*): (\\d+)\\-(\\d+) or (\\d+)\\-(\\d+)"

func parseRules(data string) map[string][2]validateFunction {
	validators := make(validateMap)
	regex, err := regexp.Compile(ruleRegex)
	if err != nil {
		log.Fatal(err)
	}
	rules := regex.FindAllStringSubmatch(data, -1)
	for _, rule := range rules {
		name := rule[1]
		var ruleValidators [2]validateFunction
		for i := 2; i < len(rule); i += 2 {
			ruleValidators[i/2-1] = toValidateFunction(rule[i : i+2])
		}
		validators[name] = ruleValidators
	}
	return validators
}

type validateFunction func(in int) bool
type validateMap map[string][2]validateFunction
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
