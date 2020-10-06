package div

import "testing"

func DivTest(t *testing.T) {
	var v int64
	v = Div(10, 5)
	if v != 15 {
		t.Error("Expected 15, got ", v)
	}
}
