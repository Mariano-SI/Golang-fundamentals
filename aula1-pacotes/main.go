package main

import(
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Olá, mundo!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("mariano.silva@gmail.com")
	fmt.Println(erro)
}