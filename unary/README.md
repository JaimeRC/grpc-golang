# Unary RPC

## Ciclo de vida:
1. Una vez que el cliente llama a un método stub, se notifica al servidor que se ha invocado el RPC con los metadatos del cliente para esta llamada, el nombre del método y la fecha límite especificada , si corresponde.
2. El servidor puede enviar de vuelta sus propios metadatos iniciales (que deben enviarse antes de cualquier respuesta) de inmediato, o esperar el mensaje de solicitud del cliente. Lo que sucede primero, es específico de la aplicación.
3. Una vez que el servidor tiene el mensaje de solicitud del cliente, hace todo el trabajo necesario para crear y completar una respuesta. Luego, la respuesta se devuelve (si tiene éxito) al cliente junto con los detalles del estado (código de estado y mensaje de estado opcional) y metadatos finales opcionales.
4. Si el estado de la respuesta es correcto, entonces el cliente recibe la respuesta, que completa la llamada del lado del cliente.

## Código
- [Interfaz proto](./proto/hello.proto)
- [Servidor](./server/main.go)
- [Cliente](./client/main.go)


## Inicializar:
- Servidor:
```
    cd server/
    go run main.go
```
- Cliente:
```
    cd client/
    go run main.go
        2020/12/26 18:47:05 Greeting: Hello Pepe
        2020/12/26 18:47:05 Greeting: Hello again Pepe
  ```
