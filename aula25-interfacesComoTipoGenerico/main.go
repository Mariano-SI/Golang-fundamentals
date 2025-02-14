package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Aula 25 - Interfaces como tipo genérico")

	//Ineterfaces como tipo generico serve para que possamos passar qualquer tipo de dado para uma função
	//Isso é bom? depende, em alguns casos pode ser útil, em outros pode ser perigoso. Um exemplo disso é p fmt.Println que aceita qualquer tipo de dado

	generica("String")
	generica(1)
	generica(true)
	generica(1.5)
	generica(struct{}{})

}