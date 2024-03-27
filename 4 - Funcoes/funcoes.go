package main

func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

func subtrair(numero1 int8, numero2 int8) int8 {
	return numero1 - numero2
}

func calculosMatematicos(numero1, numero2 int8) (int8, int8) {
	soma := somar(numero1, numero2)
	sub := subtrair(numero1, numero2)
	return soma, sub
}

var funcao = func(txt string) string {
	return txt
}

func main() {

	soma := somar(2, 2)
	sub := subtrair(3, 1)
	txt := funcao("Teste")
	resultadosoma, reultadoSubtracao := calculosMatematicos(30, 10)
	resulsoma, _ := calculosMatematicos(30, 10)

	println(soma)
	println(sub)
	println(txt)
	println(resultadosoma)
	println(reultadoSubtracao)
	println(resulsoma)

}
