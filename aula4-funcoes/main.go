package main

import "fmt"

//declaracao de funcao comum
func somar(a int, b int) int {
	return a + b
}

//funcao com varios retornos
func calculosMatematicos(a, b int) (int, int, int, int) {
	soma := a + b
	subtracao := a - b
	divisao := a / b
	multiplicacao := a * b

	return soma, subtracao, divisao, multiplicacao
}

func main() {
	resultadoSoma := somar(2, 3)
	fmt.Println(resultadoSoma)

	//variavel recebendo funcao anonima
	var myFunc = func(txt string ){
		fmt.Println(txt)
	}

	myFunc("Mariano na funcao")

	//recebendo retorno de funcao com varios valores retornados
	soma, subtracao, divisao, multiplicacao := calculosMatematicos(10, 5)
	fmt.Println(soma, subtracao, divisao, multiplicacao)

	//ignorando retorno
	_, subtracao2, _, _ := calculosMatematicos(10, 5)
	fmt.Println(subtracao2)
}