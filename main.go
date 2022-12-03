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
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles default route:
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("templates/hangman.html"))
		if r.Method == http.MethodPost {
			if r.FormValue("name") != "" && data.Attempts > 0 {
				data.RoundResult(r.FormValue("name"))
			} else if r.FormValue("difficulty") != "" {
				switch r.FormValue("difficulty") {
				case "easy":
					data.InitGame("words.txt")
				case "hard":
					data.InitGame("words3.txt")
				default:
					data.InitGame("words2.txt")
				}
			}
		}
		if data.FinalWord == "" {
			data.InitGame("words2.txt")
		}
		template.Execute(w, data)

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if r.FormValue("reset") == "true" {
				data = hangman.HangManData{}
				w.Write([]byte("ok !!"))
				return
			}
		}
		template := template.Must(template.ParseFiles("templates/index.html"))
		template.Execute(w, data)
	})

	// Launches the server:
	preferredPort := ":80"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
