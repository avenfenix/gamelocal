package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PartidaGuardada struct {
	Juego  string `json:"juego"`
	Estado string `json:"estado"`
}

func main() {
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
