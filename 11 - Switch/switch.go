package main

func main() {
	diaDaSemana := 3 // Suponha que os dias da semana sejam numerados de 1 (Domingo) a 7 (Sábado)

	switch diaDaSemana {
	case 1:
		println("Domingo")
	case 2:
		println("Segunda-feira")
	case 3:
		println("Terça-feira")
	case 4:
		println("Quarta-feira")
	case 5:
		println("Quinta-feira")
	case 6:
		println("Sexta-feira")
	case 7:
		println("Sábado")
	default:
		println("Valor inválido para dia da semana")
	}
}
