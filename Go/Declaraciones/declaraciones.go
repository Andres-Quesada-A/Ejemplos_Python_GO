package main

import (
	"fmt"
)

func main() {
	//Declaración de asignación de una constante
	const pi = 3.14159
	fmt.Println("Constante pi: ", pi)

	// Declaración y asignación de una variable
	var nombre_persona string = "Joshua"
	fmt.Println("Nombre: ", nombre_persona)

	// Declaración e inicialización corta
	//Nota: El compilador deduce que la variable num es del tipo
	num := 10
	fmt.Println("Declaración corta: ", num)

	// Declaración de una estructura
	type Persona struct {
		nombre string
		edad   int
	}
	//Creación de la instancia
	p := Persona{nombre: "Joshua", edad: 24}
	//Se accede a los campos de la estructura anterior
	fmt.Println("Nombre: ", p.nombre)
	fmt.Println("Edad: ", p.edad)

	//Declaración de asignación
	numero := 10
	//Declaración de condicional
	if numero > 0 {
		fmt.Println("El número es positivo")
	} else if numero < 0 {
		fmt.Println("El número es negativo")
	} else {
		fmt.Println("El número es cero")
	}

	//Declaración de bucle for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Llama a la función "suma" que se encuentra más abajo
	resultado := suma(2, 3)
	fmt.Println("Resultado función suma: ", resultado)

}

// Declaración de una función
func suma(a, b int) int {
	return a + b
}
