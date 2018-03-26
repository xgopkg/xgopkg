package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
}
func TestMySQLURL(t *testing.T) {
	mURL := MySQLURL()
	assert.NotEmpty(t, mURL)
}
func TestLoggerLevel(t *testing.T) {
	level := LoggerLevel()
	assert.Equal(t, 0, level)
}
