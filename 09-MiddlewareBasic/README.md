# Introducción 

Este ejemplo mostrará cómo crear un middleware de registro básico en Go.
Un middleware simplemente toma a `http.HandlerFunc` como uno de sus parámetros, lo envuelve y devuelve uno nuevo `http.HandlerFunc` para que el servidor lo llame.

Por lo que una vez que lo llame podremos ejecutar una serie de operaciones adicionales sobre la peticion, esto principalmente se usa para temas de seguridad en las aplicaciones web.

La firma de un middleware es la siguiente: 

```go
func middleware(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // the extra actions
    f(w, r)
  }
}
```