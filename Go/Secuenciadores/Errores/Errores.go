package main

import "fmt"

func main() {
	// Definir un slice de enteros
	numeros := []int{1, 2, 3, 4, 5}

	//SyntaxError
	//Este se produce cuando hay errores en el código. En este caso, faltaría el paréntesis final de la llamada a la variable
	cadena_string :="Ejemplo de error"
	fmt.Println(cadena_string

	// Iterar sobre los números
	for i := 0; i <= len(numeros); i++ {
		fmt.Println(numeros[i])
	}
	//El mensaje de error generado sería:
	//panic: runtime error: index out of range [5] with length 5
	/*Esto se debe a que la condición debería de ser i < len(numeros), ya que el índice
	del slice da inicio en el 0 y la longitud de esta es de 5, por lo tanto, el índice máximo válido es 4
	Asimismo, si se trata de imprimir el último número, es decir, el 5, dará error de índice fuera de rango
	*/
}
