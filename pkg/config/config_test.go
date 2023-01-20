package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	NewConfig()
	assert.NotNil(t, AppConfig)
	assert.Equal(t, AppConfig.Port, ":3000")
}
