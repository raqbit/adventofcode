package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strings"
)

func main() {
	input, err := shared.LoadInputFile("day2/input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	totalTwo := 0
	totalThree := 0
	for _, line := range lines {
		processed := make(map[rune]struct{}, 0)
		foundTwo := false
		foundThree := false
		for _, letter := range line {
			// Skip runes that were already counted
			if _, ok := processed[letter]; ok {
				continue
			}
			// Count instances of rune
			count := strings.Count(line, string(letter))

			if count == 2 && !foundTwo {
				totalTwo++
				foundTwo = true
			}

			if count == 3 && !foundThree {
				totalThree++
				foundThree = true
			}

			processed[letter] = struct{}{}
		}
	}

	checksum := totalTwo * totalThree
	fmt.Printf("Checksum: %d\n", checksum)

}

func part2(lines []string) {
	// Loop over lines
	for lineIndex, line := range lines {

		// Compare to all lines
		for otherLineIndex, otherLine := range lines {

			// Don't compare to self or stuff already checked
			if otherLineIndex <= lineIndex {
				continue
			}

			// Store common characters
			common := make([]rune, 0)

			// Loop over every character
			for charIndex, char := range []rune(line) {
				otherChar := []rune(otherLine)[charIndex]
				// Check if the characters of this index are the same
				if otherChar == char {
					common = append(common, char)
				}
			}
			// If there's only a single difference
			if len(line)-len(common) == 1 {
				fmt.Printf("Common Characters: '%s'\n", string(common))
				return
			}
		}
	}
}
