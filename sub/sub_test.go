package sub

import "testing"

func SubTest(t *testing.T) {
	var v int64
	v = Sub(10, 5)
	if v != 15 {
		t.Error("Expected 15, got ", v)
	}
}
