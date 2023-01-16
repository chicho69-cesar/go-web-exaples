# Introducción

Para trabajar convirtiendo archivos con la extension `.txt` a un archivo `pdf` es necesario usar el paquete `gofpdf` el cual lo podemos agregar a nuestro proyecto usamos el comando:

```bash
go get github.com/jung-kurt/gofpdf
```

## Lectura del archivo .txt

Lo primero que debemos de hacer es leer el contenido del archivo usando la funcion del paquete `os.ReadFile` el cual recibe el nombre o ruta del archivo y mediante esto obtenemos el contenido.

```go
var file string = "test.txt"

content, err := os.ReadFile(file)
if err != nil {
  log.Fatalf("%s file not found", file)
}
```

## Creacion del archivo pdf

Para crear el pdf usamos el paquete que instalamos anteriormente y le agregamos el contenido que leimos de dicho archivo. Para esto usamos la funcion New la cual recibe como parametros, la orientacion la cual puede ser Portrait (P) o Landscape (L), despues la unidad de medida la cual es recomendable usarla en milimetros mm, despues el tamaño de la hoja, la cual puede ser de muchas medidas pero la mas tipica es A4 y por ultimo viene la direccion del texto, el cual si ponemos un string vacio se coloca de forma normal, por ejemplo podemos crear un pdf vacico con el siguiente codigo.

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.AddPage()
pdf.SetFont("Arial", "B", 14)
pdf.MultiCell(190, 5, string(content), "0", "0", false)

_ = pdf.OutputFileAndClose("test.pdf")

fmt.Println("PDF CREATED...")
```
