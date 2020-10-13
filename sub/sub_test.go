package sub

import "testing"

func TestSub(t *testing.T) {
	var v int64
	v = Sub(11, 5)
	if v != 6 {
		t.Error("Expected 15, got ", v)
	}
}
