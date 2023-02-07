// Importa las dependencias necesarias
const axios = require("axios");

// Define la consulta GraphQL
const query = `
	query GetUser($id: ID!) {
		user(id: $id) {
			id
			name
			email
		}
	}
`;

// Define los parámetros de la consulta
const variables = {
	id: "1",
};

// Realiza una solicitud POST a la URL del servidor GraphQL con la consulta y los parámetros
axios.post("http://localhost:8080/graphql", {
	query,
	variables,
})
	.then(response => {
		// Imprime el resultado de la consulta
		console.log(response.data.data.user);
	})
	.catch(error => {
		// Imprime cualquier error que haya ocurrido
		console.error(error);
	});
