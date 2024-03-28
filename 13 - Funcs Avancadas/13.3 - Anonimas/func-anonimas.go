package main

import "fmt"

func main() {

	func(texto string) string {
		return fmt.Sprintf("Recebido => %s", texto)
	}("Passando parametros")

}
