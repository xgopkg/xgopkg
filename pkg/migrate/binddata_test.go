package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/logger.v1"
)

func init() {
	log.SetOutputLevel(0)
}
func TestAsset(t *testing.T) {
	data, err := Asset("migrations/201803041120_add_package.sql")
	readme, err := Asset("README.MD")
	if err != nil {
		t.Error(err)
	}
	log.Debug("content: %s", string(data))
	assert.NotEmpty(t, data)
	assert.NotEmpty(t, readme)
}
