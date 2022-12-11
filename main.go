package main

import (
	"fmt"
	"hangman-web/packages/hangman"
	"hangman-web/packages/session"
	"log"
	"net/http"
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
