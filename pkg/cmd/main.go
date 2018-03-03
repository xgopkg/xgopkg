package main

import (
	// "html/template"
	// "io"
	"log"
	"net/http"

	// "github.com/labstack/echo"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	_ "github.com/go-xorm/xorm"
	_ "github.com/mattes/migrate"
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

const (
	//HomeURL for qtrader.io home page URL
	HomeURL = "https://www.xgopkg.com"
)

func main() {
	p := PackageResource{}
	restful.DefaultContainer.Add(p.WebService())
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(),
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	//todo env get swaagerr ui dis
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/liwei/JSWorkspace/swagger-ui/dist"))))
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
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}

//PackageView package view struct
// type PackageView struct {
// 	Title string
// 	Group string
// 	Name  string
// }

//Template template impl
// type Template struct {
// templates *template.Template
// }

//Render render template
// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }
// func main() {
// t := &Template{
// 	templates: template.Must(template.ParseGlob("public/views/*.html")),
// }

// e := echo.New()
// e.Renderer = t
// e.GET("/", func(c echo.Context) error {
// c.Redirect(301, HomeURL)
// return nil
// })
//support xgopkg.com/x,xgopkg.com/x/y,xgopkg.com/x/y/z, xgopkg/x/y/z/
// e.GET("/:group/:pkg", handPkg)
// e.GET("/:group/:pkg/:subPkg", handPkg)
// e.GET("/:group/:pkg/:subPkg/:sSubPkg", handPkg)
// e.GET("/:group/:pkg/:subPkg/:sSubPkg/:sSSubPkg", handPkg)
// e.GET("/:group/:pkg/**", handPkg)

// e.Logger.Fatal(e.Start(":1323"))
// }

// func handPkg(c echo.Context) error {
// 	groupName := c.Param("group")
// 	pkgName := c.Param("pkg")
// 	isGoGet := c.QueryParam("go-get")
// 	pkg := &PackageView{
// 		Title: pkgName,
// 		Group: groupName,
// 		Name:  pkgName,
// 	}
// 	if pkg.Name != "" && isGoGet == "1" {
// 		return c.Render(http.StatusOK, "pkg.html", pkg)
// 	}
// 	c.Redirect(301, HomeURL)
// 	return nil
// }
