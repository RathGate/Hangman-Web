package main

import (
	"fmt"
	"hangman-web/packages/hangman"
	"log"
	"net/http"
	"text/template"
)

var data hangman.HangManData

func main() {
	// Handles css files:
	cssHandler := http.FileServer(http.Dir("assets/css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	jsHandler := http.FileServer(http.Dir("assets/js"))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandler))

	// Handles default route:
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		if data.FinalWord == "" {
			data.InitGame("words.txt")
		}
		template := template.Must(template.ParseFiles("assets/templates/hangman.html"))
		if r.Method == http.MethodPost {
			data.RoundResult(r.FormValue("name"))
		}
		template.Execute(w, data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println(r.FormValue("reset"))
			data = hangman.HangManData{}
		}
		template := template.Must(template.ParseFiles("index.html"))
		template.Execute(w, data)
	})

	// Launches the server:
	preferredPort := ":80"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
