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

	fmt.Println("Press 1 for part 1, press 5 for part 2")
	run(ic, intcodes)
}

func run(ic *intcode.Computer, instructions []int) {
	ic.SetInitialMemory(instructions)

	err := ic.Start()

	if err != nil {
		fmt.Printf("IntComputer error: %v\n", err)
		return
	}

	ic.Reset()
}
