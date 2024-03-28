package main

import "reflect"

func main() {

	println("Arrays e Slices")

	var array [5]int
	array[0] = 1

	println(array[0])

	array2 := [5]string{"1", "2", "3", "4", "5"}

	println(array2[4])

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	println(array3[5])

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	println(slice)

	slice = append(slice, 19)

	println(slice[12])
	println(reflect.TypeOf(slice))
	println(reflect.TypeOf(array3))

	slice2 := array2[1:3]
	println(slice2[1])

	//arrays internos
	println("-----------Arrays internos---------")
	slice3 := make([]float32, 10, 15)
	println(len(slice3))
	println(cap(slice3))

}
