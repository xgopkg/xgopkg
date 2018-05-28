package resource

import (
	"reflect"
	"strconv"

	rest "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"gopkg.in/logger.v1"

	"github.com/xorm-page/page"
	"xgopkg.com/xgopkg/pkg/mapper"
)

// PackageResource package resource
type PackageResource struct {
}

// WebService  xx
func (p PackageResource) WebService() *rest.WebService {
	tags := []string{"packages"}
	ws := new(rest.WebService)

	ws.Path("/api/v1/packages").
		Consumes(rest.MIME_JSON).
		Produces(rest.MIME_JSON)
	ws.Route(ws.GET("").
		To(p.getPackages).
		Doc("get all packages").
		Param(ws.QueryParameter("user_id", "user id")).
		Param(ws.QueryParameter("name", "package name")).
		Param(ws.QueryParameter("page_index", "page index").DataType("number").DefaultValue("1")).
		Param(ws.QueryParameter("page_size", "page size").DataType("number").DefaultValue("10")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(page.Page{Data: []mapper.Package{}}))
	return ws
}
func (p PackageResource) getPackages(request *rest.Request, res *rest.Response) {
	log.Debug("get pkgs")
	userID := request.QueryParameter("user_id")
	pkgName := request.QueryParameter("name")
	pageIndex, err := strconv.Atoi(request.QueryParameter("page_index"))
	if err != nil {
		//todo
		log.Error(err)
	}

	pageSize, err := strconv.Atoi(request.QueryParameter("page_size"))
	if err != nil {
		//todo
		log.Error(err)
	}

	pageable := &page.Pageable{PageIndex: pageIndex, PageSize: pageSize}
	if err != nil {
		//todo
		log.Error(err)
	}

	//TODO:
	//Dynamic search conditions build
	condition := new(mapper.Package)
	condition.Name = pkgName
	condition.UserID = userID
	page, err := mapper.DefaultPackageMapper().FindByCondition(condition, pageable)
	if err != nil {
		//todo
		log.Error(err)
	}

	log.Debugf("data type%v", reflect.TypeOf(page.Data))
	// list := page.Data.(reflect.Value).Interface().(*[]mapper.Package)
	// page.Data = list
	// data, err := json.Marshal(page)
	// log.Debugf("data: %+v", page.Data)
	// // log.Debugf("list: %+v", list)
	// if err != nil {
	// 	//todo
	// 	log.Error(err)
	// }
	res.WriteEntity(page)

	//TODO: query DB
}

//NewPackageResource new package resource
func NewPackageResource() *PackageResource {
	return &PackageResource{}
}
