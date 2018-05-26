package resource

import (
	rest "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"gopkg.in/logger.v1"
)

// UserResource e resource
type UserResource struct {
}

// WebService  user webservice
func (u UserResource) WebService() *rest.WebService {
	tags := []string{"users"}
	ws := new(rest.WebService)

	ws.Path("/api/v1/users").
		Consumes(rest.MIME_JSON).
		Produces(rest.MIME_JSON)
	ws.Route(ws.GET("").
		To(u.getUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags))
	return ws
}
func (u UserResource) getUsers(request *rest.Request, res *rest.Response) {
	log.Debug("get pkgs")
	res.Write([]byte("[\"hello\"]"))

	//TODO: query DB

}

//NewUserResource new User resource
func NewUserResource() *UserResource {
	return &UserResource{}
}
