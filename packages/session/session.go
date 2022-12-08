package session

import (
	"encoding/json"
	"hangman-web/packages/hangman"
)

type SessionData struct {
	Player            *Player
	GameData          *hangman.HangmanData
	CurrentDifficulty string
}

// *------------------------ START / RESET -----------------------* //

func (s *SessionData) StartNewGame(difficulty string) {
	// Deletes current data if there's already a pending game.
	s.ResetGame()

	// Automatically sets difficulty to "medium" if neither "hard" or "easy".
	// Prevents invalid values from being passed.
	if difficulty != "hard" && difficulty != "easy" {
		difficulty = "medium"
	}

	// Sets overall and player's current difficulty values if needed.
	s.CurrentDifficulty = difficulty
	if s.Player.IsLoggedIn {
		s.Player.SetPlayerDifficulty(difficulty)
	}

	// Launches the game.
	s.GameData.InitGame(difficulty)
}

// TODO: must be factorized
func (s *SessionData) ResetGame() {
	if s.GameData.HasEnded {
		s.GameData.EmptyData()
		return
	}
	if s.GameData.FinalWord != "" && s.GameData.Attempts < 10 {
		s.Player.LoseRound()

	}
	s.GameData.EmptyData()
}

func (s *SessionData) PlayRound(suggested string) {
	// Round result: -1 equals game lost, 1 equals game won.
	result := s.GameData.RoundResult(suggested)
	if result == 0 || s.GameData.HasEnded {
		return
	}

	if result == 1 {
		s.Player.WinRound()
	} else {
		s.Player.LoseRound()
	}

	s.GameData.HasEnded = true
}

// *------------------ RESPONSES TO JAVASCRIPT -------------------* //

// Get player's GameMode data depending on the requested difficulty.
// Returns a response in []byte that shall be sent back to JS.
func (p Player) GetGameModeData(request string) (data []byte) {
	if !p.IsLoggedIn {
		return []byte("404")
	}

	switch request {
	case "easy":
		data, _ = json.Marshal(p.ModeEasy)
	case "hard":
		data, _ = json.Marshal(p.ModeHard)
	default:
		data, _ = json.Marshal(p.ModeMedium)
	}
	return data
}
