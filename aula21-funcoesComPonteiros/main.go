package main

import "fmt"

//funcao que manipula como copia
func invertendoValor(numero int) int{
	return numero * -1
}

func invertendoValorComPonteiro(numero *int){
	*numero = *numero * -1
}

func editaSlice(slice []int){
	slice[0] = 100
}

func main() {
	fmt.Println("Aula 21 - Funções com Ponteiros")

	numero := 20
	numeroInvertido := invertendoValor(numero)
	fmt.Println("Numero original: ", numero) //20
	fmt.Println("Numero invertido: ", numeroInvertido) // -20 //nesse caso minha função não altera o valor original da variavel numeropois ela foi passada como copia 

	novoNumero := 40
	invertendoValorComPonteiro(&novoNumero)
	fmt.Println("Novo numero: ", novoNumero) // -40 //nesse caso minha função altera o valor original da variavel novoNumero pois ela foi passada como ponteiro


	slice := []int{1, 2, 3}
	editaSlice(slice)
	fmt.Println("Slice editado: ", slice) // [100 2 3] //nesse caso minha função altera o valor original do slice pois ela foi passada como ponteiro
}