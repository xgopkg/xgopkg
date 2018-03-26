package main

import (
	"net/http"
	"path"
	"runtime"

	// "github.com/labstack/echo"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	_ "github.com/go-xorm/xorm"
	_ "github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/mysql"
	_ "github.com/spf13/viper"
	"gopkg.in/logger.v1"

	"xgopkg.com/xgopkg/pkg/assets"
	"xgopkg.com/xgopkg/pkg/config"
	"xgopkg.com/xgopkg/pkg/mapper"
	"xgopkg.com/xgopkg/pkg/migrate"
	"xgopkg.com/xgopkg/pkg/resource"
)

// PackageResource xxx
type PackageResource struct {
}

// WebService  xx
func (p PackageResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/").To(p.hello))
	return ws
}
func (p PackageResource) hello(request *restful.Request, res *restful.Response) {
	res.AddHeader("Content-Type", restful.MIME_JSON)
	res.Write([]byte("[\"hello\"]"))
}
func init() {
	log.SetOutputLevel(config.LoggerLevel())
	mapper.Connect()
	migrate.Migrate()
}
func main() {
	p := PackageResource{}
	ws := new(restful.WebService)
	ws.Route(ws.GET("/{*}").To(resource.Render))

	restful.DefaultContainer.Add(p.WebService()).Add(ws)
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(),
		APIPath:     "/swagger.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	//todo env get swaagerr ui dis
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}
	log.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(assets.FS("public/swagger/dist"))))
	//run server two for dashboard
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/", http.StripPrefix("/", http.FileServer(assets.FS("build/"))))
		log.Fatal(http.ListenAndServe(":8001", mux))
	}()
	//Run main server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "XGoPkg",
			Description: "Resource for managing package",
			Contact: &spec.ContactInfo{
				Name:  "xgopkg",
				Email: "dev_support@xgopkg.com",
				URL:   "http://xgopkg.com",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}
