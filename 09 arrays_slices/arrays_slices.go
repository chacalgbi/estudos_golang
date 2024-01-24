package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("---------- Arrays (tem tamanhos fixos e tipos fixos)----------")
	var array1 [5]int
	fmt.Println(array1)
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // O compilador conta o tamanho do array
	fmt.Println(array3)
	fmt.Println(reflect.TypeOf(array3))

	fmt.Println("\n---------- Slices ----------")

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19} // Slice não tem tamanho fixo e usa chaves
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 20) // Adiciona um elemento ao slice
	fmt.Println(slice)

	slice2 := array2[1:3] // Slice de um array
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(slice2) // Slice é uma referência de memória, então se alterar o array, altera o slice

	fmt.Println("\n---------- Arrays Internos ----------")

	slice3 := make([]float32, 10, 11) // Cria um slice com 10 posições e um array interno com 11 posições
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho do slice
	fmt.Println(cap(slice3)) // Tamanho do array interno

	slice3 = append(slice3, 5) // Adiciona um elemento ao slice
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho do slice
	fmt.Println(cap(slice3)) // Tamanho do array interno
}
