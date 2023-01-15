# Intoducción

Este ejemplo mostrará cómo trabajar con websockets en Go. Construiremos un servidor simple que repita todo lo que le enviemos. Para esto tenemos que instalar una biblioteca muy popular para utilizar websockets, esta biblioteca es propia del famoso toolkit `Gorilla Web Toolkit` utilizando el comando:

```bash
go get github.com/gorilla/websocket
```

De ahora en adelante, cada aplicación que escribamos podrá hacer uso de esta biblioteca.

## Crear un servidor de Sockets con Go

Para crear nuestro servidor que acepte conexiones a traves de sockets, lo que vamos a hacer primero es definir un objeto `upgrader` que es el que nos va a ayudar a controlar el tamaño del buffer de lectura y escritura de nuestros sockets, para ello usamos: 


```go
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
```

Despues nos creamos un endpoint para manejar las conexiones a traves de sockets, donde vamos a crear un objeto de conexion mediante el objeto upgrader y creamos un ciclo infinito para mantener al servidor escuchando los mensajes de los sockets, leemos el tipo de mensaje, el mensaje y un posible que pueda haber y verificamos que no haya ningun error mientras escuchamos los mensajes del socket, y si lo hay rompemos la ejecucion del servidor de sockets.

```go
http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
  conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

  for {
    // Read message from browser
    msgType, msg, err := conn.ReadMessage()
    if err != nil {
      return
    }

    // Print the message to the console
    fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

    // Write message back to browser
    if err = conn.WriteMessage(msgType, msg); err != nil {
      return
    }
  }
})
```

Despues montamos nuestro documento de `websockets.html` el cual a traves de codigo *JavaScript* se va a conectar a nuestro servidor de sockets y va a establecer una comunicacion directa:

```html
<input id="input" type="text" />
<button onclick="send()">Send</button>
<pre id="output"></pre>

<script>
  var input = document.getElementById("input");
  var output = document.getElementById("output");
  var socket = new WebSocket("ws://localhost:5173/echo");

  socket.onopen = function () {
    output.innerHTML += "Status: Connected\n";
  };

  socket.onmessage = function (e) {
    output.innerHTML += "Server: " + e.data + "\n";
  };

  function send() {
    socket.send(input.value);
    input.value = "";
  }
</script>
```

En el servidor: 

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "websockets.html")
})
```