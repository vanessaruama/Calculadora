package sqr

import (
	"testing"
)

func TestSqr(t *testing.T) {
	var v float64
	v = Sqrt(25)
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}
