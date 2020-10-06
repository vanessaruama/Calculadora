package soma

import "testing"

func SomaTest(t *testing.T) {
	var v int64
	v = Soma(10, 5)
	if v != 15 {
		t.Error("Expected 15, got ", v)
	}
}
