package services

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/MoonHack2077/Parcial2-SO/config"
	"github.com/MoonHack2077/Parcial2-SO/models"
)

func TestCrearTarea(t *testing.T) {
	// Cargar .env
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("❌ No se pudo cargar el archivo .env")
	}

	// Verificar que la URI esté bien formada
	if os.Getenv("MONGO_URI_TEST") == "" {
		t.Fatal("❌ MONGO_URI está vacío o no cargado")
	}
	
	config.ConectarDB() // ← ESTA LÍNEA INICIALIZA LA CONEXIÓN

	tarea := models.Tarea{
		Titulo:      "Prueba",
		Descripcion: "Esto es una prueba",
		Completado:  false,
	}

	id, err := CrearTarea(tarea)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestObtenerTareas(t *testing.T) {
	tareas, err := ObtenerTareas()

	assert.NoError(t, err)
	assert.NotNil(t, tareas)
}

func TestActualizarTarea(t *testing.T) {
	// Crear una tarea para actualizar
	tarea := models.Tarea{
		Titulo:      "Tarea para actualizar",
		Descripcion: "Antes de actualizar",
		Completado:  false,
	}
	id, err := CrearTarea(tarea)
	assert.NoError(t, err)

	// Datos actualizados
	tareaActualizada := models.Tarea{
		Titulo:      "Tarea actualizada",
		Descripcion: "Después de actualizar",
		Completado:  true,
	}

	err = ActualizarTarea(id.Hex(), tareaActualizada)
	assert.NoError(t, err)
}

func TestEliminarTarea(t *testing.T) {
	// Crear una tarea para luego eliminar
	tarea := models.Tarea{
		Titulo:      "Tarea a eliminar",
		Descripcion: "Eliminar esta tarea en el test",
		Completado:  false,
	}
	id, err := CrearTarea(tarea)
	assert.NoError(t, err)

	// Eliminar la tarea
	err = EliminarTarea(id.Hex())
	assert.NoError(t, err)
}

