package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func NuevaPersona(nombre string, edad int) *Persona {
	p := Persona{nombre: nombre, edad: edad}
	return &p
}

func (p *Persona) ObtenerNombre() string {
	return p.nombre
}

func (p *Persona) ObtenerEdad() int {
	return p.edad
}

func main() {
	persona := NuevaPersona("Joshu	", 24)
	fmt.Println("Nombre: ", persona.ObtenerNombre())
	fmt.Println("Edad: ", persona.ObtenerEdad())
}
