package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PartidaGuardada struct {
	Juego  string `json:"juego"`
	Estado string `json:"estado"`
}

var partidasGuardadas []PartidaGuardada

func main() {
	http.HandleFunc("/guardar-partida", func(w http.ResponseWriter, r *http.Request) {
		var partida PartidaGuardada
		err := json.NewDecoder(r.Body).Decode(&partida)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		partidasGuardadas = append(partidasGuardadas, partida)
		fmt.Println("Partida guardada:", partida.Juego)

		w.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("/partidas-guardadas", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(partidasGuardadas)
	})

	http.ListenAndServe(":8080", nil)
}
