# Introducción

Go es un lenguaje de programación que incluye batería y tiene un servidor web ya incorporado. El paquete `net/http` paquete de la biblioteca estándar contiene todas las funcionalidades sobre el protocolo HTTP. Esto incluye (entre muchas otras cosas) un cliente HTTP y un servidor HTTP. En este ejemplo, descubrirá lo simple que es crear un servidor web que pueda ver en su navegador.

## Registro de un controlador de soluciones

Primero, cree un controlador que reciba todas las conexiones HTTP entrantes de navegadores, clientes HTTP o solicitudes de API. Un controlador en Go es una función con esta firma:

```go
func (w http.ResponseWriter, r *http.Request)
```

La función recibe dos parámetros:
Un http.ResponseWriter que es donde escribe su respuesta de texto/html.
Un http.Request que contiene toda la información sobre esta solicitud HTTP, incluidos elementos como la URL o los campos de encabezado.
Registrar un controlador de solicitudes en el servidor HTTP predeterminado es tan simple como esto:

```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})
```

## Escuchar las conexiones HTTP

El controlador de solicitudes por sí solo no puede aceptar ninguna conexión HTTP desde el exterior. Un servidor HTTP tiene que escuchar en un puerto para pasar conexiones al controlador de solicitudes. Debido a que el puerto 80 es en la mayoría de los casos el puerto predeterminado para el tráfico HTTP, este servidor también lo escuchará.
El siguiente código iniciará el servidor HTTP predeterminado de Go y escuchará las conexiones en el puerto 80.
Puede navegar con su navegador `http://localhost/` y ver su servidor entregando su solicitud.

```go
http.ListenAndServe(":80", nil)
```