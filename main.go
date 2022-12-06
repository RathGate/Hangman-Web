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

// Player struct stores all the information linked to the current player if logged in.
// HangmanData struct stores all the information linked to the current game played, if there's one.
var player session.Player
var gamedata hangman.HangmanData

// SessionData struct stores both instances above in order to send them to the templates.
var data = session.SessionData{
	Player:   &player,
	GameData: &gamedata,
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	gamedata.CheckRoundState()

	// Handles the POST request if any (can be game reset, login or logout):
	if r.Method == "POST" {

		if r.FormValue("reset") == "reset" {
			gamedata.EmptyData()
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// Defines the files to parse and executes the template:
	files := []string{"templates/index.html", "templates/_scoreboard.html"}
	template := template.Must(template.ParseFiles(files...))
	template.Execute(w, data)
}

func hangmanHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		// *PLAYER HAS CHOSEN A DIFFICULTY AND CLICKED START → Launches the game:
		if r.FormValue("difficulty") != "" {
			// Player tried to launch a game when there was already one;
			// if the player has already lost at least one point, loses the current game:
			if gamedata.FinalWord != "" && gamedata.Attempts < 10 {
				data.ResetGame()
			}

			// Launches the game with the appropriate dictionary file;
			//
			gamedata.InitGame(r.FormValue("difficulty"))
			// If the player is logged in, changes their current difficulty in their settings:
			player.SwitchDifficulty(r.FormValue("difficulty"))

			// *PLAYER FILLED THE SUGGESTION INPUT → Launches the round:
		} else if r.FormValue("suggestion") != "" && gamedata.Attempts > 0 {

			// Plays the round:
			result := gamedata.RoundResult(r.FormValue("suggestion"))
			// Wins or loses the game, if already won/lost, does nothing.
			if result != 0 && !gamedata.HasEnded {
				if result == 1 {
					player.WinRound()
				} else {
					player.LoseRound()

				}
				gamedata.HasEnded = true
			}
		}
		http.Redirect(w, r, "/hangman", http.StatusSeeOther)
	}

	// *REDIRECTS TO HOMEPAGE IF PAGE ACCESSED WITH NO PENDING GAME:
	if gamedata.FinalWord == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// ?PARSES THE FILES NEEDED AND EXECUTES THE TEMPLATES:
	files := []string{"templates/hangman.html", "templates/_scoreboard.html", "templates/_game.html", "templates/_results.html"}
	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, data)
}

func main() {
	// Handles css files:
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles routing:
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hangman", hangmanHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
