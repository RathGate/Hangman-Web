package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	// Handles css files:
	static := http.FileServer(http.Dir("assets/css"))
	http.Handle("/css/", http.StripPrefix("/css/", static))

	// Handles default route:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))
		if r.Method == http.MethodPost {
			main_templ, _ := template.New("main").Parse(r.FormValue("name"))
			template.Execute(w, main_templ)
		} else {
			main_templ, _ := template.New("main").Parse("Bonjour !")
			template.Execute(w, main_templ)
		}
	})

	// Launches the server:
	preferredPort := ":80"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
