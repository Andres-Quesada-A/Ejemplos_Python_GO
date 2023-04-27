# Excepción ValueError
# La excepción ValueError se produce cuando se intenta realizar una operación con un valor no válido. Por ejemplo, si intentamos convertir una cadena que no representa un número a un número entero, se producirá una excepción ValueError.
# Ejemplo de excepción ValueError
try:
    numero = int("abc")
except ValueError:
    print("La cadena no representa un número.")




# Excepción IndexError
# La excepción IndexError se produce cuando se intenta acceder a un índice que está fuera del rango de una secuencia. Por ejemplo, si intentamos acceder al cuarto elemento de una lista que solo tiene tres elementos, se producirá una excepción IndexError.
# Ejemplo de excepción IndexError
lista = [1, 2, 3]
try:
    elemento = lista[3]
except IndexError:
    print("El índice está fuera del rango.")



# Excepción TypeError
# La excepción TypeError se produce cuando se intenta realizar una operación con un tipo de datos no válido. Por ejemplo, si intentamos sumar un número entero y una cadena de caracteres, se producirá una excepción TypeError.
# Ejemplo de excepción TypeError
x = 10
y = "5"
try:
    z = x + y
except TypeError:
    print("La operación no es válida para los tipos de datos proporcionados.")





# Excepción KeyError
# La excepción KeyError se produce cuando se intenta acceder a una clave que no existe en un diccionario. Por ejemplo, si intentamos acceder a la clave "direccion" en un diccionario que solo tiene las claves "nombre" y "edad", se producirá una excepción KeyError.
# Ejemplo de excepción KeyError
diccionario = {"nombre": "Juan", "edad": 25}
try:
    valor = diccionario["direccion"]
except KeyError:
    print("La clave no existe en el diccionario.")