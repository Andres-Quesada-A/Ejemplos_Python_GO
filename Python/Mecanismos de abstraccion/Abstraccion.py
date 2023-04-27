# Funciones
# Las funciones son uno de los mecanismos de abstracción más básicos en Python. Permiten encapsular un conjunto de instrucciones en una sola entidad, que puede ser invocada desde cualquier parte del programa. Además, pueden aceptar argumentos y retornar valores.

# En este ejemplo, se presenta una función que calcula el área de un círculo dado su radio:
import math

def calcular_area_circulo(radio):
    area = math.pi * radio ** 2
    return area






# Clases
# Las clases son otro mecanismo de abstracción que permite modelar objetos complejos en Python. Permiten encapsular datos y comportamientos en una sola entidad, y pueden heredar propiedades y métodos de otras clases.
# En este ejemplo, se presenta una clase Persona que modela a una persona con nombre y edad:
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad
    
    def saludar(self):
        print(f"Hola, mi nombre es {self.nombre} y tengo {self.edad} años")






# Módulos
# Los módulos son otro mecanismo de abstracción que permite agrupar funciones y clases relacionadas en un mismo archivo. Esto facilita la organización y reutilización de código.
# En este ejemplo, se presenta un módulo operaciones que contiene funciones para sumar, restar, multiplicar y dividir dos números:
def sumar(a, b):
    return a + b

def restar(a, b):
    return a - b

def multiplicar(a, b):
    return a * b

def dividir(a, b):
    if b == 0:
        raise ValueError("No se puede dividir entre cero")
    return a / b