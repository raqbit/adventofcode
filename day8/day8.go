package main

import (
	"errors"
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
	"time"
)

type (
	operation   string
	instruction struct {
		opcode   operation
		argument int
	}
)

const (
	acc operation = "acc"
	jmp operation = "jmp"
	nop operation = "nop"
)

var (
	ErrInfiniteLoop = errors.New("infinite loop")
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	// Remove last newline
	input = input[:len(input)-1]

	if err != nil {
		panic("Could not load input")
	}

	start := time.Now()
	bags := strings.Split(input, "\n")
	bagIndex := parseInstructions(bags)
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

func part1(code []instruction) shared.Result {

	accumulator, err := runComputer(code)

	if err != ErrInfiniteLoop {
		return func() {
			fmt.Printf("Error: did not get infinite loop in part 1")
		}
	}

	return func() {
		fmt.Printf("Accumulator: %d\n", accumulator)
	}
}

func part2(code []instruction) shared.Result {
	changeIndex := -1

	for {
		// Patch code
		changeIndex = patchCode(code, changeIndex)

		// Run computer with patched code
		ret, err := runComputer(code)

		if err == nil {
			return func() {
				fmt.Printf("Accumulator: %d\n", ret)
			}
		}
	}
}

func patchCode(code []instruction, changeIndex int) int {
	if changeIndex >= 0 {
		if code[changeIndex].opcode == nop {
			code[changeIndex].opcode = jmp
		} else {
			code[changeIndex].opcode = nop
		}
	}

	// Change next jmp/nop into nop/jmp
	for i := changeIndex + 1; i < len(code); i++ {
		opcode := code[i].opcode

		if opcode == jmp {
			code[i].opcode = nop
			return i
		} else if opcode == nop {
			code[i].opcode = jmp
			return i
		}
	}

	panic("Could not find new patch to make")
}

func runComputer(code []instruction) (int, error) {
	visited := make(map[int]bool)
	ip := 0
	accumulator := 0

	for {
		// Stop if ip is outside code range
		if ip == len(code) {
			return accumulator, nil
		}

		// Detect infinite loop
		if visited[ip] {
			return accumulator, ErrInfiniteLoop
		}

		inst := code[ip]

		visited[ip] = true

		switch inst.opcode {
		case acc:
			accumulator += inst.argument
		case jmp:
			ip = ip + inst.argument
			continue
		case nop:
		}

		ip++
	}
}

func parseInstructions(lines []string) []instruction {
	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		opcode := parts[0]
		argument, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		instructions[i] = instruction{
			opcode:   operation(opcode),
			argument: argument,
		}
	}

	return instructions
}
