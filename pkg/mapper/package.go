package mapper

import (
	"time"

	"github.com/xorm-page/page"
	"gopkg.in/logger.v1"
)

//Package package ORM table struct
type Package struct {
	ID          int64     `xorm:"pk autoincr 'id'" json:"id,omitempty"`
	Name        string    `xorm:"char(100) not null 'pkg_name'" json:"name,omitempty"`
	Source      string    `xorm:"'pkg_source'" json:"source,omitempty"`
	Description string    `xorm:"varchar(255) 'pkg_desc'" json:"description,omitempty"`
	CreatedAt   time.Time `xorm:"created" json:"created_at,omitempty"`
	UpdatedAt   time.Time `xorm:"updated" json:"updated_at,omitempty"`
	UserID      string    `xorm:"user_id" json:"user_id,omitempty"`
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
func (p *PackageMapper) FindAll(pa *page.Pageable) (*page.Page, error) {
	var pkgs []Package
	sess := engine.NewSession()
	defer sess.Close()
	return page.NewBuilder().Session(engine.NewSession()).Page(pa).Data(&pkgs).Build()
}

//FindByUserID qeuery user by user_id
func (p *PackageMapper) FindByUserID(userID string, pa *page.Pageable) (*page.Page, error) {
	var pkgs []Package
	sess := engine.Alias("p").Where("p.user_id=?", userID)
	return page.NewBuilder().Session(sess).Page(pa).Data(&pkgs).Build()
}

//FindByUserIDAndPkgName qeuery user by user_id
func (p *PackageMapper) FindByUserIDAndPkgName(userID, pkgName string, pa *page.Pageable) (*page.Page, error) {
	var pkgs []Package
	sess := engine.Alias("p").Where("p.user_id=?", userID).And("p.pkg_name LIKE ?", "%"+pkgName+"%")
	return page.NewBuilder().Page(pa).Session(sess).Data(&pkgs).Build()

}

// FindByCondition dynamic query
func (p *PackageMapper) FindByCondition(condition *Package, pa *page.Pageable) (*page.Page, error) {
	var pkgs []Package
	session := engine.Alias("p").Where("1=1")
	if condition != nil {
		if condition.UserID != "" {
			session.And("p.user_id = ?", condition.UserID)
		}
		if condition.Name != "" {
			log.Debugf("pkg_name=%s", condition.Name)
			session.And("p.pkg_name LIKE ?", "%"+condition.Name+"%")
		}
	}

	return page.NewBuilder().Page(pa).Session(session).Data(&pkgs).Build()
}
