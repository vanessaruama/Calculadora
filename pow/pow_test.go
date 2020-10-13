package pow

import "testing"

func TestPow(t *testing.T) {
	var v float64
	v = Pow(2, 3)
	if v != 8 {
		t.Error("Expected 15, got ", v)
	}
}
