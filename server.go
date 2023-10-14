// Parte que sirve los archivos de juegos. Es la "base de datos" en linea (local)
// La parte cliente le envia los archivos de juego actualizados
//

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// Importa las librerías necesarias
)

type Usuario struct {
	Nombre      string `json:"nombre"`
	Dispositivo string `json:"dispositivo"`
	Juego       string `json:"juego"`
	Directorio  string `json:"directorio"`
}

func main() {
	http.HandleFunc("/guardar-partida", func(w http.ResponseWriter, r *http.Request) {
		// Maneja la solicitud para guardar la partida en el servidor
		// Guarda la partida en una base de datos o un archivo
		// ...

		fmt.Println("Partida guardada recibida desde el cliente Gamelocal")
	})

	http.ListenAndServe(":8080", nil) // Escucha en el puerto 8080

	var usuariosConectados []Usuario

	http.HandleFunc("/conectar-usuario", func(w http.ResponseWriter, r *http.Request) {
		var usuario Usuario
		err := json.NewDecoder(r.Body).Decode(&usuario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Verifica si el usuario y el juego son válidos (por ejemplo, GTA San Andreas)
		// Asocia el juego con el directorio del usuario
		// ...

		usuariosConectados = append(usuariosConectados, usuario)
		fmt.Println("Usuario conectado:", usuario.Nombre)
	})

	http.HandleFunc("/verificar-archivos", func(w http.ResponseWriter, r *http.Request) {
		// Recibe la solicitud de verificación de archivos del juego
		// Verifica la integridad de los archivos en el directorio del usuario
		// ...

		// Devuelve una respuesta al cliente (indicando si los archivos están actualizados o no)
		// ...
	})

}
