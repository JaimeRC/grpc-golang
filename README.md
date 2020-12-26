# grpc-golang

## Descripción

En gRPC, una aplicación cliente puede llamar directamente a un método en una aplicación de servidor en una máquina diferente como si fuera un objeto local, lo que le facilita la creación de aplicaciones y servicios distribuidos. Como en muchos sistemas RPC, gRPC se basa en la idea de definir un servicio, especificando los métodos que se pueden llamar de forma remota con sus parámetros y tipos de retorno. En el lado del servidor, el servidor implementa esta interfaz y ejecuta un servidor gRPC para manejar las llamadas de los clientes. En el lado del cliente, el cliente tiene un código auxiliar (denominado simplemente cliente en algunos idiomas) que proporciona los mismos métodos que el servidor.

![gRPC](./image.svg)

Los clientes y servidores de gRPC pueden ejecutarse y comunicarse entre sí en una variedad de entornos, desde servidores dentro de Google hasta su propio escritorio, y pueden escribirse en cualquiera de los idiomas admitidos por gRPC. Entonces, por ejemplo, puede crear fácilmente un servidor gRPC en Java con clientes en Go, Python o Ruby. Además, las últimas API de Google tendrán versiones gRPC de sus interfaces, lo que le permitirá crear fácilmente la funcionalidad de Google en sus aplicaciones.

## Tipos de **gRPC**:
### **Unary**
En los que el cliente envía una única solicitud al servidor y obtiene una única respuesta, al igual que una llamada de función normal

**Ciclo de vida:**
1. Una vez que el cliente llama a un método stub, se notifica al servidor que se ha invocado el RPC con los metadatos del cliente para esta llamada, el nombre del método y la fecha límite especificada , si corresponde.
2. El servidor puede enviar de vuelta sus propios metadatos iniciales (que deben enviarse antes de cualquier respuesta) de inmediato, o esperar el mensaje de solicitud del cliente. Lo que sucede primero, es específico de la aplicación.
3. Una vez que el servidor tiene el mensaje de solicitud del cliente, hace todo el trabajo necesario para crear y completar una respuesta. Luego, la respuesta se devuelve (si tiene éxito) al cliente junto con los detalles del estado (código de estado y mensaje de estado opcional) y metadatos finales opcionales.
4. Si el estado de la respuesta es correcto, entonces el cliente recibe la respuesta, que completa la llamada del lado del cliente.


**Código ejemplo:** [Unary](./unary)

**Inicializar:**
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


### **Server streaming**
Donde el cliente envía una solicitud al servidor y obtiene una transmisión para leer una secuencia de mensajes. El cliente lee el flujo devuelto hasta que no hay más mensajes. gRPC garantiza el orden de los mensajes dentro de una llamada RPC individual

**Ciclo de vida:**

Una RPC de transmisión por servidor es similar a una RPC unaria, excepto que el servidor devuelve una secuencia de mensajes en respuesta a la solicitud de un cliente. Después de enviar todos sus mensajes, los detalles del estado del servidor (código de estado y mensaje de estado opcional) y metadatos finales opcionales se envían al cliente. Esto completa el procesamiento en el lado del servidor. El cliente completa una vez que tiene todos los mensajes del servidor.

**Código ejemplo:** [Server streaming](./server_streaming)

**Inicializar:**
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

### **Client streaming**
En las que el cliente escribe una secuencia de mensajes y los envía al servidor, nuevamente utilizando una secuencia proporcionada. Una vez que el cliente ha terminado de escribir los mensajes, espera a que el servidor los lea y devuelva su respuesta. Una vez más, gRPC garantiza el orden de los mensajes dentro de una llamada RPC individual

**Ciclo de vida:**

Un RPC de flujo de cliente es similar a un RPC unario, excepto que el cliente envía un flujo de mensajes al servidor en lugar de un solo mensaje. El servidor responde con un solo mensaje (junto con sus detalles de estado y metadatos finales opcionales), normalmente, pero no necesariamente, después de haber recibido todos los mensajes del cliente.

**Código ejemplo:** [Client streaming](./client_streaming)

**Inicializar:**
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

### **Bi-directional streaming**
Donde ambos lados envían una secuencia de mensajes mediante una transmisión de lectura y escritura. Los dos flujos funcionan de forma independiente, por lo que los clientes y los servidores pueden leer y escribir en el orden que deseen: por ejemplo, el servidor podría esperar a recibir todos los mensajes del cliente antes de escribir sus respuestas, o alternativamente podría leer un mensaje y luego escribir un mensaje, o alguna otra combinación de lecturas y escrituras. Se conserva el orden de los mensajes en cada flujo

**Ciclo de vida:**

En un RPC de transmisión bidireccional, el cliente inicia la llamada invocando el método y el servidor que recibe los metadatos del cliente, el nombre del método y la fecha límite. El servidor puede optar por enviar sus metadatos iniciales o esperar a que el cliente comience a transmitir mensajes.

El procesamiento de secuencias del lado del cliente y del servidor es específico de la aplicación. Dado que los dos flujos son independientes, el cliente y el servidor pueden leer y escribir mensajes en cualquier orden. Por ejemplo, un servidor puede esperar hasta que haya recibido todos los mensajes de un cliente antes de escribir sus mensajes, o el servidor y el cliente pueden jugar "ping-pong": el servidor recibe una solicitud, luego envía una respuesta y luego el cliente envía otra solicitud basada en la respuesta, y así sucesivamente.

**Código ejemplo:** [Vi-directional streaming](./bi_directional_streaming)

**Inicializar:**
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