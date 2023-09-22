package main

import (
	"html/template"
	"log"
	"net/http"
)

var partials = []string{
	"./cmd/web/templates/base.layout.gohtml",
	"./cmd/web/templates/header.partial.gohtml",
	"./cmd/web/templates/footer.partial.gohtml",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	log.Println("Starting front end service on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	templateSlice := append(
		make([]string, 0, len(partials)+1),
		append([]string{"./cmd/web/templates/" + t}, partials...)...,
	)

	tmpl, err := template.ParseFiles(templateSlice...)

	if err != nil {
		log.Printf("error parsing template: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("error executing template: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
