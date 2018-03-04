//Package migrate for migrate database
package migrate

import (
	// "database/sql"

	"gopkg.in/logger.v1"

	mg "github.com/mattes/migrate"
	"github.com/mattes/migrate/source/go-bindata"
	// "github.com/mattes/migrate/source/go-bindata/examples/migrations"
)

func StartMigrate() {
	// dbURL := config.GetMySQLURL() + "&multiStatements=true"
	// db, _ := sql.Open("mysql", dbURL)
	// driver, _ := mysql.WithInstance(db, &mysql.Config{})
	// m, _ := migrate.NewWithDatabaseInstance(
	// 	"file:///migrations",
	// 	"mysql",
	// 	driver,
	// )

	// m.Steps(2)
	s := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) {
			log.Info(name)
			return Asset(name)
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		log.Error(err)
		//todo
	}

	// migrate.NewWithD
	m, err := mg.NewWithSourceInstance("go-bindata", d, "mysql://travis:@/xgopkg?charset=utf8")
	if err != nil {
		log.Error(err)
		//todo
	}
	m.Up() // run your migrations and handle the errors above of course
}
