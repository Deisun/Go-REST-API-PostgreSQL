package routes

import (
	"awesomeProject1/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {

	router.HandleFunc("/api/todos", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/api/todos/{id}", controllers.GetTodo).Methods("GET")
}
