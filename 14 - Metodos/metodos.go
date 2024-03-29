package main

import "fmt"

type usuario struct {
	nome      string
	sobreNome string
}

func (u usuario) salvar() {
	fmt.Println("Chamando o método salvar no contexto de usuario.")
}

func main() {
	usuario1 := usuario{"Daniel", "Jesus"}
	fmt.Println(usuario1)
	usuario1.salvar()
}
