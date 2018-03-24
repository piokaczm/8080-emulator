package eighty_eighty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestINR(t *testing.T) {
	t.Run("incrementing register", func(t *testing.T) {
		ee := New()
		ee.b = 0x00

		ee.inr(b)
		assert.Equal(t, uint8(0x01), ee.b, "increments register by one")
	})

	t.Run("setting zero flag", func(t *testing.T) {
		t.Run("when result equals zero", func(t *testing.T) {
			ee := New()
			ee.b = 0xff

			ee.inr(b)
			assert.Equal(t, uint8(0x01), ee.cc.z, "sets zero flag to one")
		})

		t.Run("when result bigger than zero", func(t *testing.T) {
			ee := New()
			ee.b = 0x00

			ee.inr(b)
			assert.Zero(t, ee.cc.z, "keeps zero flag equal to zero")
		})
	})

	t.Run("setting sign flag", func(t *testing.T) {
		t.Run("when 7th bit set", func(t *testing.T) {
			ee := New()
			ee.b = 0xfa

			ee.inr(b)
			assert.Equal(t, uint8(0x01), ee.cc.s, "sets sign flag to one")
		})

		t.Run("when 7th bit not set", func(t *testing.T) {
			ee := New()
			ee.b = 0x00

			ee.inr(b)
			assert.Zero(t, ee.cc.s, "keeps sign flag equal to zero")
		})
	})

	t.Run("setting parity flag", func(t *testing.T) {
		t.Run("when even 1s count", func(t *testing.T) {
			ee := New()
			ee.b = 0x08

			ee.inr(b)
			assert.Equal(t, uint8(0x01), ee.cc.p, "sets parity flag to one")
		})

		t.Run("when odd 1s count", func(t *testing.T) {
			ee := New()
			ee.b = 0x0c

			ee.inr(b)
			assert.Zero(t, ee.cc.p, "keeps parity flag equal to zero")
		})
	})
}

func TestDCR(t *testing.T) {
	t.Run("decrementing register", func(t *testing.T) {
		ee := New()
		ee.b = 0x02

		ee.dcr(b)
		assert.Equal(t, uint8(0x01), ee.b, "decrements register by one")
	})

	t.Run("setting zero flag", func(t *testing.T) {
		t.Run("when result equals zero", func(t *testing.T) {
			ee := New()
			ee.b = 0x01

			ee.dcr(b)
			assert.Equal(t, uint8(0x01), ee.cc.z, "sets zero flag to one")
		})

		t.Run("when result bigger than zero", func(t *testing.T) {
			ee := New()
			ee.b = 0x00

			ee.dcr(b)
			assert.Zero(t, ee.cc.z, "keeps zero flag equal to zero")
		})
	})

	t.Run("setting sign flag", func(t *testing.T) {
		t.Run("when 7th bit set", func(t *testing.T) {
			ee := New()
			ee.b = 0xfa

			ee.dcr(b)
			assert.Equal(t, uint8(0x01), ee.cc.s, "sets sign flag to one")
		})

		t.Run("when 7th bit not set", func(t *testing.T) {
			ee := New()
			ee.b = 0x01

			ee.dcr(b)
			assert.Zero(t, ee.cc.s, "keeps sign flag equal to zero")
		})
	})

	t.Run("setting parity flag", func(t *testing.T) {
		t.Run("when even 1s count", func(t *testing.T) {
			ee := New()
			ee.b = 0x0a

			ee.dcr(b)
			assert.Equal(t, uint8(0x01), ee.cc.p, "sets parity flag to one")
		})

		t.Run("when odd 1s count", func(t *testing.T) {
			ee := New()
			ee.b = 0x0e

			ee.dcr(b)
			assert.Zero(t, ee.cc.p, "keeps parity flag equal to zero")
		})
	})
}

func TestLXI(t *testing.T) {
	ee := New()

	pairs := []struct {
		registersPair  int
		firstRegister  *uint8
		secondRegister *uint8
	}{
		{bc, &ee.b, &ee.c},
		{de, &ee.d, &ee.e},
	}

	for _, testCase := range pairs {
		ee.pc = 0

		ee.lxi(0x01, 0x02, testCase.registersPair)
		assert.Equal(t, uint8(0x01), *testCase.secondRegister, "sets 2nd byte to pair's first register")
		assert.Equal(t, uint8(0x02), *testCase.firstRegister, "sets 3rd byte to pair's second register")
		assert.Equal(t, uint16(2), ee.pc, "increments pc by two additional steps")
	}
}

func TestEmulation(t *testing.T) {
	t.Run("when NOP", func(t *testing.T) {
		ee := New()
		ee.mem = []uint8{0x00}

		err := ee.Emulate()
		assert.Nil(t, err)
		assert.Equal(t, uint16(1), ee.pc, "increments pc by one")
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

func TestAddr(t *testing.T) {
	var a uint8 = 0x01
	var b uint8 = 0x02
	var expected uint16 = 0x0102

	assert.Equal(t, expected, addr(a, b))
}
