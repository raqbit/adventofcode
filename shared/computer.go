package shared

import (
	"errors"
	"fmt"
)

type Computer interface {
	Start()
	RegisterInstruction(instruction Instruction)
}

type Instruction struct {
	Name    string
	Opcode  int
	ArgC    int
	Execute func(c *IntComputer, argv []int)
}

type IntComputer struct {
	ip           int
	Memory       []int
	instructions map[int]*Instruction
	halt         bool
}

func NewIntComputer() *IntComputer {
	return &IntComputer{instructions: make(map[int]*Instruction)}
}

func (c *IntComputer) RegisterInstruction(inst *Instruction) {
	c.instructions[inst.Opcode] = inst
}

func (c *IntComputer) SetInitialMemory(initMem []int) {
	c.Memory = make([]int, len(initMem))
	copy(c.Memory, initMem)
}

func (c *IntComputer) SetNoun(noun int) {
	c.Memory[1] = noun
}

func (c *IntComputer) SetVerb(verb int) {
	c.Memory[2] = verb
}

func (c *IntComputer) Start() error {
	for {
		opcode := c.Memory[c.ip]
		inst, ok := c.instructions[opcode]
		if !ok {
			return errors.New(fmt.Sprintf("invalid instruction: %d", opcode))
		}
		argv := c.Memory[c.ip+1 : c.ip+inst.ArgC+1]
		inst.Execute(c, argv)
		if c.halt {
			break
		}
		// Increment by one plus the argument count of the instruction
		c.ip += 1 + inst.ArgC
	}

	return nil
}

func (c *IntComputer) FlagHalt() {
	c.halt = true
}

func (c *IntComputer) Reset() {
	c.Memory = nil
	c.ip = 0
	c.halt = false
}
