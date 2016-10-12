package main

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Config(t *testing.T) {
	config := loadConfig()
	//test defaults
	assert.Equal(t, config.Port, ":8081")
	assert.NotEqual(t, config.AuthenticationToken, "")
}

func Test_LoadPhrases(t *testing.T) {
	phrases := loadPhrasebook()
	assert.True(t, len(phrases) > 0)
	pair := phrases[0]
	assert.NotEqual(t, pair.English, "")
	assert.NotEqual(t, pair.Esperanto, "")
}

func Test_GetPair(t *testing.T) {
	phrases := loadPhrasebook()
	c, w, _ := gin.CreateTestContext()
	GetRandomPair(c, phrases)
	assert.Equal(t, w.Code, 200)
	assert.NotEqual(t, w.Body.String(), "")
	assert.Equal(t, w.HeaderMap.Get("Content-Type"), "application/json; charset=utf-8")
}

func Test_GetAll(t *testing.T) {
	phrases := loadPhrasebook()
	c, w, _ := gin.CreateTestContext()
	GetAllWords(c, phrases)
	assert.Equal(t, w.Code, 200)
	assert.NotEqual(t, w.Body.String(), "")
	assert.Equal(t, w.HeaderMap.Get("Content-Type"), "application/json; charset=utf-8")
}

func Test_Post_Not_Authorized(t *testing.T) {
	c, w, _ := gin.CreateTestContext()
	c.Set("Authenticated", false)
	phrases := loadPhrasebook()
	SaveNewPair(c, phrases)
	assert.Equal(t, w.Code, 401)
	assert.NotEqual(t, w.Body.String(), "")
	assert.Equal(t, w.HeaderMap.Get("Content-Type"), "application/json; charset=utf-8")
}

func Test_Middleware_Fail(t *testing.T) {
	router := gin.New()
	router.Use(checkAuthentication())
	router.POST("/", func(c *gin.Context) {
		assert.False(t, c.MustGet("Authenticated").(bool))
	})
	_ = performRequest(router, "POST", "/")
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
