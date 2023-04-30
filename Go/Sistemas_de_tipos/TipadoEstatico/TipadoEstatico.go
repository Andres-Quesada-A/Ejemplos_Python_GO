package main

import "fmt"

func main() {
	//Es necesario declarar variables en su creación e imposible poder modifcar su tipo tiempo después
	var numero int = 10
	var mensaje string = "Hola mundo"

	// Por lo que a número no se puede asignar un valor de tipo cadena al haber sido declarado anteriormente como entero
	// numero = "20"

	fmt.Println(numero)
	fmt.Println(mensaje)
}
