# Bi-directional streaming RPC

## Ciclo de vida:

En un RPC de transmisión bidireccional, el cliente inicia la llamada invocando el método y el servidor que recibe los metadatos del cliente, el nombre del método y la fecha límite. El servidor puede optar por enviar sus metadatos iniciales o esperar a que el cliente comience a transmitir mensajes.

El procesamiento de secuencias del lado del cliente y del servidor es específico de la aplicación. Dado que los dos flujos son independientes, el cliente y el servidor pueden leer y escribir mensajes en cualquier orden. Por ejemplo, un servidor puede esperar hasta que haya recibido todos los mensajes de un cliente antes de escribir sus mensajes, o el servidor y el cliente pueden jugar "ping-pong": el servidor recibe una solicitud, luego envía una respuesta y luego el cliente envía otra solicitud basada en la respuesta, y así sucesivamente.

## Código
- [Interfaz proto](./proto/feed.proto)
- [Servidor](./server/main.go)
- [Cliente](./client/main.go)

## Inicializar:
- Servidor:
```
    cd server/
    go run main.go
        sending new feed...
        sending new feed...
        sending new feed...
        sending new feed...
        sending new feed...
```
- Cliente:
```
    cd client/
    go run main.go
        Bew feed received:  New Feed Recieved: This is feed number 1
        Bew feed received:  New Feed Recieved: This is feed number 2
        Bew feed received:  New Feed Recieved: This is feed number 3
        Bew feed received:  New Feed Recieved: This is feed number 4
        Bew feed received:  New Feed Recieved: This is feed number 5
 ```