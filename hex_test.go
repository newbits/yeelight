package yeelight

import (
	"testing"
)

// TestToRgbInt passes valid and invalid values to the methods and make assertions.
func TestToRgbInt(t *testing.T) {
	_, err := Hex{Value: "invalid format"}.ToRgbInt()
	if err == nil {
		t.Errorf("Invalid format was passed, but it returned no error.")
	}

	value, err := Hex{Value: "#112233"}.ToRgbInt()
	if err != nil {
		t.Errorf("Valid format was passed, but error occured.")
	}

	// 256 * 256 * 17 (R) + 256 * 34 (G) + 51 (B) = 1122867
	if value != 1122867 {
		t.Errorf("Calculated int value is incorrect")
	}
}
