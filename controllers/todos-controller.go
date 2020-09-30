package controllers

import (
	"awesomeProject1/database"
	"awesomeProject1/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
