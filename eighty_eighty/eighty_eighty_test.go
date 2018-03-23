package eighty_eighty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmulation(t *testing.T) {
	t.Run("when NOP", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x00}

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})

	t.Run("when LXI B,D16", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x01, 0x00, 0x02}

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0x00), ee.c, "sets 2nd byte to state's c")
		assert.Equal(t, uint8(0x02), ee.b, "sets 3rd byte to state's b")
		assert.Equal(t, uint16(3), ee.pc, "increments pc by three")
	})

	t.Run("when STAX B", func(t *testing.T) {
		ee := New()
		ee.mem[0] = 0x02
		ee.a = 0x1f
		ee.b = 0x01
		ee.c = 0x02

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, ee.a, ee.mem[0x0102], "stores accumulator's value at memory address from registers b and c")
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})

	t.Run("when INX B", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x03}
		ee.b = 0x00
		ee.c = 0x01

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0x01), ee.b, "increments register b by one")
		assert.Equal(t, uint8(0x02), ee.c, "increments register c by one")
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})

	t.Run("when MVI B, D8", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x06, 0x1f}

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0x1f), ee.b, "sets register b to 2nd byte")
		assert.Equal(t, uint16(2), ee.pc, "increments pc by two")
	})

	t.Run("when LDAX B", func(t *testing.T) {
		ee := New()
		ee.b = 0x01
		ee.c = 0x02

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, ee.mem[0x0102], ee.a, "loads value from memory addr stored in registers b and c to accumulator")
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})

	t.Run("when DCX B", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x0b}
		ee.b = 0x00
		ee.c = 0x01

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0xff), ee.b, "decrements register b by one")
		assert.Equal(t, uint8(0x00), ee.c, "decrements register c by one")
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})

	t.Run("when MVI C, D8", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x0e, 0x1f}

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0x1f), ee.c, "sets register c to 2nd byte")
		assert.Equal(t, uint16(2), ee.pc, "increments pc by two")
	})

	t.Run("when LXI D,D16", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x11, 0x00, 0x02}

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0x00), ee.e, "sets 2nd byte to state's e")
		assert.Equal(t, uint8(0x02), ee.d, "sets 3rd byte to state's d")
		assert.Equal(t, uint16(3), ee.pc, "increments pc by three")
	})
}

func TestAddr(t *testing.T) {
	var a uint8 = 0x01
	var b uint8 = 0x02
	var expected uint16 = 0x0102

	assert.Equal(t, expected, addr(a, b))
}
