package main

func calculosMatematicos(a int, b int)(soma int, subtracao int){
	soma = a + b
	subtracao = a - b
	return
}
func main(){
	soma, subtracao := calculosMatematicos(10, 5)
	println(soma, subtracao)
}