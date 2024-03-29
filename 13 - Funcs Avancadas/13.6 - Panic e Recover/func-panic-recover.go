package main

import (
	"fmt"
)

func podeFalhar() {
	defer func() {
		// recover deve ser chamado dentro de uma função defer.
		// Se houve um panic, recover captura o valor do panic e continua a execução normalmente.
		if r := recover(); r != nil {
			fmt.Println("Recuperado de:", r)
		}
	}()

	fmt.Println("Executando podeFalhar")
	// Uma condição de erro fictícia.
	// Um panic interrompe a execução normal do programa e começa a propagar para cima na pilha de chamadas.
	panic("algo deu errado")
}

func main() {
	podeFalhar()
	fmt.Println("Execução continuou em main após a recuperação.")
}
