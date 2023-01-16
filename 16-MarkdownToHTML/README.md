# Introducción

Para convertir un archivo markdown `.md` a un archivo HTML `.html` usando golang, vamos a necesitar instalar el siguiente paquete:

```bash
go get github.com/gomarkdown/markdown
```

## Crear el archivo .html en base al archivo .md

Para crear el archivo html lo primero que debemos de hacer es leer el contenido de este, que es el que vamos a estar agregando al archivo html.

```go
file := "test.md"

content, err := os.ReadFile(file)
if err != nil {
  log.Fatalf("%s file not found", file)
}
```

Una vez que ya tenemos el contenido del archivo guardado en una variable se lo inyectamos a un archivo html.

```go
html := markdown.ToHTML(content, nil, nil)
fmt.Println(string(html))

fileOut := "test.html"
err = os.WriteFile(fileOut, html, 0644)
if err != nil {
  log.Fatalf("Could not write to %s", fileOut)
}

fmt.Printf("HTML outputted to %s", fileOut)
```
