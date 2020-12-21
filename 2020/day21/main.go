package main

import (
	"bufio"
	"fmt"
	"github.com/deckarep/golang-set"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var containsRegex = regexp.MustCompile(`\(contains (.*)\)`)

func main() {
	file, err := os.Open("./2020/day21/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	allergensSet := mapset.NewSet()
	var food []Food

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := containsRegex.FindAllStringSubmatch(line, -1)
		var allergens []string
		if len(matches) == 1 {
			match := matches[0][1]
			line = strings.Split(line, " (")[0]
			allergens = strings.Split(match, ", ")

			for _, a := range allergens {
				allergensSet.Add(a)
			}
		}

		food = append(food, Food{
			allergens:   allergens,
			ingredients: strings.Split(line, " "),
		})
	}

	allergenMap := make(map[string]string)

	createSet := func(f Food) mapset.Set {
		set := mapset.NewSet()

		for _, i := range f.ingredients {
			if _, ok := allergenMap[i]; ok {
				continue
			}
			set.Add(i)
		}
		return set
	}

	allergenOptions := make(map[string]mapset.Set)
	for allergen := range allergensSet.Iter() {
		println("initial check allergen", allergen.(string))
		var sets []mapset.Set
		for _, f := range food {
			if f.ContainsAllergen(allergen.(string)) {
				sets = append(sets, createSet(f))
			}
		}

		s := sets[0]
		for i := 1; i < len(sets); i++ {
			s = s.Intersect(sets[i])
		}

		allergenOptions[allergen.(string)] = s
		fmt.Println(s)
		fmt.Println()
	}

	for len(allergenMap) < len(allergenOptions) {
		for allergen, set := range allergenOptions {

			fmt.Printf("checking %s -> %v\n", allergen, set)
			if set.Cardinality() == 1 {
				ingredient := set.Pop().(string)
				allergenMap[ingredient] = allergen

				for a2, s2 := range allergenOptions {
					if a2 == allergen {
						continue
					}
					s2.Remove(ingredient)
				}
			}
		}
	}

	sum := 0
	for _, f := range food {
		for _, i := range f.ingredients {

			if _, ok := allergenMap[i]; !ok {
				sum++
			}
		}
	}

	var allergens []string
	byAllergen := make(map[string]string)
	for k, v := range allergenMap {
		byAllergen[v] = k
		allergens = append(allergens, v)
	}
	sort.Strings(allergens)

	var dangerousIngredientList []string
	for _, allergen := range allergens {
		dangerousIngredientList = append(dangerousIngredientList, byAllergen[allergen])
	}
	fmt.Println(sum)
	fmt.Println(strings.Join(dangerousIngredientList, ","))
}

type Food struct {
	allergens   []string
	ingredients []string
}

func (f *Food) ContainsAllergen(allergen string) bool {
	for _, a := range f.allergens {
		if a == allergen {
			return true
		}
	}
	return false
}

func (f *Food) String() string {
	return fmt.Sprintf("Food(ingredients=%v, allergens=%v)", f.ingredients, f.allergens)
}
