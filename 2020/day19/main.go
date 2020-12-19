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

	rules := make([]string, len(readRules))

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
		for i, str := range rules {
			if _, ok := dict[i]; ok {
				continue
			}
			if str == "a" || str == "b" {
				dict[i] = str
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

			if success {
				dict[i] = temp
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
