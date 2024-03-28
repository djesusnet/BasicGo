package main

func main() {
	println("Maps")

	usuario := map[string]string{
		"nome":      "Daniel",
		"sobrenome": "Jesus",
	}

	usuario2 := map[string]map[string]string{
		"aluno": {
			"nome": "Daniel Jesus",
		},
		"curso": {
			"nome":   "Analise de Sistemas",
			"campus": "Unip",
		},
	}

	println(usuario["nome"])
	println(usuario2["curso"]["campus"])
	delete(usuario2, "curso")

}
