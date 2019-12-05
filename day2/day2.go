package main

import (
	"fmt"
	"raqb.it/AdventOfCode/shared"
	"raqb.it/AdventOfCode/shared/intcode"
)

func main() {
	input, err := shared.LoadInputFile("input.txt")

	if err != nil {
		panic("Could not load input")
	}

	intcodes, err := intcode.ParseIntcodes(input)

	if err != nil {
		panic("Could not parse input")
	}

	ic := intcode.NewIntComputer()

	part1(ic, intcodes)
	fmt.Println("----")
	part2(ic, intcodes)
}

func part1(ic *intcode.Computer, instructions []int) {
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

func part2(ic *intcode.Computer, instructions []int) {
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
