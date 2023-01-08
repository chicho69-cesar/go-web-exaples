# Introducción

El paquete de Go `html/template` proporciona un rico lenguaje de plantillas para plantillas HTML. Se utiliza principalmente en aplicaciones web para mostrar datos de forma estructurada en el navegador de un cliente. Un gran beneficio del lenguaje de plantillas de Go es el escape automático de datos. No hay necesidad de preocuparse por los ataques XSS, ya que Go analiza la plantilla HTML y escapa de todas las entradas antes de mostrarla en el navegador.

## Primera plantilla

Escribir una plantilla en Go es muy simple. Este ejemplo muestra una lista TODO, escrita como una lista desordenada (ul) en HTML. Al renderizar plantillas, los datos que se pasan pueden ser cualquier tipo de estructuras de datos de Go. Puede ser una cadena simple o un número, incluso puede ser una estructura de datos anidada como en el ejemplo a continuación. Para acceder a los datos en una plantilla, la variable más alta es el acceso por `{{.}}`. El punto dentro de las llaves se llama canalización y el elemento raíz de los datos.

```go
data := TodoPageData{
  PageTitle: "My TODO list",
  Todos: []Todo {
    { Title: "Task 1", Done: false },
    { Title: "Task 2", Done: true },
    { Title: "Task 3", Done: true },
  },
}
```

```html
<h1>{{.PageTitle}}</h1>
<ul>
  {{range .Todos}}
    {{if .Done}}
      <li class="done">{{.Title}}</li>
    {{else}}
      <li>{{.Title}}</li>
    {{end}}
  {{end}}
</ul>
```

## Estructuras de Control

El lenguaje de plantillas contiene un rico conjunto de estructuras de control para representar su HTML. Aquí obtendrá una descripción general de los más utilizados. Para obtener una lista detallada de todas las estructuras posibles, visite: `text/template`

| Estructura de control     | Definición                 |
| ------------------------- | -------------------------- |
| {{/* a comment */}}       | Define un comentario       |
| {{.}}                     | Renderiza el elemento raíz |
| {{.Title}}                | Representa el campo "Title" en un elemento anidado |
| {{if .Done}} {{else}} {{end}} | Define una sentencia if |
| {{range .Todos}} {{.}} {{end}} | Recorre todos los "Todos" y renderiza cada uno usando {{.}} |
| {{block "content" .}} {{end}} | Define un bloque con el nombre “contenido” |

## Análisis de plantillas de archivos

La plantilla se puede analizar desde una cadena o un archivo en el disco. Como suele ser el caso, las plantillas se copian desde el disco, este ejemplo muestra cómo hacerlo. En este ejemplo, hay un archivo de plantilla en el mismo directorio que el programa Go llamado `layout.html`.

```go
tmpl, err := template.ParseFiles("layout.html")
// or
tmpl := template.Must(template.ParseFiles("layout.html"))
```

## Ejecutar una plantilla en un controlador de solicitudes

Una vez que la plantilla se analiza desde el disco, está lista para usarse en el controlador de solicitudes. La función `Execute` acepta un tipo `io.Writer` para escribir la plantilla y una `interface{}` para pasar datos a la plantilla.
Cuando se llama a la función en un encabezado `http.ResponseWriter`, el tipo de contenido se establece automáticamente en la respuesta HTTP a `Content-Type: text/html; charset=utf-8.`

```go
func(w http.ResponseWriter, r *http.Request) {
  tmpl.Execute(w, "data goes here")
}
```