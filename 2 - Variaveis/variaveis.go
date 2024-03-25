package main

func main() {

	var variavel1 string = "variável 1"
	variavel2 := "variável 2"

	var (
		variavel3 string = "variável 3"
		variavel4 string = "variável 4"
	)

	variavel5, variavel6 := "variável 5", "variável 6"

	const constante1 string = "Constante 1"

	println(variavel1)
	println(variavel2)
	println(variavel3, variavel4)
	println(variavel5, variavel6)
	println(constante1)
}
