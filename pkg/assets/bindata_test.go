package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/logger.v1"
)

func init() {
	log.SetOutputLevel(0)
}
func TestAsset(t *testing.T) {
	data, err := Asset("public/swagger/dist/index.html")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, data)
}
