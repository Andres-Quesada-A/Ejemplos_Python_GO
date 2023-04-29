# Estructura de control "if"
# La estructura de control "if" permite ejecutar un bloque de código si se cumple una condición. Si la condición es falsa, el bloque no se ejecuta.
edad = 18

if edad >= 18:    # si la edad es mayor o igual a 18
    print("Eres mayor de edad")   # se ejecuta este bloque
else:
    print("Eres menor de edad")   # si no, se ejecuta este bloque





# Estructura de control "while"
# La estructura de control "while" permite repetir un bloque de código mientras se cumpla una condición. Si la condición deja de ser verdadera, se sale del bucle.
i = 1

while i <= 10:    # mientras i sea menor o igual a 10
    print(i)    # se imprime el valor de i
    i += 1    # se incrementa i en 1





# Estructura de control "for"
# La estructura de control "for" permite iterar sobre una secuencia de valores, como una lista, tupla o rango.
frutas = ["manzana", "pera", "mango"]

for fruta in frutas:    # para cada fruta en la lista "frutas"
    print(fruta)    # se imprime la fruta




# Estructura de control "break"
# La estructura de control "break" permite salir de un bucle antes de que se cumpla la condición.
i = 1

while i <= 10:    # mientras i sea menor o igual a 10
    if i == 5:    # si i es igual a 5
        break    # se sale del bucle
    print(i)    # se imprime el valor de i
    i += 1    # se incrementa i en 1





# Estructura de control "continue"
# La estructura de control "continue" permite saltar a la siguiente iteración de un bucle sin ejecutar el código restante del bloque.
frutas = ["manzana", "pera", "mango"]

for fruta in frutas:    # para cada fruta en la lista "frutas"
    if fruta == "pera":    # si la fruta es pera
        continue    # se salta a la siguiente iteración
    print(fruta)    # se imprime la