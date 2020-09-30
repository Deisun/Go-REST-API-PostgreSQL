package main

import (
	"awesomeProject1/database"
	"awesomeProject1/routes"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
)


func main() {
	database.ConnectDB()
	defer database.DB.Close()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
