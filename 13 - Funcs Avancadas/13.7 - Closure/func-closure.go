package main

import "fmt"

func closure() func() {
	texto := "Entrou na função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Entrou na main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
