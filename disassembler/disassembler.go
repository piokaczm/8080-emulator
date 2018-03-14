package disassembler

import (
	"bytes"
	"fmt"
	"strconv"
)

type instruction struct {
	name     string
	size     int
	flags    string
	function string
	args     []string
}

// Decode reads provided data and disasseles hex values to 8080 instructions
func Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)
	var ordinal int64

	for {
		singleByte := singleByteToHex(buffer)
		if singleByte == "" {
			break
		}
		inst, err := decodeSingleHex(singleByte, buffer)
		if err != nil {
			return err
		}
		inst.print(ordinal)
		ordinal = incrementOrdinal(ordinal, inst)
	}

	return nil
}

func incrementOrdinal(ordinal int64, inst *instruction) int64 {
	return ordinal + 1 + int64(len(inst.args)) // ordinal + instruction (1) + args
}

func newInstruction(name string, size int, flags, function string) *instruction {
	in := &instruction{
		name:     name,
		size:     size,
		flags:    flags,
		function: function,
	}

	if in.size > 1 {
		in.args = make([]string, in.size-1)
	}

	return in
}

func (in *instruction) hasArgs() bool {
	return in.size > 1
}

func (in *instruction) print(ordinal int64) {
	ordinalHex := strconv.FormatInt(ordinal, 16)
	if in.hasArgs() {
		fmt.Printf("[%s] %s | args: %v\n", ordinalHex, in.name, in.args)
	} else {
		fmt.Printf("[%s] %s\n", ordinalHex, in.name)
	}
}

func decodeSingleHex(hex string, buf *bytes.Buffer) (*instruction, error) {
	in, ok := opcodes[hex]
	if !ok {
		return nil, fmt.Errorf("no opcode %q found", hex)
	}

	// read args
	if in.hasArgs() {
		for i := in.size - 2; i >= 0; i-- {
			in.args[i] = singleByteToHex(buf)
		}
	}
	return in, nil
}

func singleByteToHex(buf *bytes.Buffer) string {
	return fmt.Sprintf("%x", buf.Next(1))
}
