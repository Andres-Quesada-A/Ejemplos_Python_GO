package main

import (
	"fmt"
)

func main() {
	result := apply(double, 5) // pasa la función double como argumento a apply
	fmt.Println(result)        // imprime 10

	resultadoAdd := add(2, 6)
	fmt.Println(resultadoAdd) //Devuelve un 8 de la suma
}
func add(x, y int) int {
	return x + y
}

// Toam un entero y lo devuelve
func apply(f func(int) int, x int) int {
	return f(x) //Llama al entero y retorna el resulta
}

func double(x int) int {
	return x * 2
}
