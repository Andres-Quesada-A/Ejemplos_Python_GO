# Crear una lista
mi_lista = ['Hola', 'Mundo', '!']

# Crear un iterador a partir de la lista
mi_iterador = iter(mi_lista)

# Iterar sobre el iterador e imprimir cada elemento
for elemento in mi_iterador:
    print(elemento)






# Crear un diccionario
mi_diccionario = {'clave1': 'valor1', 'clave2': 'valor2', 'clave3': 'valor3'}

# Crear un iterador a partir del diccionario
mi_iterador = iter(mi_diccionario)

# Iterar sobre el iterador e imprimir cada clave y valor
for clave in mi_iterador:
    valor = mi_diccionario[clave]
    print(clave, valor)






# Crear una clase MiObjeto con un método __iter__() que retorna un iterador personalizado
class MiObjeto:
    def __init__(self):
        self.datos = [1, 2, 3, 4, 5]
    
    def __iter__(self):
        self.indice = 0
        return self
    
    def __next__(self):
        if self.indice < len(self.datos):
            resultado = self.datos[self.indice]
            self.indice += 1
            return resultado
        else:
            raise StopIteration

# Crear un objeto MiObjeto
mi_objeto = MiObjeto()

# Iterar sobre el objeto e imprimir cada elemento
for elemento in mi_objeto:
    print(elemento)






