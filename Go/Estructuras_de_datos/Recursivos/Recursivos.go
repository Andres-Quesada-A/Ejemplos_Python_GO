package main

func main() {
	//Arbol binario
	type Nodo struct {
		valor     int
		izquierdo *Nodo
		derecho   *Nodo
	}

	raiz := &Nodo{valor: 1}
	raiz.izquierdo = &Nodo{valor: 2}
	raiz.derecho = &Nodo{valor: 3}
	raiz.izquierdo.izquierdo = &Nodo{valor: 4}
	raiz.izquierdo.derecho = &Nodo{valor: 5}

}
