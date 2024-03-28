package main

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
}

func main() {

	println("Funcoes Recursivas")

	posicao := uint(10)
	println(fibonacci(posicao))
}
