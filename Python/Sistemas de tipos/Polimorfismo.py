# El polimorfismo por Duck Typing en Python se refiere a la capacidad de los objetos 
# de diferentes clases para responder a un mismo método o atributo. A continuación, 
# te presento algunos ejemplos de cómo se utiliza el polimorfismo por Duck Typing en Python:
# Mismo método en diferentes objetos
class Perro:
    def hablar(self):
        print("Guau!")
        
class Gato:
    def hablar(self):
        print("Miau!")
        
class Vaca:
    def hablar(self):
        print("Muu!")
        
animales = [Perro(), Gato(), Vaca()]

for animal in animales:
    animal.hablar()








# Mismo atributo en diferentes objetos
class Coche:
    def __init__(self, marca):
        self.marca = marca
        
class Moto:
    def __init__(self, marca):
        self.marca = marca
        
class Bicicleta:
    def __init__(self, marca):
        self.marca = marca
        
vehiculos = [Coche("Ford"), Moto("Honda"), Bicicleta("Trek")]

for vehiculo in vehiculos:
    print(vehiculo.marca)
















# Uso de operadores en diferentes objetos
class Entero:
    def __init__(self, valor):
        self.valor = valor
        
    def __add__(self, otro):
        return Entero(self.valor + otro.valor)
        
class Flotante:
    def __init__(self, valor):
        self.valor = valor
        
    def __add__(self, otro):
        return Flotante(self.valor + otro.valor)
        
a = Entero(5)
b = Flotante(3.5)
c = a + b
print(c.valor)