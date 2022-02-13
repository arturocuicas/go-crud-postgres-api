package routes

import (
	"github.com/gorilla/mux"
	"github.com/orutrax/go-crud-postgres-api/controllers"
)

func SetMovementsRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/movements", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/movements", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/movements/{uuid}", controllers.Get).Methods("GET")
	subRoute.HandleFunc("/movements/{uuid}", controllers.Delete).Methods("DELETE")
}
