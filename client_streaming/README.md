# gRPC: Client streaming

## Ciclo de vida:

Un RPC de flujo de cliente es similar a un RPC unario, excepto que el cliente envía un flujo de mensajes al servidor en lugar de un solo mensaje. El servidor responde con un solo mensaje (junto con sus detalles de estado y metadatos finales opcionales), normalmente, pero no necesariamente, después de haber recibido todos los mensajes del cliente.

## Código
- [Interfaz proto](./proto/sumAll.proto)
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
        sending 0 into the stream
        sending 1 into the stream
        sending 2 into the stream
        sending 3 into the stream
        sending 4 into the stream
        sending 5 into the stream
        sending 6 into the stream
        sending 7 into the stream
        sending 8 into the stream
        sending 9 into the stream
        sending 10 into the stream
        Sum of numbers:  55
 ```