package intcode

import (
	"errors"
	"fmt"
	"strconv"
)

type ParamMode int

const (
	Position  ParamMode = 0
	Immediate ParamMode = 1
)

type Instruction struct {
	Name    string
	paramC  int
	Execute func(c *Computer, modes []ParamMode, params []int)
}

type Computer struct {
	ip           int
	Memory       []int
	instructions map[int]*Instruction
	halt         bool
	err          error
}

func (c *Computer) GetValue(params []ParamMode, i int, value int) int {
	paramMode := Position

	if i < len(params) {
		paramMode = params[i]
	}

	switch paramMode {
	case Immediate:
		return value
	case Position:
		return c.Memory[value]
	default:
		fmt.Println("Error: Unknown param mode, returning 0")
	}

	return 0
}

func (c *Computer) Start() error {
	for {
		// Parse instruction at instruction pointer into opcode + params
		opcode, paramModes := parseInstruction(c.Memory[c.ip])

		// Find relevant instruction executor
		inst, ok := c.instructions[opcode]
		if !ok {
			return errors.New(fmt.Sprintf("invalid instruction: %d", opcode))
		}

		// Save instruction pointer
		previousIp := c.ip

		// Get parameters for instruction
		params := c.Memory[c.ip+1 : c.ip+inst.paramC+1]

		// Execute instruction
		inst.Execute(c, paramModes, params)

		// If error is set, halt & print error
		if c.err != nil {
			return fmt.Errorf("error while executing %s instruction: %w", inst.Name, c.err)
		}

		// If halt flag is set, halt
		if c.halt {
			break
		}

		// Only increment instruction pointer if it has not been modified
		if c.ip == previousIp {
			// Increment by one plus the argument count of the instruction
			c.ip += 1 + inst.paramC
		}

	}

	return nil
}

func parseInstruction(inst int) (int, []ParamMode) {
	// Get lower two digits
	opcode := inst % 100

	// Get upper x digits
	modeDigits := inst / 100

	// Store param modes
	paramModes := make([]ParamMode, 0)

	for {
		// No param mode digits left, break
		if modeDigits == 0 {
			break
		}

		// Get least significant digit (LSD)
		digit := modeDigits % 10

		// Append digit as param mode to modes list
		paramModes = append(paramModes, ParamMode(digit))

		// Move to next digit
		modeDigits /= 10
	}

	return opcode, paramModes
}

func (c *Computer) RegisterInstruction(opcode int, inst *Instruction) {
	c.instructions[opcode] = inst
}

func (c *Computer) SetInitialMemory(initMem []int) {
	c.Memory = make([]int, len(initMem))
	copy(c.Memory, initMem)
}

func (c *Computer) SetNoun(noun int) {
	c.Memory[1] = noun
}

func (c *Computer) SetVerb(verb int) {
	c.Memory[2] = verb
}

func (c *Computer) FlagHalt() {
	c.halt = true
}

func (c *Computer) Reset() {
	c.Memory = nil
	c.ip = 0
	c.halt = false
}

func NewIntComputer() *Computer {
	ic := &Computer{instructions: make(map[int]*Instruction)}

	ic.RegisterInstruction(1, &Instruction{
		Name:   "ADD",
		paramC: 3,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			param1 := c.GetValue(modes, 0, params[0])
			param2 := c.GetValue(modes, 1, params[1])

			c.Memory[params[2]] = param1 + param2
		},
	})

	ic.RegisterInstruction(2, &Instruction{
		Name:   "MUL",
		paramC: 3,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			param1 := c.GetValue(modes, 0, params[0])
			param2 := c.GetValue(modes, 1, params[1])

			c.Memory[params[2]] = param1 * param2
		},
	})

	ic.RegisterInstruction(99, &Instruction{
		Name:   "HLT",
		paramC: 0,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			c.FlagHalt()
		},
	})

	ic.RegisterInstruction(4, &Instruction{
		Name:   "OUT",
		paramC: 1,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			fmt.Printf("OUTPUT: %d\n", c.GetValue(modes, 0, params[0]))
		},
	})

	ic.RegisterInstruction(3, &Instruction{
		Name:   "IN",
		paramC: 1,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			fmt.Print("INPUT: ")

			// Read input
			var input string
			_, err := fmt.Scanln(&input)
			if err != nil {
				c.err = err
				return
			}

			// Parse input
			i, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				c.err = err
				return
			}

			c.Memory[params[0]] = int(i)
		},
	})

	ic.RegisterInstruction(5, &Instruction{
		Name:   "JNZ",
		paramC: 2,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			param1 := c.GetValue(modes, 0, params[0])
			param2 := c.GetValue(modes, 1, params[1])

			if param1 != 0 {
				c.ip = param2
			}
		},
	})

	ic.RegisterInstruction(6, &Instruction{
		Name:   "JZ",
		paramC: 2,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			param1 := c.GetValue(modes, 0, params[0])
			param2 := c.GetValue(modes, 1, params[1])

			if param1 == 0 {
				c.ip = param2
			}
		},
	})

	ic.RegisterInstruction(7, &Instruction{
		Name:   "LT",
		paramC: 3,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			param1 := c.GetValue(modes, 0, params[0])
			param2 := c.GetValue(modes, 1, params[1])

			if param1 < param2 {
				c.Memory[params[2]] = 1
			} else {
				c.Memory[params[2]] = 0
			}
		},
	})

	ic.RegisterInstruction(8, &Instruction{
		Name:   "EQ",
		paramC: 3,
		Execute: func(c *Computer, modes []ParamMode, params []int) {
			param1 := c.GetValue(modes, 0, params[0])
			param2 := c.GetValue(modes, 1, params[1])

			if param1 == param2 {
				c.Memory[params[2]] = 1
			} else {
				c.Memory[params[2]] = 0
			}
		},
	})

	return ic
}
