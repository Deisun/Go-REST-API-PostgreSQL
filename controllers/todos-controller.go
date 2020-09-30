package controllers

import (
	"awesomeProject1/database"
	"awesomeProject1/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetAllTodos(w http.ResponseWriter, req *http.Request) {

	rows, err := database.DB.Query("SELECT * FROM todos")
	if err != nil {
		fmt.Println("could not do DB QUERY")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {

		todo := models.Todo{}

		err = rows.Scan(&todo.ID, &todo.Name)
		if err != nil {
			log.Println("Could not do rows.Scan")

		}
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, "Problem encoding", http.StatusTeapot)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("Problem converting integer")
	}

	row := database.DB.QueryRow("SELECT * FROM todos WHERE todos_id=$1", id)

	var todo models.Todo
	err = row.Scan(&todo.ID, &todo.Name)
	if err != nil {
		log.Println("Problem with row scan")
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		http.Error(w, "Problem encoding", http.StatusTeapot)
	}
}
