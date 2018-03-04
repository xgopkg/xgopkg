package migrate

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattes/migrate/database/mysql"
)

func TestStartMigrate(t *testing.T) {
	Migrate()
}
