package main

import "fmt"

func main() {
	/*
		Salto de línea (\n)
		La secuencia de escape "\n" se utiliza para representar un salto de línea en una cadena de caracteres. Por ejemplo:
		Ejemplo de secuencia de escape para salto de línea
	*/
	fmt.Println("Esta es una cadena con una nueva línea \ny una comilla doble \"")

	/*
		La salida del programa imprimirá:
		Esta es una cadena con una nueva línea
		y una comilla doble "
	*/

	/* Tabulación (\t)
	La secuencia de escape "\t" se utiliza para representar una tabulación en una cadena de caracteres*/
	fmt.Println("Este es un\ttexto con\ttabuladores")
	//Código de salida:
	//Este es un       texto con       tabuladores

}
