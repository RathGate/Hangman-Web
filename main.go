package main

import (
	"fmt"
	"hangman-web/packages/hangman"
	"hangman-web/packages/session"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// Handles the POST request if any (can be game reset, login or logout):
	if r.Method == http.MethodPost {
		if r.FormValue("reset") == "reset" {
			fmt.Print("pio")
			data.GameData = hangman.HangManData{}
		} else if r.FormValue("username") != "" {
			data.Player.Login(r.FormValue("username"))

		} else if r.FormValue("logout") != "" {
			data.Player.Logout()
		}
	}

	// Defines the files to parse and executes the template:
	files := []string{"templates/index.html", "templates/_scoreboard.html"}
	template := template.Must(template.ParseFiles(files...))
	template.Execute(w, data)
}

func hangmanHandler(w http.ResponseWriter, r *http.Request) {

	// Game can only be launched via the main page == via its form:
	if r.Method == http.MethodPost {
		if r.FormValue("name") != "" && data.GameData.Attempts > 0 {
			data.GameData.RoundResult(r.FormValue("name"))
			if data.GameData.Won == 0 {
				switch data.GameData.CheckGameState() {
				case -1:
					data.Player.LoseRound()
					data.GameData.Won = -1
				case 1:
					data.Player.WinRound()
					data.GameData.Won = 1
				}
			}

		} else if r.FormValue("difficulty") != "" {

			if data.GameData.FinalWord != "" && data.GameData.Attempts < 10 {
				data.ResetGame()
			}

			switch r.FormValue("difficulty") {
			case "easy":
				data.GameData.InitGame("words.txt")
			case "hard":
				data.GameData.InitGame("words3.txt")
			default:
				data.GameData.InitGame("words2.txt")
			}
			data.Player.SwitchDifficulty(r.FormValue("difficulty"))
		}
	}

	// Tried to access the page when no game was launched:
	// redirects to the home page.
	if data.GameData.FinalWord == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// Defines the files to parse and executes the template:
	files := []string{"templates/hangman.html", "templates/_scoreboard.html"}
	template := template.Must(template.ParseFiles(files...))
	template.Execute(w, data)
}

var data session.SessionData

func main() {
	// Handles css files:
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles routing:
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hangman", hangmanHandler)

	// Launches the server:
	preferredPort := ":6969"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
