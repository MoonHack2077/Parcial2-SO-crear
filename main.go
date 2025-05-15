package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MoonHack2077/Parcial2-SO-crear/config"
	"github.com/MoonHack2077/Parcial2-SO-crear/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar .env solo en entorno local
	if os.Getenv("ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  No se pudo cargar .env localmente, usando variables de entorno del sistema")
		}
	}

	// Conectar a MongoDB
	config.ConectarDB()

	// Crear router
	router := mux.NewRouter()

	// Registrar endpoints
	router.HandleFunc("/tareas", controllers.CrearTarea).Methods("POST")
	router.HandleFunc("/tareas", controllers.ObtenerTareas).Methods("GET")

	// Ruta raíz de prueba
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Microservicio de Tareas 🚀")
	}).Methods("GET")

	// Correr servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🟢 Servidor corriendo en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
