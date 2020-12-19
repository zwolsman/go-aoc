package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./2020/day19/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(data)
	readRules := strings.Split(strings.Split(text, "\n\n")[0], "\n")
	messages := strings.Split(text, "\n\n")[1]

	rules := make(map[int]string)

	for _, str := range readRules {
		str = strings.ReplaceAll(str, "\"", "")
		temp := strings.Split(str, ": ")
		index, err := strconv.Atoi(temp[0])
		if err != nil {
			log.Fatal(err)
		}
		rules[index] = temp[1]
	}

	dict := make(map[int]string)

	for len(dict) < len(rules) {
		fmt.Println(len(dict), len(rules))
		for ruleId, str := range rules {
			if _, ok := dict[ruleId]; ok {
				continue
			}
			if str == "a" || str == "b" {
				dict[ruleId] = str
				continue
			}

			temp := ""
			success := true
			for _, c := range strings.Split(str, " ") {
				if c == "|" {
					temp += c
					continue
				}
				i, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}

				val, ok := dict[i]
				if !ok {
					success = false
					break
				}

				if strings.Contains(val, "|") {
					val = "(" + val + ")"
				}
				temp += val
			}

			if ruleId == 8 {
				rule42, ok := dict[42]
				if !ok {
					continue
				}
				rule8 := "(" + rule42 + ")+"
				dict[8] = rule8
				continue
			}

			if ruleId == 11 {
				rule42, ok := dict[42]
				if !ok {
					continue
				}
				rule31, ok := dict[31]
				if !ok {
					continue
				}

				rule11 := ""
				for i := 0; i < 20; i++ {
					rule11 += "(" + rule42
				}

				for i := 0; i < 20; i++ {
					rule11 += rule31 + ")?"
				}

				rule11 = rule11[:len(rule11)-1]

				dict[11] = rule11
				continue
			}

			if success && ruleId != 8 && ruleId != 11 {
				dict[ruleId] = temp
			}
		}
	}

	regex, err := regexp.Compile("^" + dict[0] + "$")
	if err != nil {
		log.Fatal(err)
	}

	validMessages := 0
	for _, m := range strings.Split(messages, "\n") {
		if regex.MatchString(m) {
			validMessages++
		}
	}

	fmt.Println(validMessages)
}
