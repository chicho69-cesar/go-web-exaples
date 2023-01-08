# Intoducción

En algún momento, desea que su aplicación web almacene y recupere datos de una base de datos. Este es casi siempre el caso cuando se trata de contenido dinámico, sirviendo formularios para que los usuarios ingresen datos o almacenando credenciales de inicio de sesión y contraseña para que sus usuarios los autentiquen. Para ello disponemos de bases de datos.
Las bases de datos vienen en todas las formas y formas. Una base de datos de uso común en toda la web es la base de datos MySQL. Ha existido durante mucho tiempo y ha demostrado su posición y estabilidad más veces de las que puede contar.
En este ejemplo, nos sumergiremos en los fundamentos del acceso a la base de datos en Go, crearemos tablas de base de datos, almacenaremos datos y los recuperaremos nuevamente.

## Instalando el paquete go-sql-driver/mysql

El lenguaje de programación Go viene con un práctico paquete llamado `database/sql` para consultar todo tipo de bases de datos SQL. Esto es útil ya que abstrae todas las características comunes de SQL en una única API para su uso. Lo que Go no incluye son los controladores de base de datos. En Go, el controlador de la base de datos es un paquete que implementa los detalles de bajo nivel de una base de datos específica (en nuestro caso, MySQL). Como ya habrás adivinado, esto es útil para mantener la compatibilidad con versiones posteriores. Dado que, en el momento de crear todos los paquetes de Go, los autores no pueden prever que todas las bases de datos cobrarán vida en el futuro y respaldar todas las bases de datos posibles sería una gran cantidad de trabajo de mantenimiento.

Para instalar el controlador de la base de datos MySQL, vaya a su terminal de elección y ejecute:

```bash
go get -u github.com/go-sql-driver/mysql
```

## Conexión a una base de datos MySQL

Lo primero que debemos verificar después de instalar todos los paquetes necesarios es si podemos conectarnos a nuestra base de datos MySQL con éxito. Si aún no tiene un servidor de base de datos MySQL en ejecución, puede iniciar una nueva instancia con Docker fácilmente. Aquí están los documentos oficiales para la imagen MySQL de Docker:https://hub.docker.com/_/mysql

Para verificar si podemos conectarnos a nuestra base de datos, importe el paquete database/sql y el paquete go-sql-driver/mysql y abra una conexión como esta:

```go
import "database/sql"
import _ "go-sql-driver/mysql"

// Configure the database connection (always check errors)
db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

// Initialize the first connection to the database, to see if everything works correctly.
// Make sure to check the error.
err := db.Ping()
```

## Creando nuestra primera tabla de base de datos

Cada entrada de datos en nuestra base de datos se almacena en una tabla específica. Una tabla de base de datos consta de columnas y filas. Las columnas dan a cada entrada de datos una etiqueta y especifican el tipo de la misma. Las filas son los valores de datos insertados. En nuestro primer ejemplo, queremos crear una tabla como esta:

| id   | username   | password  | created_at          |
| ---- | ---------- | --------- | ------------------- |
| 1    | jhondoe    | secret    | 2023-01-07 22:30:00 |
| 2    | chicho69   | secret    | 2023-01-07 22:30:01 |
| 3    | lizhet11   | secret    | 2023-01-07 22:30:02 |

Traducido a SQL, el comando para crear la tabla se verá así:

```sql
CREATE TABLE users (
  id INT AUTO_INCREMENT,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at DATETIME,
  PRIMARY KEY (id)
);
```

Ahora que tenemos nuestro comando SQL, podemos usar el paquete `database/sql` para crear la tabla en nuestra base de datos MySQL:

```go
query := `
  CREATE TABLE users (
    id INT AUTO_INCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME,
    PRIMARY KEY (id)
  );
`

// Executes the SQL query in our database. Check err to ensure there was no error.
_, err := db.Exec(query)
```

## Insertando nuestro primer usuario

Si está familiarizado con SQL, insertar nuevos datos en nuestra tabla es tan fácil como crear nuestra tabla. Una cosa a tener en cuenta es que, de forma predeterminada, Go usa declaraciones preparadas para insertar datos dinámicos en nuestras consultas SQL, que es una forma de pasar de forma segura los datos proporcionados por el usuario a nuestra base de datos sin riesgo de daños. En los primeros días de la programación web, los programadores pasaban los datos directamente con la consulta a la base de datos, lo que causaba vulnerabilidades masivas y podía romper una aplicación web completa. Por favor, no hagas eso. Es fácil hacerlo bien.

Para insertar nuestro primer usuario en nuestra tabla de base de datos, creamos una consulta SQL como la siguiente. Como puede ver, omitimos la columna id, ya que MySQL la configura automáticamente. El signo de interrogación le dice al controlador SQL que son marcadores de posición para datos reales. Aquí es donde puede ver las declaraciones preparadas de las que hablamos.

```sql
INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)
```

Ahora podemos usar esta consulta SQL en Go e insertar una nueva fila en nuestra tabla:

```go
import "time"

username := "johndoe"
password := "secret"
createdAt := time.Now()

// Inserts our data into the users table and returns with the result and a possible error.
// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
```

Para obtener la identificación recién creada para su usuario, simplemente consígala así:

```go
userID, err := result.LastInsertId()
```

## Consultando nuestra tabla de usuarios

Ahora que tenemos un usuario en nuestra tabla, queremos consultarlo y recuperar toda su información. En Go tenemos dos posibilidades para consultar nuestras tablas. Hay `db.Query` que puede consultar varias filas, para que podamos iterar y hay `db.QueryRow` en caso de que solo queramos consultar una fila específica.

Consultar una fila específica funciona básicamente como cualquier otro comando SQL que hayamos cubierto antes.
Nuestro comando SQL para consultar a un solo usuario por su ID se ve así:

```sql
SELECT id, username, password, created_at FROM users WHERE id = ?
```

En Go, primero declaramos algunas variables para almacenar nuestros datos y luego consultamos una sola fila de la base de datos de esta manera:

```go
var (
  id        int
  username  string
  password  string
  createdAt time.Time
)

// Query the database and scan the values into out variables. Don't forget to check for errors.
query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
```

## Consultando a todos los usuarios

En la sección anterior, hemos cubierto cómo consultar una sola fila de usuario. Muchas aplicaciones tienen casos de uso en los que desea consultar a todos los usuarios existentes. Esto funciona de manera similar al ejemplo anterior, pero con un poco más de codificación involucrada.

Podemos usar el comando SQL del ejemplo anterior y recortar la cláusula `WHERE`. De esta manera, consultamos a todos los usuarios existentes.

```sql
SELECT id, username, password, created_at FROM users
```

En Go, primero declaramos algunas variables para almacenar nuestros datos y luego consultamos una sola fila de la base de datos de esta manera:

```go
type user struct {
  id        int
  username  string
  password  string
  createdAt time.Time
}

rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
defer rows.Close()

var users []user
for rows.Next() {
  var u user
  err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
  users = append(users, u)
}
err := rows.Err() // check err
```

El segmento de usuarios ahora podría contener algo como esto:

```go
users {
  user {
    id:        1,
    username:  "johndoe",
    password:  "secret",
    createdAt: time.Time{
      wall: 0x0, 
      ext: 63701044325, 
      loc: (*time.Location)(nil)
    },
  },
  user {
    id:        2,
    username:  "alice",
    password:  "bob",
    createdAt: time.Time{
      wall: 0x0, 
      ext: 63701044622, 
      loc: (*time.Location)(nil)
    },
  },
}
```

## Eliminar un usuario de nuestra tabla

Finalmente, eliminar un usuario de nuestra tabla es tan sencillo como `.Exec` en las secciones anteriores:

```go
_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
```