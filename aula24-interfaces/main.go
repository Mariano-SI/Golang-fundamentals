package main

import "fmt"

type retangulo struct {
	altura float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct{
	raio float64
}

func (c circulo) area() float64 {
	return c.raio * c.raio * 3.14
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}



func main() {
	fmt.Println("Aula 24 - Interfaces")
	//INterfaces em go descrevem apenas o comportamento de um tipo
	//Alem disso nao precisamos explicitamente dizer que um tipo implementa uma interface, basta que ele tenha os métodos necessários

	r := retangulo{altura: 10, largura: 15}
	escreverArea(r)
}