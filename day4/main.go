package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readAllLines("/Users/mzwolsman/Developer/go-aoc/day4/input.txt")

	passports := parsePassports(lines)

	validPassports := 0
	for _, passport := range passports {
		if validatePassport(passport) {
			validPassports++
		}
	}

	fmt.Println("valid passports", validPassports)
}

func validatePassport(passport string) bool {
	entries := strings.Split(passport, " ")
	passportFields := make(map[string]string)

	requiredFields := map[string]func(input string) bool{
		"byr": byr,
		"iyr": iyr,
		"eyr": eyr,
		"hgt": hgt,
		"hcl": hcl,
		"ecl": ecl,
		"pid": pid,
		//"cid", <-- optional
	}

	for _, entry := range entries {
		fields := strings.Split(entry, ":")

		if len(fields) != 2 {
			fmt.Println("fields is not 2", entry, len(fields), entries, passport)
			return false
		}

		key := fields[0]
		value := fields[1]

		passportFields[key] = value
	}

	for requiredField, check := range requiredFields {
		value, ok := passportFields[requiredField]
		if !ok {
			return false
		}

		if !check(value) {
			return false
		}
	}

	return true
}

func parsePassports(lines []string) []string {
	var passports []string
	for i := 0; i < len(lines); i++ {
		passport := ""
		for j := i; j < len(lines); j++ {
			line := lines[j]

			if line == "" {
				break
			}
			passport += " " + line
			i++
		}
		passports = append(passports, strings.TrimLeft(passport, " "))
	}
	return passports
}

func readAllLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func fourDigitCheck(input string, lower, upper int) bool {
	if len(input) != 4 {
		return false
	}

	year, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	return year >= lower && year <= upper
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func byr(input string) bool {
	return fourDigitCheck(input, 1920, 2002)
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func iyr(input string) bool {
	return fourDigitCheck(input, 2010, 2020)
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func eyr(input string) bool {
	return fourDigitCheck(input, 2020, 2030)
}

// hgt (Height) - a number followed by either cm or in:
//If cm, the number must be at least 150 and at most 193.
//If in, the number must be at least 59 and at most 76.
func hgt(input string) bool {
	check := func(lower, upper int) bool {
		num, err := strconv.Atoi(input[:len(input)-2])
		if err != nil {
			return false
		}
		return num >= lower && num <= upper
	}
	if strings.HasSuffix(input, "cm") {
		return check(150, 193)
	}

	if strings.HasSuffix(input, "in") {
		return check(59, 76)
	}

	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func hcl(input string) bool {
	if len(input) != 7 {
		return false
	}
	if input[0] != '#' {
		return false
	}

	_, err := strconv.ParseInt(input[1:6], 16, 64)
	return err == nil
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func ecl(input string) bool {
	allowed := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}

	for _, entry := range allowed {
		if entry == input {
			return true
		}
	}
	return false
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func pid(input string) bool {
	if len(input) != 9 {
		return false
	}

	_, err := strconv.Atoi(input)
	return err == nil
}
