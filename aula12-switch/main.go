package main

import "fmt"

//Switch avaliando valor de uma variavel
func diaDaSemana(number int)string {
	switch number {
		case 1:
			return "Domingo"
		case 2:
			return "Segunda"
		case 3:
			return "Terça"
		case 4:
			return "Quarta"
		case 5:
			return "Quinta"
		case 6:
			return "Sexta"
		case 7:
			return "Sábado"
		default:
			return "Número inválido"
	}
}

//switch que avalia condicoes sobre uma ou varias variaveis
func diaDaSemana2(number int)string{
	switch{
		case number == 1:
			return "Domingo"
		case number == 2:
			return "Segunda"
		case number == 3:
			return "Terça"
		case number == 4:
			return "Quarta"
		case number == 5:
			return "Quinta"
		case number == 6:
			return "Sexta"
		case number == 7:
			return "Sábado"
		default:
			return "Número inválido"
	}
}

func main() {
	fmt.Println("Aula 12 - Switch")

	fmt.Println("Digite um número de 1 a 7")
	var numero int  
	fmt.Scanln(&numero)

	fmt.Println("O dia da semana é ", diaDaSemana(numero))

	
}