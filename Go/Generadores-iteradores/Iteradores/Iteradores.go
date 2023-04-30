package main

import (
	"fmt"
)

// Creación de la estructura
type Numeros struct {
	inicio, fin int
}

// Los iteradores se manejan por medio de las interfaces
func main() {
	numeros := Numeros{inicio: 1, fin: 9}
	for { //se usa el ciclo for para imprimir los números usando el método Next
		if num, ok := numeros.Next(); ok { //si Next retorna false quiere decir que ya no hay números por mostrar
			fmt.Println(num)
		} else {
			break
		}
	}
}

// Método que implementa la interfaz Iterador
func (n *Numeros) Next() (int, bool) {
	if n.inicio < n.fin {
		defer func() { n.inicio++ }()
		return n.inicio, true
	}
	return 0, false
}
