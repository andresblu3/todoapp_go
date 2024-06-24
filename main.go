package main

import (
	"log"
	"net/http"
	"todo-app/config"
	"todo-app/models"
	"todo-app/routers"

	"github.com/gorilla/handlers"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Todo{})

	router := routers.SetupRouter()

	// Configurar los encabezados de CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8080"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))
}
