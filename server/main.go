package main

import (
	"fantasy-map-server/handlers"
	"html/template"
	"net/http"
)

// Define a struct to hold template data if needed
type PageData struct {
	Title string
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./data"))
	http.Handle("/data/", http.StripPrefix("/data/", fs))

	// Serve the HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title: "Fantasy Map",
		}

		tmpl.Execute(w, data)
	})

	// Additional handlers...
	http.HandleFunc("/locations", handlers.HandleLocations)

	http.ListenAndServe(":8080", nil)
}
