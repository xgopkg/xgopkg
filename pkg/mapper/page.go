package mapper

import (
	"fmt"
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
	offset    int
	Sort      Sort
}

//Page
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
func (p *Pageable) Offset() int {
	if p.offset == 0 {
		p.offset = (p.PageIndex - 1) * p.PageSize

	}
	return p.offset
}

// func (p *Pageable) Sort() Sort {
// 	return p.Sort
// }
// func (p *Pageable) Offset() int {
// 	return (p.pageIndex - 1) * p.pageSize
// }

//Page page
type Page struct {
	PageIndex int
	PageSize  int
	Pages     int
	Total     int64
	Data      interface{}
}

//NewPage page
func NewPage(data interface{}) *Page {
	return &Page{}
}

//PageBuilder page builder
type PageBuilder struct {
	pageable *Pageable
	page     Page
}

//Page init page index & size
func (p *PageBuilder) Page(pa *Pageable) *PageBuilder {
	p.pageable = pa
	return p
}

//Total set total elments & total pages
func (p *PageBuilder) Total(total int64) *PageBuilder {
	p.page.Total = total
	return p
}

//Data set page data
func (p *PageBuilder) Data(data interface{}) *PageBuilder {
	p.page.Data = data
	return p
}

//Build return page struct
func (p *PageBuilder) Build() (*Page, error) {
	err := p.pageable.Check()
	if err != nil {
		return nil, fmt.Errorf("Pageable error,err=%s", err)
	}
	p.page.PageIndex = p.pageable.PageIndex
	p.page.PageSize = p.pageable.PageSize

	if p.page.Total%int64(p.page.PageSize) == 0 {
		p.page.Pages = int(p.page.Total / int64(p.page.PageSize))
	} else {
		p.page.Pages = int(p.page.Total/int64(p.page.PageSize)) + 1
	}

	return &(p.page), nil
}

//NewPageBuilder new page builder
func NewPageBuilder() *PageBuilder {
	return &PageBuilder{}
}

// PageableMapper pageable base mapper
type PageableMapper struct {
}

//PageBuilder get mapper pagebuilder
func (p *PageableMapper) PageBuilder() *PageBuilder {
	return NewPageBuilder()
}
