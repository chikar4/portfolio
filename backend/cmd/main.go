package main

import (
	"fmt"
	"lien/backend/handlers"
	"net/http"
)

const port = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/mainPage", handlers.MainPage)

	fmt.Println("Serveur ouvert sur http://localhost:8080")

	http.ListenAndServe(port, nil)

}
