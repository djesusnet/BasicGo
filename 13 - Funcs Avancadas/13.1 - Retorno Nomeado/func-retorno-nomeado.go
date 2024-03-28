package main

func calculosMatematicos(numero1 int, numero2 int) (soma int, subtracao int) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 30)
	println(soma, subtracao)
}
