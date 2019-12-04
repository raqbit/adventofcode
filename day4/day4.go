package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = "123257-647015"

func main() {
	from, to, err := parseRange(input)

	if err != nil {
		panic("Unable to parse input")
	}

	part1(from, to)
	fmt.Println("----")
	part2(from, to)
}

func part1(from int, to int) {
	correct := 0
	for num := from; num <= to; num++ {
		str := strconv.Itoa(num)
		increases := digitsIncrease(str)
		hasConsecutive := hasNConsecutiveChars(str, 2, true)
		if increases && hasConsecutive {
			correct++
		}
	}

	fmt.Printf("Amount of numbers that meet requirements: %d\n", correct)
}

func part2(from int, to int) {
	correct := 0
	for num := from; num <= to; num++ {
		str := strconv.Itoa(num)
		if digitsIncrease(str) && hasNConsecutiveChars(str, 2, false) {
			correct++
		}
	}

	fmt.Printf("Amount of numbers that meet requirements: %d\n", correct)
}

func hasNConsecutiveChars(str string, count int, allowExcess bool) bool {
	// Amount of characters counted
	charCount := 0

	// Previous character
	previous := -1

	// Loop through array
	for i := 0; i < len(str); i++ {

		// Current character
		char := int(str[i])

		if char == previous {
			// New of the same char counted
			charCount++

			// If we allow excess and the counter counted enough consecutive digits
			if allowExcess && charCount == count {
				return true
			}
		} else {
			// Character is different from previous (or previous was -1)

			// We previously counted 'count' characters and don't allow excess
			if !allowExcess && charCount == count {
				return true
			}

			// Reset counter to 1 because we already counted 1 of this type
			charCount = 1
		}

		previous = char
	}

	if charCount == count {
		return true
	}

	return false
}

func digitsIncrease(num string) bool {
	maxNum := 0
	for i := 0; i < len(num); i++ {
		char := int(num[i])

		// A character is lower than previously found
		if char < maxNum {
			return false
		}

		maxNum = char
	}

	return true
}

func parseRange(rangeStr string) (int, int, error) {
	rangeParts := strings.Split(rangeStr, "-")

	from, err := strconv.ParseInt(rangeParts[0], 10, 32)

	if err != nil {
		return 0, 0, err
	}

	to, err := strconv.ParseInt(rangeParts[1], 10, 32)

	if err != nil {
		return 0, 0, err
	}

	return int(from), int(to), nil
}
