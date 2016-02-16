//This application runs a very basic Gin server that will randomly give an Esperanto phrase upon request or allow
//authenticated users to add new words to the phrase book
package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//The Pair struct provides a simple Key/Value structure of an Esperanto phrase and an English translation
type Pair struct {
	Esperanto string `json:"esperanto"`
	English   string `json:"english"`
}

//For now we use a global variable to hold all of the phrases in memory
var phrases []Pair

//The Config struct holds general configuration options for the application
type Config struct {
	AuthenticationToken string
	Port                string
}

//A global Config struct for use in bootstrapping and authenticating
var config Config

//main begins the program
func main() {
	fmt.Println("Started, configuring phrasebook and authentication")
	phrases = loadPhrasebook()
	config = loadConfig()

	//setup the routes
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		GetRandomPair(c, phrases)
	})
	router.POST("/", checkAuthentication(), func(c *gin.Context) {
		SaveNewPair(c, phrases)
	})

	router.Run(config.Port)
}

//loadConfig will load up a new configuration struct with sane defaults if none provided
func loadConfig() Config {
	config := Config{}
	config.AuthenticationToken = os.Getenv("GO_EO_AUTHTOKEN")
	config.Port = os.Getenv("GO_EO_API_PORT")

	if config.AuthenticationToken == "" {
		//randomize it with bcrypt on each server start up and prompt the user to specify one
		r1 := rand.Intn(100000000)
		r2 := rand.Intn(20000000)
		plain := fmt.Sprintf("%s-%d-%d", "go-esperanto", r1, r2)
		h := md5.New()
		h.Write([]byte(plain))
		code := string(fmt.Sprintf("%x", h.Sum(nil)))
		config.AuthenticationToken = code
	}

	if config.Port == "" {
		config.Port = "8081"
	}
	config.Port = fmt.Sprintf(":%s", config.Port)

	return config
}

//loadPhrasebook will load up the phrasebook.txt file, which is a | separated file with an Esperanto
//phrase and English translation on each line
func loadPhrasebook() []Pair {
    phrases := make([]Pair, 0)
	content, err := ioutil.ReadFile("./phrasebook.txt")
	if err != nil {
		panic("Cannot load phrasebook! Abandoning...")
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		components := strings.Split(string(line), "|")
		p := Pair{components[0], components[1]}
		phrases = append(phrases, p)
	}
    return phrases
}

//CheckAuthentication will check the user's authentication token in either the
//post body or the query string (in that priority order)
func checkAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
        if config.AuthenticationToken == "" {
            //somehow it wasn't setup (most likely in testing) so set it up
            config = loadConfig()
        }
		passedAuth := c.DefaultPostForm("auth", "")
		if passedAuth == "" {
			passedAuth = c.DefaultQuery("auth", "")
		}
        
        if passedAuth == config.AuthenticationToken {
			c.Set("Authenticated", true)
		} else {
			c.Set("Authenticated", false)
		}
		c.Next()
	}
}

//GetRandomPair will get a random pair of Esperanto/English phrases and return them in JSON
//
// see the Pair struct for return format
func GetRandomPair(c *gin.Context, phrases []Pair) {
	r := rand.Intn(len(phrases))
	p := phrases[r]
	c.JSON(200, p)
}

//SaveNewPair will save a new pair of Esperanto to English definitions to the data store
func SaveNewPair(c *gin.Context, phrases []Pair) {
	if !c.MustGet("Authenticated").(bool) {
		c.JSON(401, gin.H{"status": "Unauthorized"})
	} else {
		eo := c.DefaultPostForm("esperanto", "")
		en := c.DefaultPostForm("english", "")
		if eo == "" || en == "" {
			c.JSON(400, gin.H{"status": "You must pass in both an 'esperanto' phrase and an 'english' phrase"})
		} else {
			pair := Pair{eo, en}
			file, err := os.OpenFile("./phrasebook.txt", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				c.JSON(500, gin.H{"status": "There was a problem saving that pair"})
				panic(err)
			}
			defer file.Close()

			writeString := fmt.Sprintf("\n%s|%s", pair.Esperanto, pair.English)
			if _, err = file.WriteString(writeString); err != nil {
				c.JSON(500, gin.H{"status": "There was a problem saving that pair"})
				panic(err)
			}
            _ = append(phrases, pair)
			c.JSON(200, pair)
		}
	}
}
