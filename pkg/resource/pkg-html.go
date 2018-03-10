package resource

import (
	// "github.com/arschles/go-bindata-html-template"
	"github.com/emicklei/go-restful"
	"gopkg.in/logger.v1"
	"html/template"

	"xgopkg.com/xgopkg/pkg/assets"
)

//PackageView package view struct
type PackageView struct {
	Host   string
	Source string
	Name   string
}

// Render render pkg
func Render(req *restful.Request, resp *restful.Response) {
	p := &PackageView{
		Host:   "changhong.io",
		Source: "https://git.changhong.io/chcloud/matrix/plugin/node",
		Name:   "node",
	}
	// you might want to cache compiled templates
	text, err := assets.Asset("public/views/pkg.html")
	tmpl, err := template.New("pkgTemplate").Parse(string(text))
	// tmpl, err := template.New("mytmpl", assets.Asset).Parse("public/views/pkg.html")
	log.Debugf("tmpl text: %s", text)
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	tmpl.Execute(resp.ResponseWriter, p)
}
