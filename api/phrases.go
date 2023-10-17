package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// The Pair struct provides a simple Key/Value structure of an Esperanto phrase and an English translation
type Pair struct {
	Esperanto string `json:"esperanto"`
	English   string `json:"english"`
}

// Bind allows binding the struct for use in chi render
func (data *Pair) Bind(r *http.Request) error {
	return nil
}

// For now we use a global variable to hold all of the phrases in memory
var phrases []Pair

// LoadPhrasebook will load up the phrasebook.txt file, which is a | separated file with an Esperanto
// phrase and English translation on each line
func LoadPhrasebook() []Pair {
	parsedPhrases := []Pair{}
	pbFile := config.PhraseBookLocation + "phrasebook.txt"

	content, err := os.ReadFile(pbFile)

	if err != nil {
		// since there could be some weird pathing, also check in the current directory
		content, err = os.ReadFile("./phrasebook.txt")
		if err != nil {
			msg := fmt.Sprintf("Cannot load phrase book, looking in %s", pbFile)
			panic(msg)
		}
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		components := strings.Split(string(line), "|")
		if len(components) != 2 {
			continue
		}
		p := Pair{components[0], components[1]}
		parsedPhrases = append(parsedPhrases, p)
	}
	phrases = parsedPhrases
	return parsedPhrases
}
