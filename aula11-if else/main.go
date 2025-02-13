package main

import "fmt"

func main() {
	fmt.Println("Aula 11 - Estruturas de Controle")

	numero := 2

	if numero > 5 {
		fmt.Println("O numero é maior que 5")
	} else if numero == 5 {
		fmt.Println("O numero é igual a 5")
	}else{
		fmt.Println("O numero é menor que 5")
	}

	//if com inicialização de variavel
	if outroNumero := 10; outroNumero > 5 {
		fmt.Println("O outroNumero é maior que 5")
	} else {
		fmt.Println("O outroNumero é menor que 5")
	}

	//fmt.Println(outroNumero) //erro: undefined: outroNumero. Ela so existe dentro do bloco do if
}