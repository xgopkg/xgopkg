package mapper

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "xgopkg.com/xgopkg/pkg/config"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	Connect()
	info, err := engine.DBMetas()
	assert.NoError(t, err)
	assert.NotNil(t, info)
}
