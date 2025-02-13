package main

import "fmt"

func main() {
	fmt.Println("Aula 10 - Map")

	//MAP
	//Map em Go é uma coleção de pares chave-valor

	//dentro do colchete é o tipo da chave e fora o tipo do valor
	usuario := map[string]string{
		"nome": "Fulano",
		"email": "mariano.silva@gmail.com",
	}

	fmt.Println(usuario)

	var usuario2 map[string]string = map[string]string{
		"nome": "Ciclano",
		"email": "email@ciclano.com",
	}

	usuario2["telefone"] = "999999999"

	fmt.Println(usuario2)

	delete(usuario2, "telefone")

	fmt.Println(usuario2)

	
}