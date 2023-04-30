package main

import "fmt"

func main() {
	//Slices
	// slice de números enteros
	numeros := []int{1, 2, 3, 4, 5}
	// slice de cadenas de texto
	frutas := []string{"manzana", "pera", "mango"}
	// slice de varios tipos de datos
	//Interface puede contener valores de cualquier tipo
	mezcla := []interface{}{1, "manzana", 3.14, true}
	fmt.Println("Slice de números: ", numeros)
	fmt.Println("Slice de frutas: ", frutas)
	fmt.Println("Slice mezclado: ", mezcla)

	//Arreglos
	// arreglo de dos números enteros
	coordenadas := [2]int{3, 4}
	// arreglo de tres cadenas de texto
	meses := [3]string{"enero", "febrero", "marzo"}
	fmt.Println("Coordenadas: ", coordenadas)
	fmt.Println("Meses: ", meses)

	//Mapas-Conjuntos
	// conjunto de tres cadenas de texto
	colores := map[string]bool{
		"rojo":  true,
		"verde": true,
		"azul":  true,
	}
	fmt.Println("Mapa -> Conjunto: ", colores)

	//Mapa. Equivalente a un diccionario
	persona := map[string]interface{}{
		"nombre":    "Joshua",
		"edad":      24,
		"profesion": "Profesor",
	}
	fmt.Println("Mapa -> diccionario: ", persona)
}
