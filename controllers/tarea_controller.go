package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MoonHack2077/Parcial2-SO/models"
	"github.com/MoonHack2077/Parcial2-SO/services"
)

func CrearTarea(w http.ResponseWriter, r *http.Request) {
	var tarea models.Tarea
	if err := json.NewDecoder(r.Body).Decode(&tarea); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	id, err := services.CrearTarea(tarea)
	if err != nil {
		http.Error(w, "No se pudo crear la tarea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id.Hex()})
}

func ObtenerTareas(w http.ResponseWriter, r *http.Request) {
	tareas, err := services.ObtenerTareas()
	if err != nil {
		http.Error(w, "Error al obtener las tareas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tareas)
}
