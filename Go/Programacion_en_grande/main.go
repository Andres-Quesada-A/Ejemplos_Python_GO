package main

import (
	// "Programacion_en_grande/Agregar"
	// "Programacion_en_grande/Multiplicacion"
	// "Programacion_en_grande/Resta"
	"fmt"

	"github.com/Andres-Quesada-A/Ejemplos_Python_Go/Programacion_en_grande/Agregar"
	"github.com/Andres-Quesada-A/Ejemplos_Python_Go/Programacion_en_grande/Multiplicacion"
	"github.com/Andres-Quesada-A/Ejemplos_Python_Go/Programacion_en_grande/Resta"
)

func main() {
	fmt.Println("3 + 5 =", Agregar.Agregar(3, 5))
	fmt.Println("8 - 2 =", Resta.Resta(8, 2))
	fmt.Println("4 * 6 =", Multiplicacion.Multiplicacion(4, 6))

}
