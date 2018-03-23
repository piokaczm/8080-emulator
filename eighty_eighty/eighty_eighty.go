package eighty_eighty

import "fmt"

type state struct {
	a          uint8
	b          uint8
	c          uint8
	d          uint8
	e          uint8
	h          uint8
	l          uint8
	sc         uint16
	pc         uint16
	mem        []uint8
	cc         *condCodes
	int_enable uint8
}

type condCodes struct {
	z   uint8
	s   uint8
	p   uint8
	cy  uint8
	ac  uint8
	pad uint8
}

// New returns fresh state for 8080 emulator
func New() *state {
	return &state{
		cc:  &condCodes{},
		mem: make([]uint8, 65536),
	}
}

// Emulate DOC TBA ---
func (s *state) Emulate() error {
	opCode := s.mem[s.pc:]
	defer func() { s.pc++ }()

	switch opCode[0] {
	case 0x00: // NOP
	case 0x01: // LXI B,D16
		s.c = opCode[1]
		s.b = opCode[2]
		s.pc += 2
	case 0x02: // STAX B
	case 0x03: // INX B
		s.b++
		s.c++
	case 0x04: // INR B
	case 0x05: // DCR B
	case 0x06: // MVI B, D8
		s.b = opCode[1]
		s.pc++
	case 0x07: // RLC
	case 0x08: // -
	case 0x09: // DAD B
	case 0x1a: // LDAX B
	default:
		return fmt.Errorf("bad opcode %v", opCode)
	}

	return nil
}
