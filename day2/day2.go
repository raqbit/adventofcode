package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const InputRegex = `^(?P<first>\d+)-(?P<second>\d+) (?P<char>[a-z]): (?P<password>.*)$`

type policy struct {
	first  int
	second int
	char   rune
}

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")
	policies, passwords := parseLines(lines)

	start := time.Now()
	res := part1(policies, passwords)
	duration := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", duration.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(policies, passwords)
	duration = time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", duration.String())
}

func part1(policies []policy, passwords []string) shared.Result {
	validCount := 0
	for i, password := range passwords {
		pol := policies[i]

		count := strings.Count(password, string(pol.char))

		if count >= pol.first && count <= pol.second {
			validCount++
		}
	}

	return func() {
		fmt.Printf("Valid Passwords: %d\n", validCount)
	}
}

func part2(policies []policy, passwords []string) shared.Result {
	validCount := 0
	for i, password := range passwords {
		pol := policies[i]

		if (rune(password[pol.first-1]) == pol.char) != (rune(password[pol.second-1]) == pol.char) {
			validCount++
		}
	}

	return func() {
		fmt.Printf("Valid Passwords: %d\n", validCount)
	}
}

func parseLines(lines []string) ([]policy, []string) {
	policies := make([]policy, len(lines)-1)
	passwords := make([]string, len(lines)-1)

	matcher := regexp.MustCompile(InputRegex)

	for i, line := range lines {
		if line == "" {
			continue
		}

		matches := shared.RegexMatch(matcher, line)

		var err error
		var first, second int64

		if first, err = strconv.ParseInt(matches["first"], 10, 32); err != nil {
			panic("Min is not a number")
		}

		if second, err = strconv.ParseInt(matches["second"], 10, 32); err != nil {
			panic("Max is not a number")
		}

		policies[i] = policy{
			first:  int(first),
			second: int(second),
			char:   rune(matches["char"][0]),
		}

		passwords[i] = matches["password"]
	}
	return policies, passwords
}
