package main

import "fmt"

func main() {
	fmt.Println("Executando a funcao main")
}

//Executa primeiro que a main
func init() {
	fmt.Println("Executando a funcao init ")
}
