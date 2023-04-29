# Cambio de tipo de variable
a = 5  # a es un entero
a = "cinco"  # a es ahora una cadena de texto





# Operaciones con diferentes tipos de datos
a = 5
b = "6"
c = a + int(b)
print(c)





# Uso de una variable como diferentes tipos de datos
a = "cinco"
print(a + " letras")  # "cinco letras"
a = 5
print(a + 7)  # 12






# Listas que contienen diferentes tipos de datos
lista = [1, "dos", 3.0, True]
for elemento in lista:
    print(type(elemento))