package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"strconv"
	"strings"
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	intcodes, err := parseIntcodes(input)

	if err != nil {
		panic("Could not parse input")
	}

	part1(intcodes)
	fmt.Println("----")
	part2(intcodes)
}

func part1(instructions []int) {
	// Set noun
	instructions[1] = 12
	// Set verb
	instructions[2] = 2

	finalMem := runProgram(instructions)

	fmt.Printf("First cell of final program memory: %d\n", finalMem[0])
}

func part2(instructions []int) {
AllInputs:
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			// Set noun
			instructions[1] = noun
			// Set verb
			instructions[2] = verb

			finalMem := runProgram(instructions)

			if finalMem[0] == 19690720 {
				fmt.Printf("Found valid input, noun: %d, verb: %d\n", noun, verb)
				fmt.Printf("Checksum: %d\n", 100*noun+verb)
				break AllInputs
			}
		}
	}
}

func runProgram(instructions []int) []int {
	programMemory := make([]int, len(instructions))
	copy(programMemory, instructions)

	ip := 0
run:
	for {
		// Opcode
		opcode := programMemory[ip]

		var param1 int
		var param2 int

		// HALT opcode has no params
		if ip+3 < len(programMemory) {
			// First input parameter
			param1 = programMemory[programMemory[ip+1]]
			param2 = programMemory[programMemory[ip+2]]
		}

		// Output address
		outputAddr := programMemory[ip+3]

		switch opcode {
		case 1:
			// Opcode: ADD
			programMemory[outputAddr] = param1 + param2
		case 2:
			// Opcode: MULT
			programMemory[outputAddr] = param1 * param2
		case 99:
			// Opcode: HALT
			break run
		}

		// Move instruction pointer
		ip += 4
	}

	return programMemory
}

// Parses the input lines as integers
func parseIntcodes(input string) ([]int, error) {
	unparsedIntcodes := strings.Split(strings.TrimSpace(input), ",")
	intcodes := make([]int, len(unparsedIntcodes))
	for i, uic := range unparsedIntcodes {
		code, err := strconv.ParseInt(uic, 10, 32)
		if err != nil {
			return nil, err
		}
		intcodes[i] = int(code)
	}
	return intcodes, nil
}
