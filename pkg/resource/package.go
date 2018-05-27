package resource

import (
	"encoding/json"
	"strconv"

	rest "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"gopkg.in/logger.v1"

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
		Writes(mapper.Page{}))
	return ws
}
func (p PackageResource) getPackages(request *rest.Request, res *rest.Response) {
	log.Debug("get pkgs")
	userID := request.QueryParameter("user_id")
	pkgName := request.QueryParameter("name")
	pageIndex, err := strconv.Atoi(request.QueryParameter("page_index"))
	pageSize, err := strconv.Atoi(request.QueryParameter("page_size"))
	pageable := &mapper.Pageable{PageIndex: pageIndex, PageSize: pageSize}
	if err != nil {
		//todo
	}

	//TODO:
	//Dynamic search conditions build
	condition := new(mapper.Package)
	condition.Name = pkgName
	condition.UserID = userID
	page, err := mapper.DefaultPackageMapper().FindByCondition(condition, pageable)
	if err != nil {
		//todo
	}
	data, err := json.Marshal(page)
	if err != nil {
		//todo
	}
	res.Write(data)
	return

	//TODO: query DB
}

//NewPackageResource new package resource
func NewPackageResource() *PackageResource {
	return &PackageResource{}
}
