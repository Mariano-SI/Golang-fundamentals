package main

import "fmt"

func main() {
	fmt.Println("Aula 8 - Arrays e Slices")
	//ARRAYS
	//Arrays em Go tem alocação estatica de memoria
	fmt.Println("Arrays-------------------")
	var array1 [5]int 
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	//nao é possivel declarar um array como const porque const so aceita tipos primitivos ex: int, string, float
	//const array2 [5]int = [5]int{1, 2, 3, 4, 5}

	var array2 = [3]string{"teste1", "teste2"}
	fmt.Println(array2)

	array3 := [2]string{"inferencia1", "inferencia2"}
	fmt.Println(array3)
	fmt.Println("-----------------------")

	//SLICES
	//Slices em Go tem alocação dinamica de memoria
	fmt.Println("Slices-------------------")
	slice1 := []int{1, 2, 3, 4, 5}
	//slice1[5] = 5 //gera um erro panic: runtime error: index out of range [5
	slice1 = append(slice1, 6) //adiciona um elemento ao slice
	fmt.Println(slice1)

	var slice2 []int
	slice2 = append(slice2, 1)
	fmt.Println(slice2)

}