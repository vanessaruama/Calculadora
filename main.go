//não deixar executar se der firsterroro
//fazer só com a operação que for chamada na linha de comando
//operações que faltam
//testes para todas as operações
//ver se tem um jeito de rodar todos (testes)
//adicionar operação fatorial

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/vanessaruama/Calculadora/div"
	"github.com/vanessaruama/Calculadora/fat"
	"github.com/vanessaruama/Calculadora/multi"
	"github.com/vanessaruama/Calculadora/pot"
	"github.com/vanessaruama/Calculadora/raiz"
	"github.com/vanessaruama/Calculadora/soma"
	"github.com/vanessaruama/Calculadora/sub"
)

func main() {

	var firstnumber, secondnumber int64

	arg := os.Args[1]
	firstnumber, firsterror := strconv.ParseInt(os.Args[2], 10, 64)   //converter str/int64
	secondnumber, seconderror := strconv.ParseInt(os.Args[3], 10, 64) //Arrumar array fora do balde

	if firsterror != nil {
		fmt.Println("por favor digite um número (int64) para o primeiro valor do cálculo")
		return
	}

	if seconderror != nil {
		fmt.Println("por favor digite um número (int64) para o segundo valor do cálculo")
		return
	}

	switch arg {
	case "soma":
		result := soma.Soma(firstnumber, secondnumber)
		fmt.Printf("resultado da soma: %v \n", result)
		return
	case "sub":
		result := sub.Sub(firstnumber, secondnumber)
		fmt.Printf("resultado da subtração: %v \n", result)
		return
	case "div":
		result := div.Div(firstnumber, secondnumber)
		fmt.Printf("resultado da divisão: %v \n", result)
		return
	case "multi":
		result := multi.Multi(firstnumber, secondnumber)
		fmt.Printf("resultado da multiplicação: %v \n", result)
		return
	case "pot":
		result := pot.Pow(float64(firstnumber), float64(secondnumber))
		fmt.Printf("resultado da potência: %v \n", result)
		return
	case "raiz":
		var floatnumber float64
		floatnumber = float64(firstnumber)
		result := raiz.Sqrt(floatnumber)
		fmt.Printf("resultado da raiz quadrada: %v \n", result)
		return
	case "fat":
		result := fat.Fat(firstnumber)
		fmt.Printf("resultado da potência: %v \n", result)
		return
	default:
		fmt.Printf("o cálculo informado não é válido, opções válidas: soma, sub, div, multi, pow, sqr")
	}

}
