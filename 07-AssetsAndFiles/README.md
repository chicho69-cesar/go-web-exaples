# Introducción

Este ejemplo mostrará cómo servir archivos estáticos como CSS, JavaScript o imágenes desde un directorio específico.

Para hacer esto es tan facil como definir una carpeta donde vamos a tener nuestros archivos estaticos, en este caso podemos definir una carpeta `assets` donde almacenaremos todos nuestros archivos estaticos, dentro de esta carpeta podemos definir una nueva carpeta a la cual llamaremos `css`, dentro crearemos un archivo `style.css` el cual tendra el siguiente contenido: 

```css
body {
  background-color: #09f;
}
```

Despues en nuestro programa de go, lo primero que haremos sera definir el directorio de archivos estaticos al cual haremos referencia:

```go
fs := http.FileServer(http.Dir("assets/"))
```

Despues definimos la ruta a través de la cual vamos a acceder a nuestros archivos estaticos:

```go
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

Por ultimo solo ponemos en marcha el servidor y accedemos a la ruta: `http://localhost:80/static/css/style.css`

```go
http.ListenAndServe(":80", nil)
```