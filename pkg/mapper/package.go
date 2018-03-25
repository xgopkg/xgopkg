package mapper

import (
	"time"

	"gopkg.in/logger.v1"
)

//Package package ORM table struct
type Package struct {
	ID          int64     `xorm:"pk autoincr 'id'"`
	Name        string    `xorm:"char(100) not null 'pkg_name'"`
	Source      string    `xorm:"'pkg_source'"`
	Description string    `xorm:"varchar(255) 'pkg_desc'"`
	CreatedAt   time.Time `xorm:"not null created"`
	UpdatedAt   time.Time `xorm:"not null updated"`
}

//PackageMapper package orm mapper struct
type PackageMapper struct {
}

var packageMapper *PackageMapper

// PackageMapperInstance  return single PackageMapper instance
func PackageMapperInstance() *PackageMapper {
	if packageMapper == nil {
		return &PackageMapper{}
	}
	return packageMapper
}

//Save insert pkg to table
func (p *PackageMapper) Save(pkg *Package) error {
	re, err := engine.InsertOne(pkg)
	if err != nil {
		log.Errorf("An error occurred while saving pkg to db: error= %s", err)
		return err
	}
	log.Debugf("Insert pkg to db: result= %d", re)
	return nil
}

// FindByName find Package by name
func (p *PackageMapper) FindByName(name string) *Package {
	var pkg Package
	ok, err := engine.Alias("p").Where("p.pkg_name = ?", name).Get(&pkg)
	if err != nil || !ok {
		return nil
	}
	return &pkg
}
