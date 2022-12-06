package hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HangmanData struct {
	Word         string
	FinalWord    string
	Attempts     int
	UsedLetters  []string
	UsedToString string
	HasEnded     bool
}

func (data *HangmanData) InitGame(dictFile string) {
	rand.Seed(time.Now().UnixNano())

	switch dictFile {
	case "easy":
		dictFile = "words.txt"
	case "hard":
		dictFile = "words3.txt"
	default:
		dictFile = "words.txt"
	}

	data.FinalWord = RandomWord(ReadFile(dictFile))
	data.Word = strings.Repeat("_", len(data.FinalWord))
	n := len(data.FinalWord)/2 - 1
	for i := 0; i < n; {
		r := rand.Intn(len(data.FinalWord))
		if data.Word[r] != byte('_') {
			continue
		} else {
			data.Word = data.Word[:r] + string(data.FinalWord[r]) + data.Word[r+1:]
			i++
		}
	}
	data.Attempts = 10
}

func ReadFile(filename string) []string {
	file, err := os.Open("assets/dict/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()
	return lines
}

func RandomWord(lines []string) string {
	if len(lines) == 0 {
		log.Fatal("File is not valid.")
	}

	rand.Seed(time.Now().UnixNano())
	result := lines[rand.Intn(len(lines))]
	return strings.ToUpper(result)
}

func (data *HangmanData) EmptyData() {
	*data = HangmanData{}
}
