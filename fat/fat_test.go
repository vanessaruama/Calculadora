package fat

import "testing"

func TestFat(t *testing.T) {
	var v int64
	v = Fat(5)
	if v != 120 {
		t.Error("Expected 15, got ", v)
	}
}
go 