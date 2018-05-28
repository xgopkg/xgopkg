package page

import (
	"fmt"
	"reflect"

	"github.com/go-xorm/xorm"
)

// Sort sort
type Sort struct {
	Sortby string
	Order  string
}

// Pageable pageable
type Pageable struct {
	PageIndex int
	PageSize  int
	// offset    int
	Sort Sort
}

//SetDefault default page
func (p *Pageable) SetDefault() {
	if p.PageIndex == 0 {
		p.PageIndex = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10

	}
}

//Check verify page inde & page size.
func (p *Pageable) Check() error {
	if p.PageIndex == 0 {
		return fmt.Errorf("Pageable must begin with 1, actual is %d", p.PageIndex)
	}
	if p.PageSize == 0 {
		return fmt.Errorf("Page size can not be zero")
	}
	return nil
}

//Offset db start
func (p *Pageable) offset() int {
	return (p.PageIndex - 1) * p.PageSize
}

//Page page
type Page struct {
	PageIndex int         `json:"page_index,omitempty"`
	PageSize  int         `json:"page_size,omitempty"`
	Pages     int64       `json:"pages,omitempty"`
	Total     int64       `json:"total,omitempty"`
	Data      interface{} `json:"data"`
	dataType  interface{}
}

//NewPage page
func NewPage(data interface{}) *Page {
	return &Page{}
}

//Builder page builder
type Builder struct {
	pageable *Pageable
	page     Page
	session  *xorm.Session
}

//Page init page index & size
func (p *Builder) Page(pa *Pageable) *Builder {
	p.pageable = pa
	return p
}

//Total set total elments & total pages
func (p *Builder) total(total int64) *Builder {
	p.page.Total = total
	return p
}

//Data set page data
func (p *Builder) Data(dataType interface{}) *Builder {
	p.page.Data = dataType
	return p
}

//Session set xorm session
func (p *Builder) Session(session *xorm.Session) *Builder {
	p.session = session
	return p
}

//Build return page struct
func (p *Builder) Build() (*Page, error) {
	err := p.pageable.Check()
	if err != nil {
		return nil, fmt.Errorf("Pageable error,err=%s", err)
	}
	p.page.PageIndex = p.pageable.PageIndex
	p.page.PageSize = p.pageable.PageSize

	sc := p.session.Clone()
	if reflect.TypeOf(p.page.Data).Elem().Kind() != reflect.Slice {
		return nil, fmt.Errorf("Data() parameter must be a slice pointer,not %v", reflect.TypeOf(p.page.Data))
	}
	Logger().Debugf("type: %+v\n", reflect.TypeOf(p.page.Data))
	Logger().Debugf("type: %+v\n", reflect.TypeOf(p.page.Data).Elem().Elem())
	elementType := reflect.TypeOf(p.page.Data).Elem().Elem()
	element := reflect.New(elementType)
	count, err := sc.Count(element.Interface())
	if err != nil {
		return nil, fmt.Errorf("Get count failed,err=%s", err)
	}

	p.total(count)
	Logger().Debugf("limit: %d,%d", p.pageable.PageSize, p.pageable.offset())
	err = p.session.Limit(p.pageable.PageSize, p.pageable.offset()).Find(p.page.Data)
	if err != nil {
		Logger().Debug(err)
		return nil, err
	}

	Logger().Debugf("DATA :%+v\n", p.page.Data)
	Logger().Debugf("index:%d,size:%d\n", p.pageable.PageIndex, p.pageable.PageSize)

	if p.page.Total%int64(p.page.PageSize) == 0 {
		p.page.Pages = p.page.Total / int64(p.page.PageSize)
	} else {
		p.page.Pages = p.page.Total/int64(p.page.PageSize) + 1
	}
	defer p.session.Close()
	return &(p.page), nil
}

//NewBuilder new page builder
func NewBuilder() *Builder {
	return &Builder{}
}

// PageableMapper pageable base mapper
type PageableMapper struct {
}

//Builder get mapper Builder
func (p *PageableMapper) Builder() *Builder {
	return NewBuilder()
}
