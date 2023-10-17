package api

import (
	"net/http"
	"testing"

	"github.com/mitchellh/mapstructure"
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
	code, body, _ := TestEndpoint(http.MethodGet, "/", nil, GetAllPhrasesRoute, true)
	assert.Equal(t, http.StatusOK, code)
	data, err := unmarshalSliceFromTestRoute(body)
	assert.Nil(t, err)
	assert.NotZero(t, len(data))

	// try a bad result
	code, _, _ = TestEndpoint(http.MethodGet, "/", nil, GetAllPhrasesRoute, false)
	assert.Equal(t, http.StatusForbidden, code)
}

func TestGetRandomPhraseRoute(t *testing.T) {
	LoadConfig()
	code, body, _ := TestEndpoint(http.MethodGet, "/random", nil, GetRandomPhraseRoute, true)
	assert.Equal(t, http.StatusOK, code)
	data, err := unmarshalMapFromTestRoute(body)
	assert.Nil(t, err)
	p1 := &Pair{}
	err = mapstructure.Decode(data, &p1)
	assert.Nil(t, err)
	assert.NotEqual(t, "", p1.English)
	assert.NotEqual(t, "", p1.Esperanto)

	// get another one, make sure it's not the same one
	code, body, _ = TestEndpoint(http.MethodGet, "/random", nil, GetRandomPhraseRoute, true)
	assert.Equal(t, http.StatusOK, code)
	data, err = unmarshalMapFromTestRoute(body)
	assert.Nil(t, err)
	p2 := &Pair{}
	err = mapstructure.Decode(data, &p2)
	assert.Nil(t, err)
	assert.NotEqual(t, "", p2.English)
	assert.NotEqual(t, "", p2.Esperanto)

	// try without auth
	code, _, _ = TestEndpoint(http.MethodGet, "/random", nil, GetRandomPhraseRoute, false)
	assert.Equal(t, http.StatusForbidden, code)
}
