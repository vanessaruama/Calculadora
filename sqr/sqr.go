package sqr

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	result := math.Sqrt(x)
	fmt.Println(result)
	return result
}
