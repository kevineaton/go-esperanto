//This application runs a very basic Gin server that will randomly give an Esperanto phrase upon request or allow
//authenticated users to add new words to the phrase book
package main

import (
	"fmt"
	"net/http"
	"os"
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

//For now we use a global variable to hold all of the phrases in memory
var phrases []Pair

//main begins the program
func main() {
	fmt.Println("Started, configuring phrasebook and authentication")
	config = loadConfig()
	phrases = loadPhrasebook()

	r := setupRouter()
	err := http.ListenAndServe(config.Port, r)
	if err != nil {
		fmt.Printf("\n\tERROR: %+v\n", err.Error())
	}
	os.Exit(1)
}
