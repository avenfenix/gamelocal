package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PartidaGuardada struct {
	Juego  string `json:"juego"`
	Estado string `json:"estado"`
}

type Usuario struct {
	Nombre      string `json:"nombre"`
	Dispositivo string `json:"dispositivo"`
	Juego       string `json:"juego"`
	Directorio  string `json:"directorio"`
}

func main() {

	// Información del usuario y el juego
	usuario := Usuario{
		Nombre:      "avenfenix",
		Dispositivo: "mark2",
		Juego:       "GTA San Andreas",
		Directorio:  "C:\\Users\\mark2\\Documents\\GTA San Andreas User Files",
	}

	// Conecta al usuario al servidor
	conectarUsuario(usuario)

	// Verifica la integridad de los archivos del juego
	verificarArchivosJuego(usuario)

	// Cambia esta dirección a la del servidor Gamelocal en tu red local
	servidorURL := "http://localhost:8080/guardar-partida"

	// Simula una partida guardada
	partida := PartidaGuardada{
		Juego:  "EjemploJuego",
		Estado: "DatosDelEstado",
	}

	// Serializa la partida guardada a formato JSON
	partidaJSON, err := json.Marshal(partida)
	if err != nil {
		fmt.Println("Error al serializar la partida:", err)
		return
	}

	// Crea una solicitud HTTP POST para enviar la partida al servidor
	resp, err := http.Post(servidorURL, "application/json", bytes.NewBuffer(partidaJSON))
	if err != nil {
		fmt.Println("Error al enviar la partida:", err)
		return
	}
	defer resp.Body.Close()

	// Verifica la respuesta del servidor
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Partida guardada correctamente en el servidor Gamelocal")
	} else {
		fmt.Println("Error al guardar la partida en el servidor Gamelocal")
	}

	// Espera 2 segundos antes de salir (simula el cierre del juego)
	time.Sleep(2 * time.Second)
}

func conectarUsuario(usuario Usuario) {
	servidorURL := "http://localhost:8080/conectar-usuario"
	enviarDatosAlServidor(servidorURL, usuario)
}

func verificarArchivosJuego(usuario Usuario) {
	servidorURL := "http://localhost:8080/verificar-archivos"
	respuesta := enviarDatosAlServidor(servidorURL, usuario)
	fmt.Println(respuesta)
	// Manejar la respuesta del servidor sobre la verificación de archivos del juego
	// ...
}

func enviarDatosAlServidor(url string, datos interface{}) []byte {
	datosJSON, err := json.Marshal(datos)
	if err != nil {
		fmt.Println("Error al serializar datos:", err)
		return nil
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(datosJSON))
	if err != nil {
		fmt.Println("Error al enviar datos al servidor:", err)
		return nil
	}
	defer resp.Body.Close()

	respuesta, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor:", err)
		return nil
	}

	return respuesta
}
