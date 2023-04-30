package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//La función panic indica un error, el mensaje de error puede ser adaptado a como quiera el programador
	nombre := "Joshua Mendoza"
	if !strings.Contains(nombre, " ") {
		panic("El nombre debe contener un espacio")
	}
	fmt.Println("Hola,", nombre)

	//Función Atoi convierte la cadena en un número entero
	//Si la conversión llega a fallar, retornará un error generando la excepción con la función panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("La cadena no representa un número.")
		}
	}()

	numero, err := strconv.Atoi("abc")
	if err != nil {
		panic(err)
	}
	fmt.Println(numero)
}
