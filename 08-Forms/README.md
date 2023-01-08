# Introducción 

A continuación vamos a ver la forma en la cual podemos trabajar con formularios de una pagina web desde Go, para esto vamos a crear un servidor para escuchar las peticiones http y ademas vamos a hacer uso del paquete `html/template` para renderizar las vistas html desde el navegador.
La principal forma de ver si la peticion que nos llega es un post, como resultado de hacer submit sobre un formulario tan solo utilizamos la siguiente condicion:

```go
if r.Method != http.MethodPost {
  tmpl.Execute(w, nil)
  return
}
```

Para obtener los valores que nos llegan a traves de la peticion por el formulario hacemos uso del metodo `FormValue` tal y como se muestra a continuación:

```go
field := r.FormValue("field")
```

Por ultimo para enviar un valor al formulario hacemos lo siguiente: 

```go
tmpl.Execute(w, struct {
  Value bool
} {
  Value: true,
})
```

En la vista html ponemos:

```html
{{if .Value}}
  <h1>
    Congratulations! your response is {{.Value}}
  </h1>
{{else}}
  <h1>
    Oh no :( your response is incorrect
  </h1>
{{end}}
```