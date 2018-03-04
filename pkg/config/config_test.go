package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
}
func TestGetMySQLURL(t *testing.T) {
	mURL := GetMySQLURL()
	assert.NotEmpty(t, mURL)
}
