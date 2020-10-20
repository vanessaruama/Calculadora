//Package pot realiza o cálculo da potência de um número
package pot

import (
	"math"
)

//Pow realiza o cálculo da potência de um número, multiplicando ele por ele mesmo na quantidade que foi definida no expoente
func Pow(a float64, b float64) float64 {
	return math.Pow(a, b)
}
