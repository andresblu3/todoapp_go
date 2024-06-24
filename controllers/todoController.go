package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	config.DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	config.DB.Create(&todo)
	json.NewEncoder(w).Encode(&todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	config.DB.First(&todo, params["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	config.DB.Save(&todo)
	json.NewEncoder(w).Encode(&todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	if err := config.DB.First(&todo, params["id"]).Error; err != nil {
		log.Printf("Error finding todo with id %s: %s", params["id"], err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	config.DB.Delete(&todo)
	w.WriteHeader(http.StatusNoContent)
}
