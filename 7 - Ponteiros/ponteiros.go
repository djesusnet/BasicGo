package main

func main() {

	//PONTEIRO E UMA REFERENCIA DE MEMORIA
	var variavel int
	var ponteiro *int

	variavel = 100
	ponteiro = &variavel

	println(variavel, ponteiro)

	variavel = 150
	println(variavel, *ponteiro)

}
