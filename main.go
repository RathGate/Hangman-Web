package main

import (
	"encoding/json"
	"fmt"
	"hangman-web/packages/hangman"
	"hangman-web/packages/session"
	"log"
	"net/http"
	"text/template"
)

var player session.Player
var gamedata hangman.HangManData

var data = session.SessionData{
	Player:   &player,
	GameData: &gamedata,
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// Handles the POST request if any (can be game reset, login or logout):
	if r.Method == http.MethodPost {
		if r.FormValue("reset") == "reset" {
			gamedata = hangman.HangManData{}
		} else if r.FormValue("username") != "" {
			player.Login(r.FormValue("username"))

		} else if r.FormValue("logout") != "" {
			player.Logout()
		} else if r.FormValue("getDifficultyData") != "" {
			var data []byte
			if player.Name != "" {
				switch r.FormValue("getDifficultyData") {
				case "easy":
					data, _ = json.Marshal(player.ModeEasy)
				case "hard":
					data, _ = json.Marshal(player.ModeHard)
				default:
					data, _ = json.Marshal(player.ModeMedium)
				}
				w.Write([]byte(data))
			} else {
				w.Write([]byte("404"))
			}
			return
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
		if r.FormValue("name") != "" && gamedata.Attempts > 0 {
			gamedata.RoundResult(r.FormValue("name"))
			if !gamedata.HasEnded {
				switch gamedata.CheckGameState() {
				case -1:
					player.LoseRound()
				case 1:
					player.WinRound()
				}
				gamedata.HasEnded = true
			}

		} else if r.FormValue("difficulty") != "" {

			if gamedata.FinalWord != "" && gamedata.Attempts < 10 {
				data.ResetGame()
			}

			switch r.FormValue("difficulty") {
			case "easy":
				gamedata.InitGame("words.txt")
			case "hard":
				gamedata.InitGame("words3.txt")
			default:
				gamedata.InitGame("words2.txt")
			}
			player.SwitchDifficulty(r.FormValue("difficulty"))
		}
	}

	// Tried to access the page when no game was launched:
	// redirects to the home page.
	if gamedata.FinalWord == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// Defines the files to parse and executes the template:
	files := []string{"templates/hangman.html", "templates/_scoreboard.html"}
	template := template.Must(template.ParseFiles(files...))
	template.Execute(w, data)
}

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
