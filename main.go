// This application runs a very basic HTTP server that will randomly give an Esperanto phrase upon request or allow
// authenticated users to add new words to the phrase book
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kevineaton/go-esperanto/v2/api"
)

// main begins the program
func main() {
	fmt.Println("Started, configuring phrasebook and authentication")
	config := api.LoadConfig()

	r := api.SetupRouter()
	err := http.ListenAndServe(config.Port, r)
	if err != nil {
		fmt.Printf("\n\tERROR: %+v\n", err.Error())
	}
	os.Exit(1)
}
