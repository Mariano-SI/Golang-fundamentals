package main

import "fmt"

func main() {
	fmt.Println("Aula 13 - Loops")

	//Go possui apenas o loop for que pode see utilizado de varias formas

	//For como while
	var number int = 0
	for number < 10 {
		fmt.Println(number)
		number++ //se não colocar essa linha o loop será infinito
	}

	//For tradicional iniciando a variavel no loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	} 

	//For tradicional utilizando a variavel de fora do loop
	var number2 int = 20
	for ; number2 < 30; number2++ {
		fmt.Println(number2)
	}

	//For com range
	//O range é utilizado para iterar sobre um slice, array, string, map ou channel
	//O range retorna dois valores, o indice e o valor
	var nomes = []string{"Fulano", "Ciclano", "Beltrano"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
		fmt.Println(nomes[indice])
	}

	const nome string = "Mariano"
	for indice, letra := range nome {
		fmt.Println(indice, string(letra)) //string(letra) converte o valor de letra para string já que por padrao ao iterar uma string o valor retornado é seu numero da tabela ASCII
	}

	var usuario = map[string]string{
		"nome": "Fulano",
		"email": "fulano@gmail.com",
	}

	for key, value := range usuario {
		fmt.Println(key, value)
	}
	
	
}