package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/MoonHack2077/Parcial2-SO-crear/config"
	"github.com/MoonHack2077/Parcial2-SO-crear/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CrearTarea(tarea models.Tarea) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tarea.CreadoEn = time.Now()
	collection := config.GetCollection("tareas")
	result, err := collection.InsertOne(ctx, tarea)
	if err != nil {
		fmt.Println("Error al insertar tarea:", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func ObtenerTodasTareas() ([]models.Tarea, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetCollection("tareas")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tareas []models.Tarea
	for cursor.Next(ctx) {
		var tarea models.Tarea
		if err := cursor.Decode(&tarea); err != nil {
			return nil, err
		}
		tareas = append(tareas, tarea)
	}

	return tareas, nil
}
