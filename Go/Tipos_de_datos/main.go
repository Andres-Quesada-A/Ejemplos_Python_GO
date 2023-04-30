package main

import (
	"fmt"
)

func main() {

	// Declaración de un número entero
	var numero_entero int = 5
	fmt.Println("Número entero: ", numero_entero) //Output: 5

	//Declaración de un número flotante
	var numero_flotante float32 = 3.14
	fmt.Println("Número flotante: ", numero_flotante) //Output: 3.14

	//Declaración de un dato booleano
	var dato_booleano bool = true
	fmt.Println("Dato booleano: ", dato_booleano) //Output: true

	//Declaración de una cadena de carácteres
	var cadena_caracteres string = "Hola Mundo"
	fmt.Println("Cadena de caracteres: ", cadena_caracteres) //Output: Hola Mundo

	//Datos compuestos

	//Declaración de un dato numérico sin signo
	var dato_entero_sin_signo uint = 50
	fmt.Println("Dato entero sin signo: ", dato_entero_sin_signo) //Output: 50

	//Declaración de un dato numérico entero 8Bits
	var dato_entero_8bits int8 = -128
	fmt.Println("Dato entero 8bits: ", dato_entero_8bits) //Output: -128

	//Declaración de un dato numérico entero 16 Bits
	var dato_entero_16bits int16 = 32767
	fmt.Println("Dato entero 16 bits: ", dato_entero_16bits) //Output: 32767

	//Declaración de un dato numérico entero 32 Bits
	var dato_entero_32bits int32 = 2147483647
	fmt.Println("Dato entero 32 bits: ", dato_entero_32bits) //Output: 2147483647

	//Declaración de un dato numérico entero 64 Bits
	var dato_entero_64bits int64 = 9223372036854775807
	fmt.Println("Dato entero 64 bits: ", dato_entero_64bits) //Output: 9223372036854775807

	//Declaración de un dato numérico entero 32 Bits sin signo
	var dato_entero_32bits_sin_signo uint32 = 4294967295
	fmt.Println("Dato entero 32 bits sin signo: ", dato_entero_32bits_sin_signo) //Output: 4294967295

	//Declaración de un dato numérico entero 64 Bits sin signo
	var dato_entero_64bits_sin_signo uint64 = 18446744073709551615
	fmt.Println("Dato entero 64 bits sin signo: ", dato_entero_64bits_sin_signo) //Output: 18446744073709551615

	//Declaración de un dato numérico complejo
	var dato_numero_complejo complex64 = 1 + 2i
	fmt.Println("Dato numérico complejo: ", dato_numero_complejo) //Output: (1 + 2i)

	//Declaración de una matriz
	var dato_matriz [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Matriz: ", dato_matriz) //Output: [[1, 2, 3], [4, 5, 6}, [7, 8, 9]]

	//Declaración de un slice
	var dato_slice []int = []int{1, 2, 3}
	fmt.Println("Dato Slice: ", dato_slice) //Output: [1, 2, 3]

	//Declaración de un mapa
	var dato_mapa map[string]int = map[string]int{"uno": 1, "dos": 2, "tres": 3}
	fmt.Println("Mapa: ", dato_mapa) //Output: map[dos:2 tres:3 uno:1]

}
