package div

import "testing"

func TestDiv(t *testing.T) {
	var v int64
	v = Div(8, 2)
	if v != 4 {
		t.Error("Expected 15, got ", v)
	}
}
