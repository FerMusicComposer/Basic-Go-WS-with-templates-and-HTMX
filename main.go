package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Go Web Service 1.0")
	fmt.Println("Listening on port 8080...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "The Dark Knight", Director: "Christopher Nolan"},
				{Title: "Schindler's List", Director: "Steven Spielberg"},
				{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
				{Title: "Inception", Director: "Christopher Nolan"},
				{Title: "Forrest Gump", Director: "Robert Zemeckis"},
			},
		}
		tmpl.Execute(w, films)
	})

	http.HandleFunc("/add-film", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
