package session

import (
	"fmt"
	"hangman-web/packages/hangman"
)

type Player struct {
	Name              string
	TotalGameCount    int
	CurrentDifficulty *GameModeData
	ModeEasy          GameModeData
	ModeMedium        GameModeData
	ModeHard          GameModeData
}

func (p *Player) InitDifficulty() {
	p.ModeEasy.Name = "easy"
	p.ModeMedium.Name = "medium"
	p.ModeHard.Name = "hard"
	p.CurrentDifficulty = &p.ModeMedium
}

type GameModeData struct {
	Name             string
	WinCount         int
	LoseCount        int
	MaxWinStreak     int
	CurrentWinStreak int
}

type SessionData struct {
	Player   *Player
	GameData *hangman.HangmanData
}

func (p *Player) Login(name string) {
	p.Name = name
	p.InitDifficulty()
}

func (p *Player) Logout() {
	*p = Player{}
	fmt.Println(p)
}

func (p *Player) SwitchDifficulty(difficulty string) {
	if p.Name == "" {
		return
	}

	switch difficulty {
	case "easy":
		p.CurrentDifficulty = &p.ModeEasy

	case "hard":
		p.CurrentDifficulty = &p.ModeHard

	default:
		p.CurrentDifficulty = &p.ModeMedium
	}
}

func (p *Player) WinRound() {
	if p.Name == "" {
		return
	}

	p.TotalGameCount++
	p.CurrentDifficulty.WinCount++
	p.CurrentDifficulty.CurrentWinStreak++
	if p.CurrentDifficulty.CurrentWinStreak > p.CurrentDifficulty.MaxWinStreak {
		p.CurrentDifficulty.MaxWinStreak = p.CurrentDifficulty.CurrentWinStreak
	}
}

func (p *Player) LoseRound() {
	if p.Name == "" {
		return
	}

	p.TotalGameCount++
	p.CurrentDifficulty.LoseCount++
	p.CurrentDifficulty.CurrentWinStreak = 0
}

func (s *SessionData) ResetGame() {
	if s.GameData.HasEnded {
		s.GameData = &hangman.HangmanData{}
		return
	}
	if s.GameData.FinalWord != "" && s.GameData.Attempts < 10 {
		s.Player.LoseRound()

	}
	s.GameData = &hangman.HangmanData{}
}
