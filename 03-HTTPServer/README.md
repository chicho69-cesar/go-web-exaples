# Intoducción

En este ejemplo, aprenderá cómo crear un servidor HTTP básico en Go. Primero, hablemos de lo que debería ser capaz de hacer nuestro servidor HTTP. Un servidor HTTP básico tiene algunos trabajos clave de los que ocuparse.
*Procesar solicitudes dinámicas:* Procesar solicitudes entrantes de usuarios que navegan por el sitio web, inician sesión en sus cuentas o publican imágenes.
*Servir activos estáticos:* Servir archivos JavaScript, CSS e imágenes a los navegadores para crear una experiencia dinámica para el usuario.
Aceptar conexiones: El servidor HTTP debe escuchar en un puerto específico para poder aceptar conexiones desde Internet.

## Procesar solicitudes dinámicas

El paquete `net/http` contiene todas las utilidades necesarias para aceptar solicitudes y manejarlas dinámicamente. Podemos registrar un nuevo controlador con la `http.HandleFunc` función. Su primer parámetro toma una ruta para coincidir y una función para ejecutar como segundo. En este ejemplo: cuando alguien navega por sus sitios web ( http://example.com/), será recibido con un bonito mensaje.

```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to my website!")
})
```

Para el aspecto dinámico, `http.Request` contiene toda la información sobre la solicitud y sus parámetros. Puede leer parámetros GET con `r.URL.Query().Get("token")` o parámetros POST (campos de un formulario HTML) con `r.FormValue("email")`.

## Sirviendo activos estáticos

Para servir activos estáticos como JavaScript, CSS e imágenes, usamos el incorporado `http.FileServer` y lo apuntamos a una ruta de URL. Para que el servidor de archivos funcione correctamente, necesita saber desde dónde entregar los archivos. Podemos hacer esto así:

```go
fs := http.FileServer(http.Dir("static/"))
```

Una vez que nuestro servidor de archivos está en su lugar, solo necesitamos señalarle una ruta de URL, tal como lo hicimos con las solicitudes dinámicas. Una cosa a tener en cuenta: para servir los archivos correctamente, debemos eliminar una parte de la ruta de la URL. Por lo general, este es el nombre del directorio en el que se encuentran nuestros archivos.

```go
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

## Aceptar conexiones

Lo último para terminar nuestro servidor HTTP básico es escuchar en un puerto para aceptar conexiones desde Internet. Como puede adivinar, Go también tiene un servidor HTTP incorporado, podemos comenzar rápidamente. Una vez iniciado, puede ver su servidor HTTP en su navegador.

```go
http.ListenAndServe(":80", nil)
```