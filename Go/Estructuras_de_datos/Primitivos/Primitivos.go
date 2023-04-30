package main

import (
	"fmt"
)

func main() {
	numero_entero := 42                         //Número entero
	numero_flotante := 3.14                     //Número flotante
	var dato_numero_complejo complex64 = 2 + 2i //Número complejo

	//Cadenas de texto
	cadena := "Hola mundo"
	caracter := "a"

	//Booleano
	booleano_verdadero := true
	booleano_falso := false

	//Nil
	//Indica que una variable no tiene un valor asignado
	var variable_nil chan int = nil

	fmt.Println("Dato numérico complejo: ", dato_numero_complejo)
	fmt.Println("Número entero: ", numero_entero)
	fmt.Println("Numero flotante: ", numero_flotante)
	fmt.Println("Cadena de texto: ", cadena)
	fmt.Println("Cadena de caracter: ", caracter)
	fmt.Println("Booleano verdadero: ", booleano_verdadero)
	fmt.Println("Booleano falso: ", booleano_falso)
	fmt.Println("Variable tipo nil: ", variable_nil)
}
