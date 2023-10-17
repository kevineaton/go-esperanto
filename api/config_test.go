package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	LoadConfig()
	//test defaults
	assert.Equal(t, config.Port, ":8081")
	assert.NotEqual(t, config.AuthenticationToken, "")
}
