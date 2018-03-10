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
	names, err := AssetDir("test")
	assert.NoError(t, err)
	log.Debugf("test dir file names : %v", names)
	assert.NotEmpty(t, AssetNames())
	assert.Equal(t, true, len(names) > 1)
}
