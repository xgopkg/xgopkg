package main

import (
	"net/http"

	"github.com/emicklei/go-restful"
	_ "github.com/go-xorm/xorm"
	_ "github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/mysql"
	_ "github.com/spf13/viper"
	"gopkg.in/logger.v1"

	"xgopkg.com/xgopkg/pkg/config"
	"xgopkg.com/xgopkg/pkg/mapper"
	"xgopkg.com/xgopkg/pkg/migrate"
	"xgopkg.com/xgopkg/pkg/route"
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
	//TODO: support flag
	// use flag:
	// just as --port ...

	//run server two for dashboard
	go func() {
		mux := http.NewServeMux()
		route.API(mux)
		route.WebUI(mux)
		log.Fatal(http.ListenAndServe(":8001", mux))
	}()
	route.Package()
	//Run main server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
