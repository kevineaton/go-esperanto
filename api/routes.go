package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// GetAllPhrasesRoute gets all the phrases for the system
func GetAllPhrasesRoute(w http.ResponseWriter, r *http.Request) {
	isValidToken := checkAuthenticatedHelper(&w, r)
	if !isValidToken {
		sendError(w, http.StatusForbidden, "bad token or no token found; ensure it is passed in the X-API-TOKEN header")
		return
	}

	sendResponse(w, http.StatusOK, phrases)
}

// GetRandomPhraseRoute gets a random phrase
func GetRandomPhraseRoute(w http.ResponseWriter, r *http.Request) {
	isValidToken := checkAuthenticatedHelper(&w, r)
	if !isValidToken {
		sendError(w, http.StatusForbidden, "bad token or no token found; ensure it is passed in the X-API-TOKEN header")
		return
	}
	// use a seed to get better random number
	rand.Seed(time.Now().UnixNano())
	randID := rand.Intn(len(phrases))
	p := phrases[randID]
	sendResponse(w, http.StatusOK, p)
}

type apiReturn struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func sendError(w http.ResponseWriter, httpCode int, message string) {
	ret := apiReturn{
		Message: message,
		Data:    nil,
	}

	data, _ := json.Marshal(ret)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(data)
}

func sendResponse(w http.ResponseWriter, httpCode int, payload interface{}) {
	ret := apiReturn{
		Data: payload,
	}
	response, _ := json.Marshal(ret)
	// add new line for nicer terminal output
	response = append(response, []byte("\n")[0])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}
