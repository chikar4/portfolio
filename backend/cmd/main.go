package main

import (
	"fmt"
	"lien/backend/routes"
	"net/http"
)

const port = ":8080"

func main() {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("frontend/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	routes.RegisterRoutes(mux)

	fmt.Println("Serveur ouvert sur http://localhost:8080")

	http.ListenAndServe(port, nil)

}
