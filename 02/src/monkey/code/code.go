package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	// start from 0, increase by byte
	OpConstant Opcode = iota
)

// definition for opcode
type Definition struct {
	Name          string
	OperandWidths []int // the number of bytes each operand takes up
}

var definitions = map[Opcode]*Definition{
	// limits constant's maximum number to 65536
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}
	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)
	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}
