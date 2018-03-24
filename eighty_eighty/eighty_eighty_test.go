package eighty_eighty

import (
	"log"
	"strconv"
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
		ee.mem[0] = 0x0a
		ee.mem[0x0102] = 0xff

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

	t.Run("when STAX D", func(t *testing.T) {
		ee := New()
		ee.mem[0] = 0x12
		ee.a = 0x1f
		ee.d = 0x01
		ee.e = 0x02

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, ee.a, ee.mem[0x0102], "stores accumulator's value at memory address from registers d and e")
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})

	t.Run("when INX D", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x13}
		ee.d = 0x00
		ee.e = 0x01

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0x01), ee.d, "increments register d by one")
		assert.Equal(t, uint8(0x02), ee.e, "increments register e by one")
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})

	t.Run("when MVI D, D8", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x16, 0x1f}

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0x1f), ee.d, "sets register c to 2nd byte")
		assert.Equal(t, uint16(2), ee.pc, "increments pc by two")
	})

	t.Run("when LDAX D", func(t *testing.T) {
		ee := New()
		ee.d = 0x01
		ee.e = 0x02
		ee.mem[0] = 0x1a
		ee.mem[0x0102] = 0xff

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, ee.mem[0x0102], ee.a, "loads value from memory addr stored in registers b and c to accumulator")
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})

	t.Run("when DCX D", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x1b}
		ee.d = 0x00
		ee.e = 0x01

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint8(0xff), ee.d, "decrements register d by one")
		assert.Equal(t, uint8(0x00), ee.e, "decrements register e by one")
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
	})
}

func TestConditionCodesChecks(t *testing.T) {
	t.Run("zero check", func(t *testing.T) {
		ee := New()

		t.Run("when result equals zero", func(t *testing.T) {
			var result uint16 = 0x00
			ee.cc.setZ(result)

			assert.Equal(t, uint8(1), ee.cc.z, "sets zero flag to one")
		})

		t.Run("when result does not equal zero", func(t *testing.T) {
			var result uint16 = 0x01
			ee.cc.z = 1
			ee.cc.setZ(result)

			assert.Zero(t, ee.cc.z, "sets zero flag to zero")
		})
	})

	t.Run("sign check", func(t *testing.T) {
		ee := New()

		t.Run("when 7th bit is set", func(t *testing.T) {
			var result uint16 = 0xff
			ee.cc.setS(result)

			assert.Equal(t, uint8(1), ee.cc.s, "sets sign flag to one")
		})

		t.Run("when 7th bit is not set", func(t *testing.T) {
			var result uint16 = 0x00
			ee.cc.s = 1
			ee.cc.setS(result)

			assert.Zero(t, ee.cc.s, "sets sing flag to zero")
		})
	})

	t.Run("parity check", func(t *testing.T) {
		ee := New()

		t.Run("when 1s count is even", func(t *testing.T) {
			val, err := strconv.ParseInt("1100", 2, 8)
			if err != nil {
				log.Fatalf("cant parse provided binary: %s", err.Error())
			}
			var result uint16 = uint16(val)
			ee.cc.setP(result)

			assert.Equal(t, uint8(1), ee.cc.p, "sets parity flag to one")
		})

		t.Run("when 1s count is odd", func(t *testing.T) {
			val, err := strconv.ParseInt("1101", 2, 8)
			if err != nil {
				log.Fatalf("cant parse provided binary: %s", err.Error())
			}
			var result uint16 = uint16(val)
			ee.cc.p = 1
			ee.cc.setP(result)

			assert.Zero(t, ee.cc.p, "sets zero flag to zero")
		})
	})
}

func TestAddr(t *testing.T) {
	var a uint8 = 0x01
	var b uint8 = 0x02
	var expected uint16 = 0x0102

	assert.Equal(t, expected, addr(a, b))
}
