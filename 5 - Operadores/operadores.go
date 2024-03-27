package main

func main() {
	//Aritmeticos
	soma := 1 + 2
	subtracao := 3 - 1
	divisao := 20 / 2
	multiplicacao := 2 * 2
	restoDivisao := 10 % 2

	println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	//Atribuicao

	var variavel string = "Teste"
	varivel2 := "Teste"
	println(variavel, varivel2)

	//Operadores Relacionais
	println(1 > 2)
	println(1 >= 2)
	println(1 == 2)
	println(1 <= 2)
	println(1 < 2)
	println(1 == 2)

	//Operadores Logicos
	verdadeiro, falso := true, false
	println(verdadeiro && falso)
	println(verdadeiro || falso)
	println(!verdadeiro)

	//Operadores Unarios
	numero := 10
	numero++
	numero += 17
	numero--
	numero *= 10
	numero /= 10
	numero %= 10
	println(numero)

}
