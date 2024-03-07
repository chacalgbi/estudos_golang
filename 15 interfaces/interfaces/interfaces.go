package main

import "fmt"

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return 3.14 * c.raio * c.raio
}

func imprimirArea(f forma) {
	fmt.Printf("A área é: %0.2f\n", f.area())
}

func main() {
	r := retangulo{10.71, 15.9}
	c := circulo{10.13}

	imprimirArea(r)
	imprimirArea(c)
}
