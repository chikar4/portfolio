package routes

import (
	"lien/backend/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	// * routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/mainPage", handlers.MainPage)
}
