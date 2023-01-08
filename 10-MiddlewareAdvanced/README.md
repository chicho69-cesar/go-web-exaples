# Introducción

Este ejemplo mostrará cómo crear una versión más avanzada de middleware en Go.
Un middleware en sí mismo simplemente toma a `http.HandlerFunc` como uno de sus parámetros, lo envuelve y devuelve uno nuevo `http.HandlerFunc` para que el servidor lo llame.
Aquí definimos un nuevo tipo `Middleware` que eventualmente hace que sea más fácil encadenar varios middleware juntos. Esta idea está inspirada en la charla de Mat Ryers sobre la creación de API. Puedes encontrar una explicación más detallada incluyendo la charla aquí .
Este fragmento explica en detalle cómo se crea un nuevo middleware. En el ejemplo completo a continuación, reducimos esta versión con un código repetitivo.

Primero creamos nuestro tipo Middleware:

```go
type Middleware func(http.HandlerFunc) http.HandlerFunc
```

Despues definimos nuestra funcion para crear un Midleware:

```go
func createNewMiddleware() Middleware {
  // Create a new Middleware
  middleware := func(next http.HandlerFunc) http.HandlerFunc {
    // Define the http.HandlerFunc which is called by the server eventually
    handler := func(w http.ResponseWriter, r *http.Request) {
      // ... do middleware things

      // Call the next middleware/handler in chain
      next(w, r)
    }

    // Return newly created handler
    return handler
  }

  // Return newly created middleware
  return middleware
}
```