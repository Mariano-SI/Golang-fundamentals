package main

import "fmt"

func init() {
	fmt.Println("Função init")
}

func main() {
	fmt.Println("Aula 22 - Função Init")

	//Função init é uma função que é executada antes da função main
	//Podemos ter mais de uma função init em varios arquivos do pacote
	//Casos comuns de uso: inicialização de variaveis, conexão com banco de dados, etc

}