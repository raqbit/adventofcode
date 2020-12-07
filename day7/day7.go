package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	NoOtherBags = "no other"
	ShinyBag    = "shiny gold"
)

type bag struct {
	bags []bagStack
}

type bagStack struct {
	bagType string
	amount  int
}

const BagStatementRegex = `^(?P<type>[a-z ]+) bags contain (?P<rule>.*).$`
const BagRuleRegex = `(?:(?P<amount>\d+) (?P<type>[a-z ]+)|(no other)) bags?`

func main() {
	input, err := shared.LoadInputFile("input.txt")

	// Remove last newline
	input = input[:len(input)-1]

	if err != nil {
		panic("Could not load input")
	}

	start := time.Now()
	bags := strings.Split(input, "\n")
	bagIndex := parseBagRules(bags)
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(bagIndex)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(bagIndex)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(bags map[string]bag) shared.Result {
	counter := 0

	for typ, b := range bags {
		// Ignore top-level shiny bags
		if typ == ShinyBag {
			continue
		}

		// Check if this bag contains a shiny
		if containsShiny(bags, b) {
			counter++
		}
	}

	return func() {
		fmt.Printf("# of bags that contain %s: %d\n", ShinyBag, counter)
	}
}

func containsShiny(bags map[string]bag, bag bag) bool {
	for _, stack := range bag.bags {
		if stack.bagType == ShinyBag || containsShiny(bags, bags[stack.bagType]) {
			return true
		}
	}

	return false
}

func countBags(bags map[string]bag, bag bag) int {
	count := 0
	for _, stack := range bag.bags {
		count += stack.amount + stack.amount*countBags(bags, bags[stack.bagType])
	}

	return count
}

func part2(bags map[string]bag) shared.Result {
	count := countBags(bags, bags[ShinyBag])

	return func() {
		fmt.Printf("# of bags in %s: %d\n", ShinyBag, count)
	}
}

func parseBagRules(bagRules []string) map[string]bag {
	bagIndex := make(map[string]bag)

	for _, bagString := range bagRules {
		bagRule := shared.RegexMatch(regexp.MustCompile(BagStatementRegex), bagString)
		parentBagType := bagRule["type"]
		bagContents := bagRule["rule"]

		matcher := regexp.MustCompile(BagRuleRegex)
		matches := matcher.FindAllStringSubmatch(bagContents, -2)

		stack := make([]bagStack, 0)

		for _, match := range matches {
			if match[3] == NoOtherBags {
				continue
			}

			amount, err := strconv.Atoi(match[1])

			if err != nil {
				panic(fmt.Errorf("could not parse stack amount: %w", err))
			}

			stack = append(stack, bagStack{
				amount:  amount,
				bagType: match[2],
			})
		}

		bagIndex[parentBagType] = bag{bags: stack}
	}

	return bagIndex
}
