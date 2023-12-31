package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis"},
				{Title: "Blade Runner", Director: "Ridley"},
				{Title: "The thing", Director: "John"},
			},
		}

		tmpl.Execute(w, films)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
