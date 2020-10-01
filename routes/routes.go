package routes

import (
	"awesomeProject1/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {

	router.HandleFunc("/api/todos", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/api/todos/{id}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/api/todos/create", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/update/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todos/delete/{id}", controllers.RemoveTodo).Methods("DELETE")
}
