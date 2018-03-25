package mapper

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"xgopkg.com/xgopkg/pkg/migrate"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattes/migrate/database/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/logger.v1"
)

var pkgMapper *PackageMapper

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}
func setUp() {
	log.SetOutputLevel(0)
	connectDb()
	migrate.Migrate()
	importTestSQL("data.sql")
	log.Debug("setup")
	pkgMapper = &PackageMapper{}

}
func importTestSQL(filename string) {
	sql, err := migrate.Asset(filepath.Join("test", filename))
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	engine.Import(bytes.NewBuffer(sql))
}
func tearDown() {
	// importTestSQL("clean.sql")
	log.Debug("tear down: clean tables")
}

func TestPackageMapper_Save(t *testing.T) {
	pkg := &Package{
		Name:        "test",
		Description: "matrix 容器云",
		Source:      "https://git.changhong.io/chcloud/matrix/plugin/",
	}

	err := pkgMapper.Save(pkg)
	assert.NoError(t, err)
}
func TestPackageMapper_FindByName(t *testing.T) {
	pkg := pkgMapper.FindByName("matrix")
	log.Debugf("pkg is %+v", pkg)
	assert.Equal(t, "matrix", pkg.Name)
}
