# page
- go get pkg & import
``` go
go get github.com/xorm-page/page
```
- query
``` go
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
```
- Response body for rest api by unmarshal 'Page' struct.
```json
{
  "page_index": 2,
  "page_size": 3,
  "pages": 2,
  "total": 6,
  "data": [
    {
      "id": 4,
      "name": "test5",
      "source": "https://github.com/xogopkg/test",
      "description": "test pkg",
      "created_at": "2018-03-09T21:12:19+08:00",
      "updated_at": "2018-03-09T21:12:19+08:00",
      "user_id": "1a08b350-0ed1-4a59-b0c9-5706882bd19b"
    },
    {
      "id": 5,
      "name": "test6",
      "source": "https://github.com/xogopkg/test",
      "description": "test pkg",
      "created_at": "2018-03-09T21:12:19+08:00",
      "updated_at": "2018-03-09T21:12:19+08:00",
      "user_id": "1a08b350-0ed1-4a59-b0c9-5706882bd19b"
    },
    {
      "id": 6,
      "name": "test7",
      "source": "https://github.com/xogopkg/test",
      "description": "test pkg",
      "created_at": "2018-03-09T21:12:19+08:00",
      "updated_at": "2018-03-09T21:12:19+08:00",
      "user_id": "1a08b350-0ed1-4a59-b0c9-5706882bd19b"
    }
  ]
}
```