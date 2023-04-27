# TypeError
# El TypeError se produce cuando se intenta realizar una operación con un tipo de datos no válido. Por ejemplo:

# Suma de un entero y una cadena de caracteres
x = 10
y = "5"
z = x + y  # Genera un TypeError

# El mensaje de error generado sería algo así como:
# TypeError: unsupported operand type(s) for +: 'int' and 'str'




# IndexError
# El IndexError se produce cuando se intenta acceder a un índice que está fuera del rango de una secuencia. Por ejemplo:

# Acceso a un elemento inexistente de una lista
lista = [1, 2, 3]
elemento = lista[3]  # Genera un IndexError

# El mensaje de error generado sería algo así como:
# IndexError: list index out of range




# KeyError
# El KeyError se produce cuando se intenta acceder a una clave que no existe en un diccionario. Por ejemplo:
# Acceso a una clave inexistente en un diccionario
diccionario = {"nombre": "Juan", "edad": 25}
valor = diccionario["direccion"]  # Genera un KeyError

# El mensaje de error generado sería algo así como:
KeyError: 'direccion'






# ValueError
# El ValueError se produce cuando se intenta realizar una operación con un valor no válido. Por ejemplo:
# Conversión de una cadena de caracteres no numérica a entero
numero = int("abc")  # Genera un ValueError


# El mensaje de error generado sería algo así como:
# ValueError: invalid literal for int() with base 10: 'abc'





# SyntaxError
# El SyntaxError se produce cuando hay un error de sintaxis en el código. Por ejemplo:
# Falta de paréntesis en una llamada de función
resultado = print("Hola, mundo!"  # Genera un SyntaxError



# El mensaje de error generado sería algo así como:
# SyntaxError: unexpected EOF while parsing