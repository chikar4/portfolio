package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("frontend/templates/home.html")
	if err != nil {
		fmt.Println("La page home n'arrive pas a chargée")
		log.Printf("fail to parse tmpl %v", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("fail to execute tmpl %v", err)
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("frontend/templates/mainPage.html")
	if err != nil {
		fmt.Println("La page MainPage n'arrive pas a chargée")
		log.Printf("fail to parse tmpl %v", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("fail to execute tmpl %v", err)
	}
}
