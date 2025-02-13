package main

import "fmt"

func recuperarExecucao(){
	if r := recover(); r != nil {
		fmt.Println("Recuperamos a execução com sucesso")
	}
}

func alunoEstaProvado(n1, n2 float32) bool {
	defer recuperarExecucao() 
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	}else if media < 6 {
		return false
	}
	panic("A MÉDIA É EXATAMENTE 6!!!") //Panic é um erro fatal, ele para a execução do programa. Antes disso ele apenas chama as defer functions
}

func main() {
	fmt.Println("Aula 19 - Panic e Recover")
	fmt.Println(alunoEstaProvado(6, 6))
	fmt.Println("Pós execução")
}