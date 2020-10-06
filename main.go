//não deixar executar se der firsterroro
//fazer só com a operação que for chamada na linha de comando
//operações que faltam
//testes para todas as operações
//ver se tem um jeito de rodar todos (testes)

package main

import (
	"fmt"
	"os"
	"strconv"
	"testes/div"
	"testes/multi"
	"testes/pow"
	"testes/soma"
	"testes/sqr"
	"testes/sub"
)

func main() {

	var firstnumber, secondnumber int64

	//var sqrnumber float64

	//argsWithProg := os.Args
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

	//fmt.Println(argsWithProg)
	//fmt.Println(arg, firstnumber, secondnumber)

	/*
		if arg == "soma" {
			result := soma.Soma(firstnumber, secondnumber)
			fmt.Printf("resultado da soma: %v \n", result)
			return
		} else if arg == "sub" {
			result := sub.Sub(firstnumber, secondnumber)
			fmt.Printf("resultado da subtração: %v \n", result)
			return
		} else if arg == "div" {
			result := div.Div(firstnumber, secondnumber)
			fmt.Printf("resultado da divisão: %v \n", result)
			return
		} else if arg == "multi" {
			result := multi.Multi(firstnumber, secondnumber)
			fmt.Printf("resultado da multiplicação: %v \n", result)
			return
		} else if arg == "pow" {
			result := pow.Pow(firstnumber)
			fmt.Printf("resultado da potência: %v \n", result)
			return
		} else {
			fmt.Printf("arg está inválido, opções válidas: soma, sub, div, multi, pow")
		}
	*/

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
	case "pow":
		result := pow.Pow(float64(firstnumber), float64(secondnumber)) //pacote para potencia
		fmt.Printf("resultado da potência: %v \n", result)
		return
	case "sqr":
		var floatnumber float64
		floatnumber = float64(firstnumber)
		result := sqr.Sqrt(floatnumber)
		fmt.Printf("resultado da raiz quadrada: %v \n", result)
		return
	default:
		fmt.Printf("o cálculo informado não é válido, opções válidas: soma, sub, div, multi, pow, sqr")
	}

}
