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
		main_templ, _ := template.New("main").Parse("Bonjour !")
		template.Execute(w, main_templ)
	})

	// Launches the server:
	fmt.Printf("Starting server at port 80\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
