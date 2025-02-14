package main

import "fmt"

type usuario struct {
	idade int
	nome string
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Aula 23 - Métodos")

	//A diferença entre função e método é que o método é uma função que está associada a um tipo por exemplo struct ou interface

	usuario1 := usuario{idade: 30, nome: "João"}
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.salvar()

	usuario2 := usuario{idade: 25, nome: "Maria"}
	usuario2.salvar()
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}