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

	ic := shared.NewIntComputer()

	ic.RegisterInstruction(&shared.Instruction{
		Name:   "ADD",
		Opcode: 1,
		ArgC:   3,
		Execute: func(c *shared.IntComputer, argv []int) {
			c.Memory[argv[2]] = c.Memory[argv[0]] + c.Memory[argv[1]]
			fmt.Printf("$%d + $%d => $%d\n", argv[2], argv[0], argv[1])
		},
	})

	ic.RegisterInstruction(&shared.Instruction{
		Name:   "MULT",
		Opcode: 2,
		ArgC:   3,
		Execute: func(c *shared.IntComputer, argv []int) {
			c.Memory[argv[2]] = c.Memory[argv[0]] * c.Memory[argv[1]]
			fmt.Printf("$%d * $%d => $%d\n", argv[2], argv[0], argv[1])
		},
	})

	ic.RegisterInstruction(&shared.Instruction{
		Name:   "HALT",
		Opcode: 99,
		ArgC:   0,
		Execute: func(c *shared.IntComputer, argv []int) {
			c.FlagHalt()
			fmt.Println("HALT")
		},
	})

	part1(ic, intcodes)
	fmt.Println("----")
	part2(ic, intcodes)
}

func part1(ic *shared.IntComputer, instructions []int) {
	ic.SetInitialMemory(instructions)
	ic.SetNoun(12)
	ic.SetVerb(2)

	err := ic.Start()

	if err != nil {
		fmt.Printf("IntComputer error: %v\n", err)
		return
	}
	fmt.Printf("First cell of final program memory: %d\n", ic.Memory[0])
	ic.Reset()
}

func part2(ic *shared.IntComputer, instructions []int) {
AllInputs:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			ic.SetInitialMemory(instructions)
			ic.SetNoun(noun)
			ic.SetVerb(verb)

			err := ic.Start()

			if err != nil {
				fmt.Printf("IntComputer error: %v\n", err)
				return
			}

			if ic.Memory[0] == 19690720 {
				fmt.Printf("Found valid input, noun: %d, verb: %d\n", noun, verb)
				fmt.Printf("Checksum: %d\n", 100*noun+verb)
				break AllInputs
			}

			ic.Reset()
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
