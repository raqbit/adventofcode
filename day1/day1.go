package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
	"time"
)

const Target = 2020

type result func()

var noop result = func() {
	println("no results")
}

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")
	numbers := parseLines(lines)

	start := time.Now()
	res := part1(numbers)
	duration := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %d ms, %d us\n", duration.Milliseconds(), duration.Microseconds())

	fmt.Println("----")

	start = time.Now()
	res = part2(numbers)
	duration = time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %d ms, %d us\n", duration.Milliseconds(), duration.Microseconds())
}

func part1(entries []int) result {
	set := make(map[int]bool)

	for _, entry := range entries {
		other := Target - entry
		if _, ok := set[other]; ok {
			return func() {
				fmt.Printf("Product: %d * %d = %d\n", entry, other, entry*other)
			}
		}

		set[entry] = true
	}

	return noop
}

func part2(entries []int) result {
	set := make(map[int]bool)

	for i, a := range entries {
		for j, b := range entries {
			if i == j {
				continue
			}

			ab := a + b

			if ab >= Target {
				continue
			}

			c := Target - ab

			if _, ok := set[c]; ok {
				return func() {
					fmt.Printf("Product: %d * %d * %d = %d\n", a, b, c, a*b*c)
				}
			}

			set[b] = true
		}
	}

	return noop
}

// Parses the input lines as integers
func parseLines(lines []string) []int {
	numbers := make([]int, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			panic("Input line is not a number")
		}
		numbers[i] = int(num)
	}
	return numbers
}
