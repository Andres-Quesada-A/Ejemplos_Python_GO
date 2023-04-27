import random

# Definir un generador que produce números aleatorios entre 0 y 99
def generar_datos():
    while True:
        yield random.randint(0, 99)

# Obtener datos del generador en tiempo real
generador = generar_datos()
for i in range(10):
    dato = next(generador)
    print("Dato generado:", dato)

# Usar los datos generados en un proceso
for i in range(10):
    dato = next(generador)
    resultado = dato * 2
    print("Resultado:", resultado)



# El código utiliza el módulo "random" para generar números aleatorios. 
# La función "generar_datos" define un generador que produce números 
# aleatorios entre 0 y 99 usando un bucle while y el comando "yield". 
# Luego, se crea una instancia del generador llamada "generador".

# El primer bucle for itera 10 veces, obteniendo el siguiente dato generado 
# del generador con la función "next" y lo imprime en la consola junto con un mensaje.

# El segundo bucle for también itera 10 veces y obtiene el siguiente dato 
# generado del generador usando la función "next", luego multiplica este 
# dato por 2 y lo imprime en la consola junto con un mensaje.


# Salida esperada
# Dato generado: 12
# Dato generado: 87
# Dato generado: 2
# Dato generado: 72
# Dato generado: 58
# Dato generado: 92
# Dato generado: 41
# Dato generado: 17
# Dato generado: 33
# Dato generado: 74
# Resultado: 58
# Resultado: 154
# Resultado: 84
# Resultado: 150
# Resultado: 146
# Resultado: 66
# Resultado: 44
# Resultado: 90
# Resultado: 92
# Resultado: 178