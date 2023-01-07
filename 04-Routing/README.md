# Introducción

El paquete de Go `net/http` proporciona muchas funcionalidades para el protocolo HTTP. Una cosa que no hace muy bien es el enrutamiento de solicitudes complejas, como segmentar una URL de solicitud en parámetros únicos. Afortunadamente, existe un paquete muy popular para esto, que es bien conocido por la buena calidad del código en la comunidad Go. En este ejemplo, se verá cómo usar el `gorilla/mux` para crear rutas con parámetros con nombre, controladores GET/POST y restricciones de dominio.

## Instalando el paquete gorilla/mux

`gorilla/mux` es un paquete que se adapta al enrutador HTTP predeterminado de Go. Viene con muchas funciones para aumentar la productividad al escribir aplicaciones web. También cumple con la firma del controlador de solicitudes predeterminado de Go `func (w http.ResponseWriter, r *http.Request)`, por lo que el paquete se puede mezclar y combinar con otras bibliotecas HTTP como middleware o aplicaciones existentes. Usar el comando `go get` para instalar el paquete desde GitHub así:

```bash
go get -u github.com/gorilla/mux
```

## Creación de un nuevo enrutador

Primero cree un nuevo enrutador de solicitud. El enrutador es el enrutador principal para su aplicación web y luego se pasará como parámetro al servidor. Recibirá todas las conexiones HTTP y las pasará a los controladores de solicitudes que registrará en él. Puede crear un nuevo enrutador así:

```go
r := mux.NewRouter()
```

## Registro de un controlador de solicitudes

Una vez que tenga un nuevo enrutador, puede registrar controladores de solicitudes como de costumbre. La única diferencia es que, en lugar de llamar `http.HandleFunc(...)` a , usted llama `HandleFunc` a su enrutador de esta manera: `r.HandleFunc(...)`.

## Parámetros de URL

La mayor fortaleza del enrutador `gorilla/mux` es la capacidad de extraer segmentos de la URL de solicitud. Como ejemplo, esta es una URL en su aplicación:

```txt
/books/go-programming-blueprint/page/10
```

Esta URL tiene dos segmentos dinámicos:
Slug del título del libro (go-programming-blueprint)
Página (10)
Para que un controlador de solicitudes coincida con la URL mencionada anteriormente, reemplace los segmentos dinámicos con marcadores de posición en su patrón de URL de la siguiente manera:

```go
r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
  // get the book
  // navigate to the page
})
```

Lo último es obtener los datos de estos segmentos. El paquete viene con la función `mux.Vars(r)` que toma `http.Request` como parámetro y devuelve un mapa de los segmentos.

```go
func(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  vars["title"] // the book title slug
  vars["page"] // the page
}
```

## Configuración del enrutador del servidor HTTP

`nil` ¿Alguna vez te has preguntado qué es lo que `http.ListenAndServe(":80", nil)` pasa? Es el parámetro para el enrutador principal del servidor HTTP. Por defecto es `nil`, lo que significa usar el enrutador predeterminado del paquete `net/http`. Para hacer uso de su propio enrutador, reemplace `nil` con la variable de su enrutador r.

## Características del enrutador gorilla/mux

### Métodos

Restringir el controlador de solicitudes a métodos HTTP específicos.

```go
r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
```

### Nombres de host y subdominios

Restringir el controlador de solicitudes a nombres de host o subdominios específicos.

```go
r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")
```

### Esquemas

Restringir el controlador de solicitudes a http/https.

```go
r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")
```

### Prefijos de ruta y subenrutadores

Restringir el controlador de solicitudes a prefijos de ruta específicos.

```go
bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)
```