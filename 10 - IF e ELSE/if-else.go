package main

func main() {

	println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		println("Numero maior que 15")
	} else {
		println("Numero menor que 15 ou igual a 15")
	}

	// Exemplo de IF init
	if outroNumero := numero; outroNumero > 0 {
		println("Numero é maior que zero")
	}
}
