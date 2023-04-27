def leer_archivo(nombre_archivo):
    # Abrir el archivo en modo lectura y asignar el archivo a la variable "f".
    print(nombre_archivo, "aqui")
    with open(nombre_archivo, 'r') as f:
        print("here")
    # Leer el contenido del archivo y asignarlo a la variable "contenido".
        contenido = f.read()
    # Devolver el contenido del archivo leído.
    print(contenido, "final")
    return contenido