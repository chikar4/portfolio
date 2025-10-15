package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./home.html")
	if err != nil {
		fmt.Println("La page home n'arrive pas a chargée")
	}

	tmpl.Execute(w, nil)
}

func MainPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./mainPage.html")
	if err != nil {
		fmt.Println("La page MainPage n'arrive pas a chargée")
	}

	tmpl.Execute(w, nil)
}
