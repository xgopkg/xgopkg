package route

import (
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	_ "github.com/go-xorm/xorm"
	"gopkg.in/logger.v1"

	"xgopkg.com/xgopkg/pkg/assets"
	// "xgopkg.com/xgopkg/pkg/config"

	rs "xgopkg.com/xgopkg/pkg/resource"
)

//PackageRoute go get
func Package() {
	ws := new(restful.WebService)
	log.Debug("pkg route")
	ws.Route(ws.GET("/{*}").To(rs.Render))
}

//API route go-restful api
func API(mux *http.ServeMux) {
	ws := new(restful.WebService)
	ws.Path("/api/v1")
	container := restful.NewContainer()
	container.ServeMux = mux
	container.Add(ws)
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(),
		APIPath:     "/swagger.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	container.Add(restfulspec.NewOpenAPIService(config))
}

//WebUI route swagger & web.
func WebUI(mux *http.ServeMux) {
	mux.Handle("/apidocs/",
		http.StripPrefix("/apidocs/",
			http.FileServer(assets.FS("public/swagger/dist"))))
	mux.Handle("/",
		http.StripPrefix("/",
			http.FileServer(assets.FS("build/"))))
}
