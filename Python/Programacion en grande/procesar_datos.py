from leer_archivo import leer_archivo

def procesar_datos(nombre_archivo):
    # Leer el contenido del archivo usando la función leer_archivo
    contenido = leer_archivo(nombre_archivo)
    # Separar el contenido en líneas
    lineas = contenido.split('\n')
    # Calcular el número de líneas, palabras y caracteres
    num_lineas = len(lineas)
    num_palabras = sum(len(linea.split()) for linea in lineas)
    num_caracteres = sum(len(linea) for linea in lineas)
    # Devolver los resultados en una tupla
    return num_lineas, num_palabras, num_caracteres
