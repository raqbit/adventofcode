package main

import (
	"fmt"
	"math"
	"raqb.it/AdventOfCode/shared"
	"sort"
	"strings"
	"time"
)

const (
	Front rune = 'F'
	Back  rune = 'B'
	Right rune = 'R'
	Left  rune = 'L'
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	lines := strings.Split(input, "\n")

	start := time.Now()
	parseTime := time.Since(start)
	fmt.Printf("Parsed in %s\n", parseTime.String())

	fmt.Println("----")

	start = time.Now()
	res := part1(lines)
	part1Time := time.Since(start)
	res()
	fmt.Printf("Finished part 1 in %s\n", part1Time.String())

	fmt.Println("----")

	start = time.Now()
	res = part2(lines)
	part2Time := time.Since(start)
	res()
	fmt.Printf("Finished part 2 in %s\n", part2Time.String())

	fmt.Println("----")

	fmt.Printf("Total time: %s\n", parseTime+part1Time+part2Time)
}

func part1(passes []string) shared.Result {
	highest := 0

	for _, pass := range passes {
		if pass == "" {
			continue
		}

		seatID := getPassSeatID(pass)
		if seatID > highest {
			highest = seatID
		}
	}

	return func() {
		fmt.Printf("Highest Seat ID: %d\n", highest)
	}
}

func getPassSeatID(pass string) int {
	rowPart := pass[:7]
	colPart := pass[7:]

	row := calculatePlace(rowPart, Front, Back)
	col := calculatePlace(colPart, Left, Right)

	return row*8 + col
}

func calculatePlace(part string, frontRune, backRune rune) int {
	currMin := 0
	currMax := int(math.Pow(2, float64(len(part))))

	for _, c := range part {
		if c == frontRune {
			currMax -= (currMax - currMin) / 2
		} else if c == backRune {
			currMin += (currMax - currMin) / 2
		}
	}

	return currMin
}

func part2(passes []string) shared.Result {
	foundSeats := make([]int, 0)

	for _, pass := range passes {
		if pass == "" {
			continue
		}

		foundSeats = append(foundSeats, getPassSeatID(pass))
	}

	sort.Ints(foundSeats)

	prevSeat := 0

	for i, seat := range foundSeats {
		if i != 0 && seat-1 != prevSeat {
			return func() {
				fmt.Printf("Hole at %d\n", prevSeat+1)
			}
		}
		prevSeat = seat
	}

	return shared.NoopResult
}
