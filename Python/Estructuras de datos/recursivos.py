# Estructuras de datos recursivas
# Árboles binarios
class Nodo:
    def __init__(self, valor):
        self.valor = valor
        self.izquierdo = None
        self.derecho = None
        
raiz = Nodo(1)
raiz.izquierdo = Nodo(2)
raiz.derecho = Nodo(3)
raiz.izquierdo.izquierdo = Nodo(4)
raiz.izquierdo.derecho = Nodo(5)






# Listas enlazadas
class Nodo:
    def __init__(self, valor):
        self.valor = valor
        self.siguiente = None
        
nodo1 = Nodo(1)
nodo2 = Nodo(2)
nodo3 = Nodo(3)

nodo1.siguiente = nodo2
nodo2.siguiente = nodo3