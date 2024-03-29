package main

import "fmt"

func generica(i interface{}) {
	fmt.Println(i)
}

func main() {
	generica("teste")
	generica(10)
	generica(true)
}
