# grpc-golang

## Descripción

En gRPC, una aplicación cliente puede llamar directamente a un método en una aplicación de servidor en una máquina diferente como si fuera un objeto local, lo que le facilita la creación de aplicaciones y servicios distribuidos. Como en muchos sistemas RPC, gRPC se basa en la idea de definir un servicio, especificando los métodos que se pueden llamar de forma remota con sus parámetros y tipos de retorno. En el lado del servidor, el servidor implementa esta interfaz y ejecuta un servidor gRPC para manejar las llamadas de los clientes. En el lado del cliente, el cliente tiene un código auxiliar (denominado simplemente cliente en algunos idiomas) que proporciona los mismos métodos que el servidor.

![gRPC](./image.svg)

Los clientes y servidores de gRPC pueden ejecutarse y comunicarse entre sí en una variedad de entornos, desde servidores dentro de Google hasta su propio escritorio, y pueden escribirse en cualquiera de los idiomas admitidos por gRPC. Entonces, por ejemplo, puede crear fácilmente un servidor gRPC en Java con clientes en Go, Python o Ruby. Además, las últimas API de Google tendrán versiones gRPC de sus interfaces, lo que le permitirá crear fácilmente la funcionalidad de Google en sus aplicaciones.

## Tipos de **gRPC**:
### **Unary**
En los que el cliente envía una única solicitud al servidor y obtiene una única respuesta, al igual que una llamada de función normal

[Más info...](./unary/README.md)

### **Server streaming**
Donde el cliente envía una solicitud al servidor y obtiene una transmisión para leer una secuencia de mensajes. El cliente lee el flujo devuelto hasta que no hay más mensajes. gRPC garantiza el orden de los mensajes dentro de una llamada RPC individual

[Más info...](./server_streaming/README.md)


### **Client streaming**
En las que el cliente escribe una secuencia de mensajes y los envía al servidor, nuevamente utilizando una secuencia proporcionada. Una vez que el cliente ha terminado de escribir los mensajes, espera a que el servidor los lea y devuelva su respuesta. Una vez más, gRPC garantiza el orden de los mensajes dentro de una llamada RPC individual

[Más info...](./client_streaming/README.md)


### **Bi-directional streaming**
Donde ambos lados envían una secuencia de mensajes mediante una transmisión de lectura y escritura. Los dos flujos funcionan de forma independiente, por lo que los clientes y los servidores pueden leer y escribir en el orden que deseen: por ejemplo, el servidor podría esperar a recibir todos los mensajes del cliente antes de escribir sus respuestas, o alternativamente podría leer un mensaje y luego escribir un mensaje, o alguna otra combinación de lecturas y escrituras. Se conserva el orden de los mensajes en cada flujo

[Más info...](./bi_directional_streaming/README.md)
