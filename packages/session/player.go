package session

type Player struct {
	Name              string
	TotalGameCount    int
	IsLoggedIn        bool
	CurrentDifficulty *GameModeData
	ModeEasy          GameModeData
	ModeMedium        GameModeData
	ModeHard          GameModeData
}

type GameModeData struct {
	Name             string
	WinCount         int
	LoseCount        int
	MaxWinStreak     int
	CurrentWinStreak int
}

// *----------------------- LOGIN / LOGOUT -----------------------* //

func (p *Player) Login(name string) {
	if len(name) > 20 {
		name = "Nice try :)"
	}

	p.Name = name
	p.IsLoggedIn = true
	p.InitDifficulty()
}

func (p *Player) Logout() {
	*p = Player{}
}

// *----------------------- DIFFICULTY ------------------------* //

func (p *Player) InitDifficulty() {
	p.ModeEasy.Name = "easy"
	p.ModeMedium.Name = "medium"
	p.ModeHard.Name = "hard"
}

func (p *Player) SetPlayerDifficulty(difficulty string) {
	if !p.IsLoggedIn {
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

// *--------------------- ROUND RESULTS ---------------------* //

func (p *Player) WinRound() {
	if !p.IsLoggedIn {
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
	if !p.IsLoggedIn {
		return
	}

	p.TotalGameCount++
	p.CurrentDifficulty.LoseCount++
	p.CurrentDifficulty.CurrentWinStreak = 0
}
