package eighty_eighty

import "fmt"

const (
	// registers fast access
	a = iota
	b
	c
	d
	e
	h
	l

	// registers pairs fast access
	bc = iota
	de
)

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
		s.lxi(opCode[1], opCode[2], bc)
	case 0x02: // STAX B
		s.mem[addr(s.b, s.c)] = s.a
	case 0x03: // INX B
		s.b++
		s.c++
	case 0x04: // INR B
	case 0x05: // DCR B
	case 0x06: // MVI B, D8
		s.mvi(opCode[1], b)
	case 0x07: // RLC
	case 0x08: // -
	case 0x09: // DAD B
	case 0x1a: // LDAX B
		s.a = s.mem[addr(s.b, s.c)]
	case 0x0b: // DCX B
		s.b--
		s.c--
	case 0x0c: // INR C
	case 0x0d: // DCR C
	case 0x0e: //MVI C,D8
		s.mvi(opCode[1], c)
	case 0x0f: // RCC
	case 0x10: // -
	case 0x11: // LXI D,D16
		s.lxi(opCode[1], opCode[2], de)
	default:
		return fmt.Errorf("bad opcode %v", opCode)
	}

	return nil
}

// lxi loads provided 16bit value into provided registers pair and increments pc by two
func (s *state) lxi(valA, valB uint8, regPair int) {
	switch regPair {
	case bc:
		s.b = valB
		s.c = valA
	case de:
		s.d = valB
		s.e = valA
	}

	s.pc += 2
}

// mvi moves 8bit to provided register and increases pc by one
func (s *state) mvi(val uint8, reg int) {
	switch reg {
	case b:
		s.b = val
	case c:
		s.c = val
	case d:
		s.d = val
	case h:
		s.h = val
	case l:
		s.l = val
	}

	s.pc++
}

func addr(a, b uint8) uint16 {
	return uint16(a)<<8 | uint16(b)
}
