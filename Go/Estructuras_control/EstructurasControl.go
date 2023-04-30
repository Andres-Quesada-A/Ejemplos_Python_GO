package main

import "fmt"

func main() {
	//Estructura de control "if"
	//La estructura de control "if" permite ejecutar un bloque de código si se cumple una condición. Si la condición es falsa, el bloque no se ejecuta.
	edad := 18
	if edad > 18 { //# si la edad es mayor o igual a 18
		fmt.Println("Eres mayor de edad") //se ejecuta este bloque
	} else {
		fmt.Println("Eres menor de edad") //o de lo contrario, se ejecuta este bloque
	}

	//Estructura de control Switch
	diaSemana := "Jueves"
	switch diaSemana {
	case "Lunes":
		fmt.Println("Hoy es lunes")
	case "Martes":
		fmt.Println("Hoy es martes")
	case "Miércoles":
		fmt.Println("Hoy es miércoles")
	default:
		fmt.Println("Hoy es otro día de la semana")
	}

	//Estructura de control for
	frutas := []string{"manzana", "pera", "mango"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	//Estructura de control for-each (Range)
	//Permite ejecutar el bloque para cada elemento de la colección
	numeros := []int{1, 2, 3, 4, 5}
	for indice, valor := range numeros {
		fmt.Printf("El número en el índice %d es %d\n", indice, valor)
	}

	//Estructura de control While
	//Esta estructura en Go se implementa con un for
	contador := 0
	for contador < 10 {
		fmt.Println(contador)
		contador++
	}
}
