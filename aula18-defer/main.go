package main

import "fmt"


func funcao1(){
	fmt.Println("Executando a função 1")
}
func funcao2(){
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32)(bool, float32 ){
	defer fmt.Println("Média calculada. Resultado será retornado") //Por mais que isso esteja no inicio do bloco ele será a ultima coisa a ser executada nele, antes de retornar o valor
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	
	media := (n1 + n2) / 2

	return (media >= 6), media
}

func main() {
	fmt.Println("Aula 18 - Defer")

	//A clausula defer adia a execução de um determinado bloco de código ex: fechar um arquivo, fechar uma conexão com o banco de dados
	//DEFER = ADIAR 

	//Aqui executa em ordem primeiro logou a 1 e depois a 2
	funcao1()
	funcao2()

	//Aqui com o uso do defer a função 2 é executada antes da função 1 que so é executada no final
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}