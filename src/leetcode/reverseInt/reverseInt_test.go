package reverseInt

import "testing"

//整数反转
func TestReverseInt(t *testing.T) {
	if r := ReverseInt(123); r != 321 {
		t.Errorf("123 reverse be 321,but %d go", r)
	}

	if r := ReverseInt(-123); r != -321 {
		t.Errorf("-123 reverse be -321,but %d go", r)
	}

	if r := ReverseInt(120); r != 21 {
		t.Errorf("120 reverse be 21,but %d go", r)
	}
}
