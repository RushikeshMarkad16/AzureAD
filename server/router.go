package server

import (
	"net/http"

	"github.com/RushikeshMarkad16/AzureAD/handler"
	"github.com/gorilla/mux"
)

// InitRouter ...
func InitRouter() (router *mux.Router) {

	router = mux.NewRouter()

	router.HandleFunc("/login", handler.HandleLandingPage).Methods(http.MethodGet)
	router.HandleFunc("/auth-callback", handler.HandleCallback).Methods(http.MethodGet)

	return
}
