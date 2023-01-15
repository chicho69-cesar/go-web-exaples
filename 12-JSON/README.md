# Introducción

En este ejemplo veremos como codificar y decodificar información en formato JSON usando el paquete `encoding/json`.

Para ello vamos a crear un servidor `http` el cual reciba a traves de peticiones http elementos de una estructura que definiremos como User y es la que se muestra a continuacion: 

```go
type User struct {
	Firstname string `json:"firstname"`
	Lastname 	string `json:"lastname"`
	Age 			int 	 `json:"age"`
}
```

Como podemos ver el tercer campo de cada elemento de la estructura define el nombre que vamos a decodificar o a codificar en formato JSON desde la peticion http.

## Decodificar JSON a tipo User

Para hacer la decodificacion de un JSON que nos llega como cuerpo de una peticion http lo que necesitamos es un objeto decoder del paquete json, y decodificar el cuerpo de una request sobre una variable de tipo User:

```go
http.HandleFunc("/decode", func(w http ResponseWriter, r *http.Request) {
  var user User
  json.NewDecoder(r.Body).Decode(&user)

  fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
})
```

## Codificar de tipo User a JSON

Para realizar la codificacion hacemos el proceso inverso a la decodificacion, y a traves de una variable de tipo User la codificamos a JSON con un objeto Encoder del paquete json.

```go
http.HandleFunc("/enconde", func(w http.ResponseWriter, r *http.Request) {
  peter := User {
    Firstname: "Jhon",
    Lastname: "Doe",
    Age: 25,
  }

  json.NewEncoder(w).Encode(peter)
  bs, _ := json.Marshal(peter)
  fmt.Println(string(bs))
})
```