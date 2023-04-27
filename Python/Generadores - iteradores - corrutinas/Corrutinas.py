# Corrutina simple
import asyncio

# Definir una corrutina que recibe un mensaje y lo imprime
async def imprimir_mensaje(mensaje):
    print("Mensaje recibido:", mensaje)

# Ejecutar la corrutina
asyncio.run(imprimir_mensaje("Hola, mundo!"))





# Corrutina con comunicación asincrónica
# Definir una corrutina que recibe un mensaje, espera un segundo y devuelve una respuesta
async def responder_mensaje(mensaje):
    print("Mensaje recibido:", mensaje)
    await asyncio.sleep(1)
    return "Respuesta al mensaje: " + mensaje

# Ejecutar la corrutina y obtener la respuesta
async def obtener_respuesta(mensaje):
    respuesta = await responder_mensaje(mensaje)
    print(respuesta)

asyncio.run(obtener_respuesta("Hola, mundo!"))
