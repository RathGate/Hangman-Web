package hangman

import (
	"hangman-web/packages/utils"
	"sort"
	"strings"
)

// Checks if the word had been entirely uncovered.
func (data HangmanData) CheckGameState() int {
	if data.Word == data.FinalWord && data.Attempts > 0 {
		return 1
	} else if data.Attempts == 0 {
		return -1
	}
	return 0
}

func (data *HangmanData) RevealLetter(answer string) {
	for i, letter := range data.FinalWord {
		if answer == string(letter) {
			data.Word = data.Word[:i] + string(data.FinalWord[i]) + data.Word[i+1:]
		}
	}
}

// Returns false if letter already suggested, else returns true.
func (data *HangmanData) AddUsedLetters(letter string) bool {
	for _, char := range data.UsedLetters {
		if char == letter {
			return false
		}
	}
	data.UsedLetters = append(data.UsedLetters, letter)
	sort.Strings(data.UsedLetters)
	data.UsedToString = strings.Join(data.UsedLetters, " ")
	return true
}

func (data *HangmanData) LosePoints(points int) {
	data.Attempts -= points
	if data.Attempts < 0 {
		data.Attempts = 0
	}
}

// Round results initiates the round.
// Returns -1 if the game is lost, 1 if it's won, else 0.
func (data *HangmanData) RoundResult(answer string) int {
	answer = strings.ToUpper(answer)

	// Answer is invalid:
	if !utils.IsAlpha(answer) {
		return 0

		// Answer is one-char long:
	} else if len(answer) == 1 {
		if !data.AddUsedLetters(answer) {
			return 0
		}
		if strings.Contains(data.FinalWord, string(answer[0])) {
			data.RevealLetter(answer)
		} else {
			data.LosePoints(1)
		}
		// Answer = at least 2 characters.
	} else {
		if answer == data.FinalWord {
			data.Word = data.FinalWord
			return 1
		} else {
			data.LosePoints(2)
		}
	}
	return data.CheckGameState()
}
