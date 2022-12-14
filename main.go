package main

import (
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
	Player:            &player,
	GameData:          &gamedata,
	CurrentDifficulty: "medium",
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// Clears gamedata if it isn't empty and the result screen has already been displayed:
	// (== no need to display the resume/quit screen on a finished game)
	if gamedata.HasEnded {
		gamedata.EmptyData()
	}

	if r.Method == "POST" {
		if r.FormValue("reset") == "reset" {
			data.ResetGame()

		} else if r.FormValue("username") != "" {
			player.Login(r.FormValue("username"))
			player.SetPlayerDifficulty(data.CurrentDifficulty)

		} else if r.FormValue("logout") != "" {
			player.Logout()

			// Sent by JS in order to refresh the displayed player data on the main screen.
		} else if requestedDifficulty := r.FormValue("requestedData"); requestedDifficulty != "" {
			w.Write(player.GetGameModeData(requestedDifficulty))
			// Early return because HTML page is updated via JS, there's no need to reload.
			return
		}

		// Reloads the page with redirection, used to empty the values of the previous POST request.
		// No real function to reload the page other than using JS or redirecting.
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// Parses the files needed and executes the template.
	files := []string{"templates/index.html", "templates/_scoreboard.html"}
	template := template.Must(template.ParseFiles(files...))
	template.Execute(w, data)
}

func hangmanHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		if difficulty := r.FormValue("difficulty"); difficulty != "" {
			if difficulty == "current" {
				difficulty = data.CurrentDifficulty
			}
			data.StartNewGame(difficulty)

		} else if suggestion := r.FormValue("suggestion"); suggestion != "" && gamedata.Attempts > 0 {
			// Plays the round only if at least 1 attempt left.
			data.PlayRound(suggestion)
		}

		// Reloads the page with redirection, used to empty the values of the previous POST request.
		// No real function to reload the page other than using JS or redirecting.
		http.Redirect(w, r, "/hangman", http.StatusSeeOther)
	}

	// REDIRECTS TO HOMEPAGE IF PAGE ACCESSED WITH NO PENDING GAME:
	if gamedata.FinalWord == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// Parses the files needed and executes the template.
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
