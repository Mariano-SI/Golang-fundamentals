package main

import "fmt"

type pessoa struct{
	nome string
	sobreNome string
	idade uint8
	altura uint8
}

type estudante struct{
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Aula 7 - Abordagem de Herança")

	p1 := pessoa{"Mariano", "Santos", 30, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia de Software", "USP"}
	fmt.Println(e1)

	var e2 estudante = estudante{p1, "Engenharia de Software", "USP"}
	fmt.Println(e2)

}