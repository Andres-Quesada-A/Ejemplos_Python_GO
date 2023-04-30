package main

//El polimorfismo en Go se logra con las interfaces.

import (
	"fmt"
)

// Definición de una interfaz que define un comportamiento común para una familia de tipos
type Animal interface {
	Sonido() string
}

// Definición de una estructura de tipo Perro
type Perro struct {
	Nombre string
}

// Implementación del método Sonido de la interfaz Animal para la estructura Perro
func (p Perro) Sonido() string {
	return "guau"
}

// Definición de una estructura de tipo Gato
type Gato struct {
	Nombre string
}

// Implementación del método Sonido de la interfaz Animal para la estructura Gato
func (g Gato) Sonido() string {
	return "miau"
}

// Función que recibe un parámetro de tipo Animal e imprime su sonido
func ImprimirSonido(animal Animal) {
	fmt.Println(animal.Sonido())
}

func main() {
	// Creación de una instancia de tipo Perro y otra de tipo Gato
	perro := Perro{Nombre: "Canelo"}
	gato := Gato{Nombre: "Neblina"}

	// Llamada a la función ImprimirSonido pasando como parámetros un Perro y un Gato
	ImprimirSonido(perro)
	ImprimirSonido(gato)
}
