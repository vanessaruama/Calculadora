//Package fat realiza o cálculo fatorial de um número inteiro
package fat

//Fat realiza o cálculo fatorial de um número inteiro
func Fat(a int64) int64 {

	fact := 1

	for i := 1; int64(i) <= a; i++ {

		fact = fact * i

	}
	return int64(fact)
}
