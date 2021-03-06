package eighty_eighty

import (
	"fmt"
)

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
		s.stax(bc)
	case 0x03: // INX B
		s.inx(bc)
	case 0x04: // INR B; Z, S, P, AC
		s.inr(b)
	case 0x05: // DCR B
		s.dcr(b)
	case 0x06: // MVI B, D8
		s.mvi(opCode[1], b)
	case 0x07: // RLC
		s.rlc(a)
	case 0x08: // -
	case 0x09: // DAD B
		s.dad(bc)
	case 0x0a: // LDAX B
		s.ldax(bc)
	case 0x0b: // DCX B
		s.dcx(bc)
	case 0x0c: // INR C
	case 0x0d: // DCR C
	case 0x0e: // MVI C,D8
		s.mvi(opCode[1], c)
	case 0x0f: // RCC
	case 0x10: // -
	case 0x11: // LXI D,D16
		s.lxi(opCode[1], opCode[2], de)
	case 0x12: // STAX D
		s.stax(de)
	case 0x13: // INX D
		s.inx(de)
	case 0x14: // INR D
	case 0x15: // DCR D
	case 0x16: // MVI D, D8
		s.mvi(opCode[1], d)
	case 0x17: // RAL
	case 0x18: // -
	case 0x19: // DAD D
	case 0x1a: // LDAX D
		s.ldax(de)
	case 0x1b: // DCX D
		s.dcx(de)
	default:
		return fmt.Errorf("bad opcode %v", opCode)
	}

	return nil
}

// dad "double adds" a 16bit value located in the provided registers pair and stores the result in
// hl registers
func (s *state) dad(regPair int) {
	var result uint16

	switch regPair {
	case bc:
		bcValue := (uint16(s.b) << 8) + uint16(s.c)
		hlValue := (uint16(s.h) << 8) + uint16(s.l)
	}

	s.cc.setCY(result)
}

// rlc rotates accumulator left
func (s *state) rlc(reg int) {
	var result uint16
	switch reg {
	case a:
		shiftedA := uint16(s.a) << 1
		result = shiftedA | 1<<0
		s.a = uint8(result)
	}

	s.cc.setCY(result)
}

// inr increments value of single register and sets proper Z, S and P condition codes
func (s *state) inr(reg int) {
	var result uint16

	switch reg {
	case b:
		result = uint16(s.b) + 1
		s.b++
	}

	s.cc.setZ(result)
	s.cc.setS(result)
	s.cc.setP(result)
}

// dcr decrements value of single register and sets proper Z, S and P condition codes
func (s *state) dcr(reg int) {
	var result uint16

	switch reg {
	case b:
		result = uint16(s.b) - 1
		s.b--
	}

	s.cc.setZ(result)
	s.cc.setS(result)
	s.cc.setP(result)
}

// inx increments values stored in provided registers pair
func (s *state) inx(regPair int) {
	switch regPair {
	case bc:
		s.b++
		s.c++
	case de:
		s.d++
		s.e++
	}
}

// inx decrements values stored in provided registers pair
func (s *state) dcx(regPair int) {
	switch regPair {
	case bc:
		s.b--
		s.c--
	case de:
		s.d--
		s.e--
	}
}

// ldax loads value stored in memory address provided by registers pair in accumulator
func (s *state) ldax(regPair int) {
	var address uint16

	switch regPair {
	case bc:
		address = addr(s.b, s.c)
	case de:
		address = addr(s.d, s.e)
	}

	s.a = s.mem[address]
}

// stax stores data from acumulator to memory address provided from registers pair
func (s *state) stax(regPair int) {
	var address uint16

	switch regPair {
	case bc:
		address = addr(s.b, s.c)
	case de:
		address = addr(s.d, s.e)
	}

	s.mem[address] = s.a
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

// mvi moves 8bit value to provided register and increases pc by one
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
