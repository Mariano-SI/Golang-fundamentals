package main

import "fmt"

//funcoes variaticas em go sao funcoes que aceitam um numero variavel de argumentos

func soma(numeros ...int)int{
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func showEnterpriseAndEmployees(enterprise string, employees ...string){
	for _, employee := range employees{
		fmt.Println(enterprise, employee)
	}
}

func main(){
	var resultado1 =  soma(1, 2, 3, 4, 5)

	showEnterpriseAndEmployees("Google", "Fulano", "Ciclano", "Beltrano")

	fmt.Println(resultado1)

}