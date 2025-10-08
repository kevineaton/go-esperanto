package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPhrases(t *testing.T) {
	LoadConfig()
	assert.True(t, len(phrases) > 0)
	pair := phrases[0]
	assert.NotEqual(t, pair.English, "")
	assert.NotEqual(t, pair.Esperanto, "")
}

func TestGetAllPhrasesRoute(t *testing.T) {
	LoadConfig()
	code, body, err := TestEndpoint(http.MethodGet, "/", nil, GetAllPhrasesRoute, true)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)
	pair := []Pair{}
	err = testEndpointResultToStruct(body, &pair)
	assert.Nil(t, err)
	assert.NotZero(t, len(pair))

	// try a bad result
	code, _, _ = TestEndpoint(http.MethodGet, "/", nil, GetAllPhrasesRoute, false)
	assert.Equal(t, http.StatusForbidden, code)
}

func TestGetRandomPhraseRoute(t *testing.T) {
	LoadConfig()
	code, body, err := TestEndpoint(http.MethodGet, "/random", nil, GetRandomPhraseRoute, true)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)
	p1 := &Pair{}
	err = testEndpointResultToStruct(body, &p1)
	assert.Nil(t, err)
	assert.NotEqual(t, "", p1.English)
	assert.NotEqual(t, "", p1.Esperanto)

	// get another one, make sure it's not the same one
	code, body, err = TestEndpoint(http.MethodGet, "/random", nil, GetRandomPhraseRoute, true)
	assert.Equal(t, http.StatusOK, code)
	p2 := &Pair{}
	err = testEndpointResultToStruct(body, &p2)
	assert.Nil(t, err)
	assert.NotEqual(t, "", p2.English)
	assert.NotEqual(t, "", p2.Esperanto)

	// try without auth
	code, _, _ = TestEndpoint(http.MethodGet, "/random", nil, GetRandomPhraseRoute, false)
	assert.Equal(t, http.StatusForbidden, code)
}
