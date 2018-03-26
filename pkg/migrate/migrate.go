//Package migrate for migrate database
package migrate

import (
	"path/filepath"

	mg "github.com/mattes/migrate"
	"github.com/mattes/migrate/source/go-bindata"
	"gopkg.in/logger.v1"

	"xgopkg.com/xgopkg/pkg/config"
)

//Migrate start migrations
func Migrate() {
	dbURL := config.MySQLURL()
	names, err := AssetDir("migrations")
	if err != nil {
		log.Error(err)
	}
	s := bindata.Resource(names,
		func(name string) ([]byte, error) {
			log.Info(name)
			return Asset(filepath.Join("migrations", name))
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		log.Error(err)
		//todo
	}

	log.Debug("database url: ", dbURL)

	// migrate.NewWithD
	m, err := mg.NewWithSourceInstance("go-bindata", d, "mysql://"+dbURL)
	if err != nil {
		log.Error(err)
		//todo
	}
	m.Up() // run your migrations and handle the errors above of course
}
