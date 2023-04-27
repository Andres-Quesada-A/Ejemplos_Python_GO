from procesar_datos import procesar_datos

nombre_archivo = 'archivo.txt'
num_lineas, num_palabras, num_caracteres = procesar_datos(nombre_archivo)

print(f"El archivo {nombre_archivo} tiene:")
print(f"{num_lineas} líneas")
print(f"{num_palabras} palabras")
print(f"{num_caracteres} caracteres")
