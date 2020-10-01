package controllers

import (
	"awesomeProject1/database"
	"awesomeProject1/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var todos []models.Todo

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET requests allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := database.DB.Query("SELECT * FROM todos")
	if err != nil {
		http.Error(w, "Unable to query DB", http.StatusBadRequest)
		return
	}
	defer rows.Close()

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
	id, err := strconv.Atoi(mux.Vars(r)["id"])
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

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The data sent was bad"))
		return
	}

	stmt, err := database.DB.Prepare("INSERT INTO todos (name) VALUES ($1)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(todo.Name); err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		log.Println("there was a problem encoding")
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := database.DB.Prepare("UPDATE todos SET name = $2 WHERE todos_id = $1")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(id, todo.Name); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-type", "application/json")
}

func RemoveTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	}

	stmt, err := database.DB.Prepare("DELETE FROM todos WHERE todos_id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		log.Println(err)
	}
}
