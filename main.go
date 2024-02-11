package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Year     int
	Director string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := []Film{
			{Title: "The Shawshank Redemption", Year: 1994, Director: "Frank Darabont"},
			{Title: "The Godfather", Year: 1972, Director: "Francis Ford Coppola"},
			{Title: "The Dark Knight", Year: 2008, Director: "Christopher Nolan"},
			{Title: "Blade Ruuner", Year: 1982, Director: "Ridley Scott"},
		}
		data := map[string][]Film{
			"Films": films,
			"Series": {
				{Title: "Breaking Bad", Year: 2008, Director: "Vince Gilligan"},
				{Title: "Game of Thrones", Year: 2011, Director: "David Benioff"},
				{Title: "The Crown", Year: 2016, Director: "Peter Morgan"},
			},
		}
		// log each film
		for _, film := range films {
			fmt.Println(film.Title, film.Year, film.Director)
		}
		tmpl.Execute(w, data)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
