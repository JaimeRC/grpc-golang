# Server streaming RPC 

## Ciclo de vida:

Una RPC de transmisión por servidor es similar a una RPC unaria, excepto que el servidor devuelve una secuencia de mensajes en respuesta a la solicitud de un cliente. Después de enviar todos sus mensajes, los detalles del estado del servidor (código de estado y mensaje de estado opcional) y metadatos finales opcionales se envían al cliente. Esto completa el procesamiento en el lado del servidor. El cliente completa una vez que tiene todos los mensajes del servidor.

## Código
- [Interfaz proto](./proto/countdown.proto)
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
        count:10
        count:9
        count:8
        count:7
        count:6
        count:5
        count:4
        count:3
        count:2
        count:1
 ```