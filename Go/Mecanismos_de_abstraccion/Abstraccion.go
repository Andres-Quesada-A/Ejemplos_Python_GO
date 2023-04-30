package main

import (
	"fmt"
)

func main() {
	//Llama a la función "suma"
	resultado := suma(2, 3)
	fmt.Println("Resultado de la función suma: ", resultado)
	//Llama a la función "resta"
	resultado_Resta := resta(2, 3)
	fmt.Println("Resultado de la función resta: ", resultado_Resta)
	//Llama a la función "multiplicación"
	resultado_multiplicacion := multiplicacion(2, 3)
	fmt.Println("Resultado de la función multiplicación: ", resultado_multiplicacion)
	//Llama a la función "suma"

	//Estructura -> Clase
	//En este ejemplo, se presenta una clase Animal que modela a un animal con nombre, edad, y si es mascota:
	type Animal struct {
		Nombre    string
		Edad      int
		EsMascota bool
	}
	animal := Animal{Nombre: "Neblina", Edad: 4, EsMascota: true}
	fmt.Println("Nombre: ", animal.Nombre)
	fmt.Println("Edad: ", animal.Edad)
	fmt.Println("¿Es mascota?: ", animal.EsMascota)

}

// Módulos de operaciones básicas con paso de parámetros
func suma(a int, b int) int {
	return a + b
}

func resta(a int, b int) int {
	return a - b
}

func multiplicacion(a int, b int) int {
	return a * b
}
