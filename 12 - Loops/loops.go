package main

import "fmt"

func main() {

	//Neste exemplo, o loop for itera de 0 a 4, imprimindo o valor de i a cada iteração.
	for i := 0; i < 5; i++ {
		println("Valor de i:", i)
	}

	//Este exemplo faz basicamente o mesmo que o exemplo anterior, mas a condição de continuação do loop (i < 5)
	//está diretamente na declaração do for, fazendo com que o loop funcione como um while.
	i2 := 0
	for i2 < 5 {
		println("Valor de i:", i2)
		i2++
	}

	//Neste exemplo, o loop range itera sobre cada elemento do slice numeros, fornecendo tanto o índice quanto o valor de cada elemento.
	numeros := []int{10, 20, 30, 40, 50}
	for i, valor := range numeros {
		println("Índice: %d, Valor: %d\n", i, valor)
	}

	//Neste exemplo, o loop range itera sobre cada par chave-valor do map frutas, imprimindo tanto a chave quanto o valor.
	frutas := map[string]string{
		"A": "Abacaxi",
		"B": "Banana",
		"C": "Cereja",
	}
	for chave, valor := range frutas {
		fmt.Printf("Chave: %s, Valor: %s\n", chave, valor)
	}

}
