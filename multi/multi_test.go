package multi

import "testing"

func TestMulti(t *testing.T) {
	var v int64
	v = Multi(10, 5)
	if v != 50 {
		t.Error("Expected 15, got ", v)
	}
}
