package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strings"
	"unicode"
)

func main() {
	input, err := shared.LoadInputFile("day5/input.txt")

	if err != nil {
		panic("Could not load input")
	}

	result := part1(input)
	fmt.Println("----")
	// Using the result of part1 in part2 because it has already been done reacting without removing characters
	part2(result)
}

func part1(input string) string {
	result := react(input)

	fmt.Printf("Final string: %s\n", result)
	fmt.Printf("Length: %d\n", len(result))
	return result
}

func part2(input string) {
	checkedTypes := make(map[rune]bool)
	lowest := len(input)
	for _, unitType := range input {
		filteredInput := input
		lowerUnitType := unicode.ToLower(unitType)

		// Filter already checked types
		if _, ok := checkedTypes[lowerUnitType]; ok {
			continue
		}

		filteredInput = strings.Replace(filteredInput, string(lowerUnitType), "", -1)
		filteredInput = strings.Replace(filteredInput, string(unicode.ToUpper(lowerUnitType)), "", -1)

		result := react(filteredInput)
		resultLen := len(result)
		if resultLen < lowest {
			lowest = resultLen
		}
		checkedTypes[lowerUnitType] = true
	}

	fmt.Printf("Smallest polymer: %d\n", lowest)
}

func react(input string) string {
	stack := make([]rune, 0)

	for _, currChar := range input {
		n := int32(len(stack) - 1) // Top char

		if n < 0 {
			stack = append(stack, currChar)
			continue
		}

		lastChar := stack[n]

		// Checking for aA || Aa
		// 32 is the distance between the lower and uppercase in ascii/unicode
		if currChar-32 == lastChar || currChar+32 == lastChar {
			// Pop last char from stack
			stack = stack[:n]
		} else {
			// Add current character to stack
			stack = append(stack, currChar)
		}
	}

	return string(stack)
}
