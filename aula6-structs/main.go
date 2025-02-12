package main

import "fmt"

//structs são tipos de dados que permitem armazenar diferentes tipos de dados em um mesmo tipo
type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Aula 6 - Structs")

	//declarando uma variavel do tipo usuario
	var usuario1 usuario
	fmt.Println(usuario1) //assume os valores padrão dos tipos de dados
	usuario1.nome = "Mariano"
	usuario1.idade = 30
	fmt.Println(usuario1)


	//declarando uma variavel do tipo usuario com valores
	endereco1 := endereco{"Rua dos Bobos", 0}
	usuario2 := usuario{"Joaozinho", 20, endereco1}
	fmt.Println(usuario2)


	//declarando uma variavel do tipo usuario com valores parciais
	usuario3 := usuario{idade: 25}
	fmt.Println(usuario3)
}