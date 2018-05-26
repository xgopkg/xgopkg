package resource

import (
	rest "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"gopkg.in/logger.v1"
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
		Metadata(restfulspec.KeyOpenAPITags, tags))
	return ws
}
func (p PackageResource) getPackages(request *rest.Request, res *rest.Response) {
	log.Debug("get pkgs")
	res.Write([]byte("[\"hello\"]"))
	//TODO: query DB
}

//NewPackageResource new package resource
func NewPackageResource() *PackageResource {
	return &PackageResource{}
}
