# Introducción

Este ejemplo mostrará cómo almacenar datos en cookies de sesión usando el popular paquete `gorilla/sessions` en Go.
Las cookies son pequeños fragmentos de datos que se almacenan en el navegador de un usuario y se envían a nuestro servidor con cada solicitud. En ellos, podemos almacenar, por ejemplo, si un usuario ha iniciado sesión o no en nuestro sitio web y averiguar quién es realmente (en nuestro sistema).
En este ejemplo, solo permitiremos que los usuarios autenticados vean nuestro mensaje secreto en la página `/secret`. Para obtener acceso a él, primero tendrá que visitar `/login` para obtener una cookie de sesión válida, que lo registra. Además, puede visitar `/logout` para revocar su acceso a nuestro mensaje secreto.

Lo primero que vamos a hacer es instalar el paquete de gorilla/sessions, para lo cual hacemos uso del comando: 

```bash
go get github.com/gorilla/sessions
```

## Instanciar la clave de acceso para las cookies

Lo primero que debemos hacer es definir una store para almacenar las cookies de los usuarios, para lo que hacemos es crear primero un array de bytes con la clave secreta y despues crear la instancia de nuestra cookie store.

```go
key = []byte("super-secret-key")
store = sessions.NewCookieStore(key)
```

Con esto ya podremos definir nuestras funciones handler para las peticiones http que reciba el servidor a traves de las cuales vamos a manejar las sesiones.

## Login

Lo primero que debemos de hacer es generar un cookie de sesion a traves de la cual vamos a ver si el usuario ya ha sido autenticado en nuestro servidor(*la parte de la autenticacion la vamos a obviar, se va a considerar como el usuario si esta autenticado*).
Despues lo que hacemos es marcar la sesion como autenticada y guardarla.

```go
func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
  // ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}
```

## Acceder al mensaje secreto

Una vez que el usuario ya ha iniciado sesion en nuestro servidor, este ya podra acceder a la peticion `http` que esta reservada solo para usuarios logeados, para poder comprobar que un usuario realmente esta logeado en nuestro servidor lo que hacemos es que obtenemos nuestra cookie de sesion y a partir de esta verificamos si el usuario esta autenticado o si diña cookie existe, si no se cumple ninguna de las dos, lanzamos un error con un codigo de respuesta 403 o `Forbidden`, si todo va bien le mostramos la información ya que el usuario si tiene acceso.

```go
func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}
```

## Logout

La funcion de logout es realmente sencilla ya que lo unico que debemos de hacer es en la cookie marcar como que el usuario no esta autenticado, despues guardamos los cambios sobre la cookie para el usuario y de esta manera dicho ya no podra acceder al contenido asegurado.

```go
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
```