package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Publication struct {
	Title    string `json:"title"`
	Authors  string `json:"authors"`
	Link     string `json:"link"`
	Abstract string `json:"abstract"`
	Journal  string `json:"journal"`
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func loadPublications() []Publication {
	file, err := os.Open("data/publications.json")
	if err != nil {
		return []Publication{}
	}
	defer file.Close()

	var pubs []Publication
	json.NewDecoder(file).Decode(&pubs)

	return pubs
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Publications []Publication
	}{
		Publications: loadPublications(),
	}

	templates.ExecuteTemplate(w, "index.html", data)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
