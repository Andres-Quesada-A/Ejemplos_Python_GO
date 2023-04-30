package main

import "fmt"

func main() {

	//Declaración de un arreglo
	var Array [5]int // Arreglo de 5 enteros
	fmt.Println("Arreglo: ", Array)

	//Slices
	//Declaración de un slice de enteros
	Slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice: ", Slice)
	//Declaración de un slice de string. Agrega elementos al slice
	var Apellidos []string
	Apellidos = append(Apellidos, "Martinez")
	Apellidos = append(Apellidos, "Figueroa")
	fmt.Println("Slice de apellidos: ", Apellidos)
	//Modificación del valor de un elemento del slice
	Apellidos[0] = "Gutierrez"
	fmt.Println("Apellido modificado: ", Apellidos)

	//Creación de un mapa con datos string y numéricos enteros
	Mapa := map[string]int{"Mariana": 25, "Joshua": 24}
	fmt.Println("Mapa: ", Mapa)

	//Estructuras
	type Persona struct {
		nombre string
		edad   int
	}
	//Creación de la instancia
	p := Persona{nombre: "Joshua", edad: 24}
	//Se accede a los campos de la estructura anterior
	fmt.Println("Nombre: ", p.nombre)
	fmt.Println("Edad: ", p.edad)
	var miPersona Persona // una estructura de tipo Persona
	fmt.Println((miPersona))

	//Declaración de un puntero
	var num_Entero int = 42
	var Puntero *int = &num_Entero
	fmt.Println("Puntero", Puntero)

}
