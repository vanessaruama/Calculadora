package multi

import "testing"

func MultiTest(t *testing.T) {
	var v int64
	v = Multi(10, 5)
	if v != 15 {
		t.Error("Expected 15, got ", v)
	}
}
