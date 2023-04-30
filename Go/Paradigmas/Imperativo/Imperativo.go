package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	suma := 0

	for i := 0; i < len(numeros); i++ {
		suma += numeros[i]
	}

	fmt.Println("La suma de los números es:", suma)

}
