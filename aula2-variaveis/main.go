package main

import "fmt"

func main() {
	//declaração com tipagem explicita
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)
	variavel1 = "Variável 1 modificada"

	//declaração com tipagem implicita
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	var(
		variavel3 string = "Variavel 3"
		variavel4 string
	)

	fmt.Println(variavel3, variavel4)

	//multipla declaraçao com inferencia
	variavel5, variavel6 := "Variável 5", 2

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)
}