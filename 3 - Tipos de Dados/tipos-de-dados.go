package main

import "errors"

func main() {

	// Numeros

	var numero int16 = -100
	println(numero)

	var numero2 uint16 = 100
	println(numero2)

	//Alias para INT32
	var numero3 rune = 12345
	println(numero3)

	// Byte - Uint8
	var numero4 byte = 123
	println(numero4)

	numeroFloat3 := 123.45
	println(numeroFloat3)

	var numeroFloat float32 = 123.45
	println(numeroFloat)

	var numeroFloat2 float64 = 12355255225522.45
	println(numeroFloat2)

	//Strings
	var str string = "Teste"
	println(str)

	str2 := "Teste 2"
	println(str2)

	char := 'a'
	println(char)

	// Boleano

	var boleano bool = true
	println(boleano)

	//Errors
	var erro error = errors.New("Erro interno")
	println(erro.Error())

}
