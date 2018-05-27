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
	importTestSQL("clean.sql")
	log.Debug("tear down: clean tables")
}

func TestPackageMapper_Save(t *testing.T) {
	pkg := &Package{
		Name:        "test",
		Description: "test",
		UserID:      "8ac48532-7870-46d2-96cb-6eb273669806",
		Source:      "https://github.com/xgopkg/testpkg/",
	}

	err := pkgMapper.Save(pkg)
	assert.NoError(t, err)
}
func TestPackageMapper_FindByName(t *testing.T) {
	pkg := pkgMapper.FindByName("test2")
	log.Debugf("pkg is %+v", pkg)
	assert.Equal(t, "test2", pkg.Name)
}

func TestPackageMapper_FindByUserID(t *testing.T) {
	pageable := &Pageable{}
	pageable.PageIndex = 1
	pageable.PageSize = 10
	page, err := pkgMapper.FindByUserID("1a08b350-0ed1-4a59-b0c9-5706882bd19b", pageable)
	assert.NoError(t, err)
	log.Debugf("pkg page:%+v", page)
	log.Debugf("pkg page:%+v", page.Data)
	assert.NotNil(t, page)
	assert.NotEmpty(t, page.Data)
	assert.Equal(t, "test2", page.Data.([]Package)[0].Name)
}
