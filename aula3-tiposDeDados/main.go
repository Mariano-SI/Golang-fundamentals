package main

import (
	"errors"
	"fmt"
)

func main() {
	//numeros inteiros:
	//int8, int16, int32, int64 e int o numero significa a quantidade de bits que suportam. Qundo usamos apenas int ele se baseia na arquitetura de bits do seu pc. no meu caso cria um int64

	var numero1 int64 = 1000000000000
	fmt.Println(numero1)

	//numeros reais:
	//float32 e float64

	const numeroReal1 float32 = 1000.43
	fmt.Println(numeroReal1)

	numeroReal2 := 1000.55
	fmt.Println(numeroReal2)

	//strings: 
	// string. OBS: em Go não existe char, apenas string

	var frase string = "Olá, mundo!"
	frase2 := "Olá, mundo 2!"

	fmt.Println(frase)
	fmt.Println(frase2)

	//valor default para variaveis nao inicializadas

	var texto string
	fmt.Println(texto) // string vazia

	var numero int
	fmt.Println(numero) // 0 independente de ser float ou int

	var booleano bool
	fmt.Println(booleano) // false

	var erro error
	fmt.Println(erro) // nil


	// booleanos
	// bool = true ou false

	var status bool = true
	fmt.Println(status)

	status2 := false
	fmt.Println(status2)

	//tipo error

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2) 
}