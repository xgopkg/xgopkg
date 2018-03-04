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
	log.Debugf("names : %v", AssetNames())
	assert.NotEmpty(t, AssetNames())
}
