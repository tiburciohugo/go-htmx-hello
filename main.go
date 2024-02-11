package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Year     int
	Director string
}

type Actor struct {
	Name    string
	Birth   string
	Country string
}

type Data struct {
	Films  []Film
	Series []Film
	Actors []Actor
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		data := Data{
			Films: []Film{
				{Title: "The Shawshank Redemption", Year: 1994, Director: "Frank Darabont"},
				{Title: "The Godfather", Year: 1972, Director: "Francis Ford Coppola"},
				{Title: "The Dark Knight", Year: 2008, Director: "Christopher Nolan"},
				{Title: "Blade Runner", Year: 1982, Director: "Ridley Scott"},
			},
			Series: []Film{
				{Title: "Breaking Bad", Year: 2008, Director: "Vince Gilligan"},
				{Title: "Game of Thrones", Year: 2011, Director: "David Benioff"},
				{Title: "The Crown", Year: 2016, Director: "Peter Morgan"},
			},
			Actors: []Actor{
				{Name: "Tom Hanks", Birth: "9 July 1956", Country: "USA"},
				{Name: "Morgan Freeman", Birth: "1 June 1937", Country: "USA"},
				{Name: "Al Pacino", Birth: "25 April 1940", Country: "USA"},
				{Name: "Christian Bale", Birth: "30 January 1974", Country: "UK"},
			},
		}

		tmpl.Execute(w, data)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
