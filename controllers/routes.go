package controllers

import (
	"github.com/gorilla/mux"
	"github.com/Arthelon/n10n/controllers/api"
	"net/http"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	authRouter := router.PathPrefix("/auth").Subrouter()

	apiRouter.HandleFunc("/", api.API_Verify).Methods("GET")
	apiRouter.HandleFunc("/verify", Use(api.API_Verify, RequireUserToken)).Methods("GET")

	authRouter.HandleFunc("/login", AUTH_Post_Login).Methods("POST")
	authRouter.HandleFunc("/register", AUTH_Post_Register).Methods("POST")
	return router
}

func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
