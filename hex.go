package yeelight

import (
	"errors"
	"image/color"
)

type Hex struct {
	Value string
}

// ToRgbInt will convert the value in string to a int value
func (h Hex) ToRgbInt() (value int, err error) {
	c := color.RGBA{}
	c.A = 0xff

	if h.Value[0] != '#' {
		return value, errors.New("missing '#' at the beginning of the string")
	}

	switch len(h.Value) {
	case 7:
		c.R = hexToByte(h.Value[1])<<4 + hexToByte(h.Value[2])
		c.G = hexToByte(h.Value[3])<<4 + hexToByte(h.Value[4])
		c.B = hexToByte(h.Value[5])<<4 + hexToByte(h.Value[6])
	case 4:
		c.R = hexToByte(h.Value[1]) * 17
		c.G = hexToByte(h.Value[2]) * 17
		c.B = hexToByte(h.Value[3]) * 17
	default:
		err = errors.New("incorrect value")
	}

	value = 256*256*int(c.R) + 256*int(c.G) + int(c.B)

	return
}

func hexToByte(b byte) byte {
	switch {
	case b >= '0' && b <= '9':
		return b - '0'
	case b >= 'a' && b <= 'f':
		return b - 'a' + 10
	case b >= 'A' && b <= 'F':
		return b - 'A' + 10
	}

	return 0
}
