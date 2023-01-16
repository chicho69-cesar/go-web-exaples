# Hashing de contraseñas con bcrypt

Este ejemplo mostrará cómo codificar contraseñas usando bcrypt. Para esto tenemos que usar la biblioteca de golang bcrypt, aunque primero debemos de instalarla:

```bash
go get golang.org/x/crypto/bcrypt
```

De ahora en adelante, cada aplicación que escribamos podrá hacer uso de esta biblioteca.

## Creación de Hash de contraseña

Una de las principales funciones que nos ofrece esta libreria es convertir una contraseña en un hash cifrado, para esto usamos la funcion `GenerateFromPassword` la cual acepta dos parametros, donde el primero es un array de bytes de la contraseña que queremos cifrar y el segundo es el costo en bytes del cifrado, el cual tiene como limite 14 bytes.

```go
func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}
```

## Comparar dos claves cifradas

Para comparar un texto con una contraseña cifrada, solo debemos de utilizar la funcion `CompareHashAndPassword` la cual recibe como parametros un array de bytes con el hash y un array de bytes con el texto con la que lo queremos comparar.

```go
func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
```
