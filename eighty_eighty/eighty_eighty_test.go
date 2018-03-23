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
}
