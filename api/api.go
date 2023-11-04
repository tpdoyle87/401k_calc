package api

import (
	"401k_calculator/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupAPI() http.Handler {
	router := mux.NewRouter()
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST"})
	origins := handlers.AllowedOrigins([]string{"www.example.com"})
	router.HandleFunc(
		"/calculate",
		controllers.Calculate(),
	).Methods("POST")

	apiHandler := handlers.CORS(credentials, methods, origins)(router)

	return apiHandler
}
