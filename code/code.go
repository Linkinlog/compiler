package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
	OpWithoutDescription
)

type Definition struct {
	Name string
	Description string

	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {
		"OpConstant",
		"OpConstant is an instruction that pushes a constant to the stack.",
		[]int{2},
	},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("undefined opcode %v", op)
	}

	return def, nil
}

// Make creates an instruction with the given opcode and operands.
// An instruction consists of an op code and a sequence of operands.
func Make(op Opcode, operands ...int) (instruction []byte) {
	def, err := Lookup(byte(op))
	if err != nil {
		return instruction
	}

	instruction = allocateInstructionSlice(def.OperandWidths)

	// The first byte is the opcode.
	instruction[0] = byte(op)

	writeOperands(instruction, def.OperandWidths, operands...)

	return instruction
}

func writeOperands(instruction []byte, operandWidths []int, operands ...int) {
	offset := 1

	for i, o := range operands {
		width := operandWidths[i]
		writeOperand(instruction, o, width)
		offset += width
	}
}

func writeOperand(dest []byte, operand int, width int) {
	switch width {
	case 2:
		binary.BigEndian.PutUint16(dest, uint16(operand))
	}
}

func allocateInstructionSlice(operandWidths []int) (instruction []byte) {
	instructionLen := 1 + sum(operandWidths)

	return make([]byte, instructionLen)
}

func sum(slice []int) (sum int) {
	for _, s := range slice {
		sum += s
	}
	return sum
}
