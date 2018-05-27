package route

import (
	"net/http"

	rest "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"gopkg.in/logger.v1"

	"xgopkg.com/xgopkg/pkg/assets"
	rs "xgopkg.com/xgopkg/pkg/resource"
)

//Package go get
func Package() {
	ws := new(rest.WebService)
	log.Debug("pkg route")
	ws.Route(ws.GET("/{*}").To(rs.Render))
}
func init() {
	APIServerContainer()
}

//API route go-restful api
func API() {

	// ws := new(restful.WebService)
	// ws.Path("/api/v1")
	APIServerContainer().
		Add(rs.NewPackageResource().WebService()).
		Add(rs.NewUserResource().WebService())

	config := restfulspec.Config{
		WebServices:                   APIServerContainer().RegisteredWebServices(),
		APIPath:                       "/swagger.json",
		APIVersion:                    "1",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	APIServerContainer().Add(restfulspec.NewOpenAPIService(config))

}

//WebUI route swagger & web.
func WebUI() {

	APIServerContainer().ServeMux.Handle("/apidocs/",
		http.StripPrefix("/apidocs/",
			http.FileServer(assets.FS("public/swagger/dist"))))

	// APIServerContainer().ServeMux.Handle("/",
	// 	http.StripPrefix("/",
	// 		http.FileServer(assets.FS("build/"))))

}

var apiContainer *rest.Container

//APIServerContainer get apiserver container
func APIServerContainer() *rest.Container {
	if apiContainer == nil {
		apiContainer = rest.NewContainer()
	}
	return apiContainer
}
