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

	part1(input)
	fmt.Println("----")
	part2(input)
}

func part1(input string) {
	result := react(input)

	fmt.Printf("Final string: %s\n", result)
	fmt.Printf("Length: %d\n", len(result))
}

func part2(input string) {
	checkedTypes := map[rune]struct{}{}
	for _, unitType := range input {
		filteredInput := input
		lowerUnitType := unicode.ToLower(unitType)

		// Filter already checked types
		if _, ok := checkedTypes[rune(lowerUnitType)]; ok {
			continue
		}

		filteredInput = strings.Replace(filteredInput, string(lowerUnitType), "", -1)
		filteredInput = strings.Replace(filteredInput, string(unicode.ToUpper(lowerUnitType)), "", -1)

		result := react(filteredInput)
		fmt.Printf("Filtering %s: %d\n", string(lowerUnitType), len(result))
		checkedTypes[rune(lowerUnitType)] = struct{}{}
	}
}

func react(input string) string {
	currString := input
	newString := ""
	index := 0

	for {
		currChar := currString[index]
		nextChar := currString[index+1]

		increment := 1

		// Aa || aA
		if unicode.IsUpper(rune(currChar)) && rune(nextChar) == unicode.ToLower(rune(currChar)) ||
			unicode.IsLower(rune(currChar)) && rune(nextChar) == unicode.ToUpper(rune(currChar)) {
			// If we did find it, skip these
			increment = 2
		} else {
			// Else simply add it to the new string
			newString += string(currChar)
		}

		// Do not process last character since it does not have a next char
		if index+increment+1 >= len(currString) {
			newString += string(nextChar)
			if newString != currString {
				currString = newString
			} else {
				break
			}
			newString = ""
			index = 0
			continue
		}
		index += increment
	}

	return newString
}
