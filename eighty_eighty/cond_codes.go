package eighty_eighty

import (
	"strconv"
)

type condCodes struct {
	z   uint8 // zero flag -> if result of operation is equal to 0 it's set to 1
	s   uint8 // sign flag -> set to 1 when bit 7 is set
	p   uint8 // parity flag -> check if results 1 bits count is even or odd
	cy  uint8 // carry bit flag -> check if result requires carrying a bit of a higher order
	ac  uint8 // auxiliary carry bit; leaving it for now
	pad uint8
}

func (c *condCodes) setZ(result uint16) {
	if result&0xff == 0 {
		c.z = 1
	} else {
		c.z = 0
	}
}

//  TODO: make sure it's not inverted
func (c *condCodes) setS(result uint16) {
	if result&(1<<7) == 0 {
		c.s = 0
	} else {
		c.s = 1
	}
}

func (c *condCodes) setP(result uint16) {
	var ones int

	for _, char := range strconv.FormatInt(int64(result), 2) {
		if string(char) == "1" {
			ones++
		}
	}

	if ones%2 == 0 {
		c.p = 1
	} else {
		c.p = 0
	}
}

func (c *condCodes) setCY(result uint16) {
	if result > 0xff {
		c.cy = 1
	} else {
		c.cy = 0
	}
}
