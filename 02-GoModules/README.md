# Modulos en Go

Los modulos en Go nos ayudan a tener una correcta organizacion del codigo que realizamos en los proyectos, los modulos se organizan por carpetas, comunmente el archivo principal donde tenemos la función `main` que inicia el proyecto se encuentra en la carpeta raiz del proyecto y el paquete suele llamarse tambien como:

```go
package main
```

Despues los archivos que se encuentren en carpetas subyacentes a este archivo suelen tener como nombre de package el nombre del directorio en el que se encuentran, por ejemplo: 

```go
// /maths/operations.go
package maths
```

```go
// /api/controllers/user-controller.go
package controllers
```

```go
// /server/server.go
package server
```

Para iniciar un proyecto de Go donde vamos a usar modulos debemos de usar el comando:

```bash
go mod init example.com/my-project
```

Donde *example.com/my-project* va a ser la ruta a traves de la cual vamos a poder importar los demas packages que definamos en nuestro proyecto. Una vez que ejecutamos este comando se nos va a crear un archivo `go.mod` el cual lo vamos a tratar como por ejemplo lo que en JavaScript seria el `package.json` es decir donde vamos a tener todos los packages que instalemos en nuestro proyecto.
La estructura del archivo quedaria algo asi:

```go
module example.com/my-project

go 1.19
```

## Instalar paquetes de terceros

Para instalar paquetes de terceros en nuestro proyecto debemos de usar el comando `go get package`, donde en el package ponemos la ruta del package que queremos agregar al proyecto, un ejemplo donde instalamos gorilla/mux seria:

```bash
go get github.com/gorilla/mux
```

El archivo `go.mod` con la instalacion del paquete quedaria algo así:

```go
module example.com/my-project

go 1.19

require github.com/gorilla/mux v1.8.0 // indirect
```

## Cachear los paquetes y otras utilidades

El subcomando de go `go mod` tiene varias opciones de ser utilizado como ya lo hemos visto, otras de las utilidades que tiene es por ejemplo el poder cachear los packages en nuestras computadoras para que cuando lo queramos volver a usar lo podamos descargar de una manera mas rapida. Para esto usamos el comando:

```bash
go mod download package
```

Para remover los paquetes que no estan siendo usados en nuestro proyecto usamos el comando `tidy` el cual remueve del archivo `go.mod` los paquetes que agregamos al proyecto pero que no los estamos utilizando.

```bash
go mod tidy
```