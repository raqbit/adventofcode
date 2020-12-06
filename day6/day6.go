package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strings"
	"time"
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	// Remove last newline
	input = input[:len(input)-1]

	if err != nil {
		panic("Could not load input")
	}

	groups := strings.Split(input, "\n\n")

	start := time.Now()
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(groups)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(groups)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(groups []string) shared.Result {
	answerCount := 0

	for _, g := range groups {
		for c := 'a'; c <= 'z'; c++ {
			if strings.ContainsRune(g, c) {
				answerCount++
			}
		}
	}

	return func() {
		fmt.Printf("Sum of counts: %d\n", answerCount)
	}
}
func part2(groups []string) shared.Result {
	answerCount := 0

	for _, g := range groups {
		personCount := len(strings.Split(g, "\n"))
		runes := []rune(g)

		for i := 'a'; i <= 'z'; i++ {
			count := strings.Count(string(runes), string(i))
			if count == personCount {
				answerCount++
			}
		}
	}

	return func() {
		fmt.Printf("Sum of counts: %d\n", answerCount)
	}
}
