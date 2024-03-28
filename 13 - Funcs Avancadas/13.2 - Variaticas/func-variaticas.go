package main

// Func pode passar de 0 a N parametros do tipo int
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9)
	println(totalSoma)
}
