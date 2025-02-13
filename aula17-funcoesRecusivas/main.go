package main

import "fmt"

func fatorial(n uint)uint{
	if n == 0{
		return 1
	}
	return n * fatorial(n - 1)
}

func fibonacci(posicao int)int{
	if posicao <= 1{
		return posicao
	}
	return fibonacci(posicao - 1) + fibonacci(posicao - 2)
}


func main() {
	fmt.Println("Aula 17 - Funções Recursivas")

	fmt.Println("Fatorial de 5 é ", fatorial(5))
	
}