package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Define los tipos de datos en el esquema
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Define las consultas disponibles
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Aquí puedes obtener los datos del usuario usando el ID proporcionado en los argumentos
					// y devolver los datos en un mapa
					return map[string]interface{}{
						"id":    "123",
						"name":  "John Doe",
						"email": "johnddoe@example.com",
					}, nil
				},
			},
		},
	},
)

func main() {
	// Crea un nuevo esquema a partir de los tipos de datos y las consultas
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	if err != nil {
		fmt.Printf("Error al crear el esquema: %v", err)
		return
	}

	// Crea un endpoint HTTP para el servidor GraphQL
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		// Obtiene la consulta del cuerpo de la solicitud
		query := r.URL.Query().Get("query")
		if query == "" {
			query = "{}"
		}

		// Ejecuta la consulta con el esquema y devuelve el resultado en formato JSON
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: query,
		})

		if len(result.Errors) > 0 {
			fmt.Printf("Error en la consulta: %v", result.Errors)
		}

		json.NewEncoder(w).Encode(result)
	})

	// Inicia el servidor
	fmt.Println("Iniciando el servidor en http://localhost:8080/graphql")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %v", err)
		return
	}
}
