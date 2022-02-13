package main

import (
	"github.com/gorilla/mux"
	"github.com/orutrax/go-crud-postgres-api/commons"
	"github.com/orutrax/go-crud-postgres-api/routes"
	"log"
	"net/http"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetMovementsRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Server is running in the port => 9000")
	log.Println(server.ListenAndServe())
}
