package main

import (
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	println("Escrevendo do pacote main")
	auxiliar.Escrever()
	auxiliar.Escrever2()
	err := checkmail.ValidateFormat("ç$€§/az@gmail.com")
	println(err.Error())
}
