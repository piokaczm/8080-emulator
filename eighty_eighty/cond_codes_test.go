package eighty_eighty

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroFlagSetting(t *testing.T) {
	cc := &condCodes{}

	t.Run("when result equals zero", func(t *testing.T) {
		var result uint16 = 0x00
		cc.setZ(result)

		assert.Equal(t, uint8(1), cc.z, "sets zero flag to one")
	})

	t.Run("when result does not equal zero", func(t *testing.T) {
		var result uint16 = 0x01
		cc.z = 1
		cc.setZ(result)

		assert.Zero(t, cc.z, "sets zero flag to zero")
	})
}

func TestSignFlagSetting(t *testing.T) {
	cc := &condCodes{}

	t.Run("when 7th bit is set", func(t *testing.T) {
		var result uint16 = 0xff
		cc.setS(result)

		assert.Equal(t, uint8(1), cc.s, "sets sign flag to one")
	})

	t.Run("when 7th bit is not set", func(t *testing.T) {
		var result uint16 = 0x00
		cc.s = 1
		cc.setS(result)

		assert.Zero(t, cc.s, "sets sing flag to zero")
	})
}

func TestParityFlagSetting(t *testing.T) {
	cc := &condCodes{}

	t.Run("when 1s count is even", func(t *testing.T) {
		val, err := strconv.ParseInt("1100", 2, 8)
		if err != nil {
			log.Fatalf("cant parse provided binary: %s", err.Error())
		}
		var result uint16 = uint16(val)
		cc.setP(result)

		assert.Equal(t, uint8(1), cc.p, "sets parity flag to one")
	})

	t.Run("when 1s count is odd", func(t *testing.T) {
		val, err := strconv.ParseInt("1101", 2, 8)
		if err != nil {
			log.Fatalf("cant parse provided binary: %s", err.Error())
		}
		var result uint16 = uint16(val)
		cc.p = 1
		cc.setP(result)

		assert.Zero(t, cc.p, "sets zero flag to zero")
	})
}

func TestCarryBitFlagSetting(t *testing.T) {
	cc := &condCodes{}

	t.Run("when result produces carry bit", func(t *testing.T) {
		var result uint16 = 256
		cc.setCY(result)

		assert.Equal(t, uint8(1), cc.cy, "sets parity flag to one")
	})

	t.Run("when result does not produce carry bit", func(t *testing.T) {
		var result uint16 = 255
		cc.cy = 1
		cc.setCY(result)

		assert.Zero(t, cc.cy, "sets zero flag to zero")
	})
}
