package sqr

import (
	"fmt"
	"testing"
)

func SqrTest(t *testing.T) {
	var v float64
	v = Sqrt(25)
	fmt.Println(v)
	if v != 10 {
		t.Error("Expected 5, got ", v)
	}
}
