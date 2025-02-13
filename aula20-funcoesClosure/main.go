package main

import "fmt"

func closure()func(){
	texto := "Dentro da função closure"
	funcao := func(){
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	fmt.Println("Aula 20 - Funções Closure")
	//Funcoes Closure sao funcoes que referenciam variaveis de um escopo externo

	texto := "Dentro da função main"
	fmt.Println(texto) //Dentro da função main

	funcaoNova := closure()
	funcaoNova() //Dentro da função closure
}