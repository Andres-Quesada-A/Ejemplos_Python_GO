# Funciones puras
# Definir una función que suma dos números
def sumar(a, b):
    return a + b

# Definir una función que multiplique dos números
def multiplicar(a, b):
    return a * b

# Definir una función que calcule el área de un círculo
def area_circulo(radio):
    pi = 3.14159
    return multiplicar(pi, multiplicar(radio, radio))

# Definir una función que calcule el perímetro de un círculo
def perimetro_circulo(radio):
    pi = 3.14159
    return multiplicar(pi, multiplicar(radio, 2))

# Imprimir el área y el perímetro de un círculo con radio 5
radio = 5
print("El área del círculo es:", area_circulo(radio))
print("El perímetro del círculo es:", perimetro_circulo(radio))








# Comprensión de listas
# Definir una lista de números
numeros = [1, 2, 3, 4, 5]

# Utilizar la comprensión de lista para obtener una lista de números pares
pares = [numero for numero in numeros if numero % 2 == 0]

# Imprimir la lista de números pares
print("Los números pares son:", pares)







# Funciones de orden superior
def aplicar_funcion(funcion, lista):
    return [funcion(elemento) for elemento in lista]

# Definir una función que multiplique un número por 2
def multiplicar_por_dos(numero):
    return numero * 2
