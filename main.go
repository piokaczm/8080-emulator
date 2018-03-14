package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/piokaczm/8080-emulator/disassembler"
)

type call struct {
	*disassembler.Instruction
	args []string
}

func newCall(inst *disassembler.Instruction) *call {
	return &call{
		Instruction: inst,
	}
}

func main() {
	data, err := ioutil.ReadFile("invaders/invaders.h")
	if err != nil {
		log.Fatalf(err.Error())
	}

	calls := make([]*call, 0)

	buffer := bytes.NewBuffer(data)
	for {
		singleByte := readByteToHex(buffer)
		if singleByte == "" {
			break
		}
		parsedInstruction, err := buildCall(singleByte, buffer)
		if err != nil {
			log.Fatalf(err.Error())
		}
		calls = append(calls, parsedInstruction)
	}

	for _, inst := range calls {
		fmt.Printf("%s | args: %v\n", inst.Name, inst.args)
	}
}

func readByteToHex(buf *bytes.Buffer) string {
	return fmt.Sprintf("%x", buf.Next(1))
}

func buildCall(hex string, buf *bytes.Buffer) (*call, error) {
	instr, err := disassembler.Decode(hex)
	if err != nil {
		return nil, err
	}

	c := newCall(instr)
	if c.Size > 1 {
		for i := 1; i < c.Size; i++ {
			c.args = append(c.args, readByteToHex(buf))
		}
	}

	return c, nil
}
