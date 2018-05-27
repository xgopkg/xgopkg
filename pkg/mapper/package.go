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
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	UserID      string    `xorm:"user_id"`
}

//PackageMapper package orm mapper struct
type PackageMapper struct {
	PageableMapper
}

var packageMapper *PackageMapper

// DefaultPackageMapper  return single PackageMapper instance
func DefaultPackageMapper() *PackageMapper {
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

// GetByName find Package by name
func (p *PackageMapper) GetByName(name string) *Package {
	var pkg Package
	ok, err := engine.Alias("p").Where("p.pkg_name = ?", name).Get(&pkg)
	if err != nil || !ok {
		return nil
	}
	return &pkg
}

// FindByName find Package by name
func (p *PackageMapper) FindByName(name string, pa *Pageable) (*Page, error) {
	var pkg []Package
	engine.Alias("p").Where("p.pkg_name LIKE ?", "%"+name+"%").Limit(pa.PageSize, pa.Offset()).Find(&pkg)
	count, err := engine.Alias("p").Where("p.pkg_name LIKE %?%", name).Count(&Package{})
	if err != nil {
		//todo
		log.Error(err)
	}
	log.Debugf("total: %d", count)

	return p.PageBuilder().Page(pa).Data(pkg).Total(count).Build()
}

//FindAll find all pkg
func (p *PackageMapper) FindAll(pa *Pageable) (*Page, error) {
	var pkg []Package
	engine.Limit(pa.PageSize, pa.Offset()).Find(&pkg)
	count, err := engine.Count(&Package{})
	if err != nil {
		//todo
		log.Error(err)
	}
	log.Debugf("total: %d", count)

	return p.PageBuilder().Page(pa).Data(pkg).Total(count).Build()
}

//FindByUserID qeuery user by user_id
func (p *PackageMapper) FindByUserID(userID string, pa *Pageable) (*Page, error) {
	var pkg []Package
	engine.Alias("p").Where("p.user_id=?", userID).Limit(pa.PageSize, pa.Offset()).Find(&pkg)
	count, err := engine.Count(&Package{})
	if err != nil {
		//todo
		log.Error(err)
	}
	log.Debugf("total: %d", count)

	return p.PageBuilder().Page(pa).Data(pkg).Total(count).Build()
}

//FindByUserIDAndPkgName qeuery user by user_id
func (p *PackageMapper) FindByUserIDAndPkgName(userID, pkgName string, pa *Pageable) (*Page, error) {
	var pkg []Package
	engine.Alias("p").Where("p.user_id=?", userID).And("p.pkg_name LIKE ?", "%"+pkgName+"%").Limit(pa.PageSize, pa.Offset()).Find(&pkg)
	count, err := engine.Count(&Package{})
	if err != nil {
		//todo
		log.Error(err)
	}
	log.Debugf("total: %d", count)

	return p.PageBuilder().Page(pa).Data(pkg).Total(count).Build()
}

// FindByCondition dynamic query
func (p *PackageMapper) FindByCondition(condition *Package, pa *Pageable) (*Page, error) {
	var pkg []Package
	session := engine.Alias("p").Where("1=1")
	defer session.Close()
	if condition != nil {
		if condition.UserID != "" {
			session.And("p.user_id = ?", condition.UserID)
		}
		if condition.Name != "" {
			log.Debugf("pkg_name=%s", condition.Name)
			session.And("p.pkg_name LIKE ?", "%"+condition.Name+"%")
		}
	}
	session2 := session.Clone()
	defer session2.Close()
	count, err := session2.Count(&Package{})
	if err != nil {
		//todo
		log.Error(err)
	}

	session.Limit(pa.PageSize, pa.Offset()).Find(&pkg)
	return p.PageBuilder().Page(pa).Data(pkg).Total(count).Build()

}
