package main

import "fmt"

func main() {
	fmt.Println("Aula 9 - Ponteiros")

	//PONTEIROS
	//Variaveis do tipo valor em Go são passadas por valor ou seja a variavel declarada guarda o valor
	//Os tipos das variaveis valor em Go sao: int, string, float, bool , structs e arrays

	//teste com int
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	variavel1 = 20

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//teste com struct
	type Pessoa struct {
		Nome string
		Idade int
	}

	pessoa1 := Pessoa{"Fulano", 30}
	pessoa2 := pessoa1
	
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	pessoa1.Nome = "Ciclano"

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	//teste com array
	array1 := [5]int{1, 2, 3, 4, 5}
	array2 := array1

	fmt.Println(array1)
	fmt.Println(array2)

	array1[0] = 10

	fmt.Println(array1)
	fmt.Println(array2)


	//Variaveis do tipo referencia em Go sao passadas por referencia ou seja a variavel declarada guarda o endereco de memoria
	//Os tipos referência em Go são: slices, maps, channels, ponteiros

	//declarando um ponteiro explicito
	var valor1 int = 10
	var ponteiro1 *int = &valor1

	fmt.Println(valor1) //valor = 10
	fmt.Println(ponteiro1) //endereco de memoria

	*ponteiro1 = 20 //acessei o valor pelo ponteiro e lhe alterei

	fmt.Println(valor1) //valor = 20

	//exemplo com slices

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = 10

	fmt.Println(slice1)
	fmt.Println(slice2) //ambos foram alterados
	

}