package routes

import (
	"gostorage/controllers"
	"gostorage/middlewares"

	"github.com/gorilla/mux"
)

func storageRoutes(r *mux.Router) {
	controller := controllers.NewStorageController()

	r.Use(middlewares.IsAuthenticated)
	r.HandleFunc("/upload", controller.Upload).Methods("POST")
}
