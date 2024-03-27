package main

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	println("Exemplo Structs")

	var user usuario
	user.idade = 34
	user.nome = "Daniel Jesus"

	println("Idade: ", user.idade, "Nome: ", user.nome)
}
