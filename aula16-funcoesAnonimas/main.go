package main

func main(){
	//funcao anonima
	func(){
		println("Função anonima")
	}()

	//funcao anonima com parametro
	func(msg string){
		println(msg)
	}("Função anonima com parametro")

	//funcao anonima atribuida a uma variavel
	f := func(msg string){
		println(msg)
	}
	f("Função anonima atribuida a uma variavel")
}