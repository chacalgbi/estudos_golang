package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(10)
	generica(10.5)
	generica(true)
	generica(struct{}{})
	generica([]int{1, 2, 3})
	generica(map[string]int{"um": 1, "dois": 2})
	generica(func() {})
	generica(nil)
}
