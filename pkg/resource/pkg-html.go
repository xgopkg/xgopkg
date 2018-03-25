package resource

import (
	"strings"

	// "github.com/arschles/go-bindata-html-template"
	"github.com/emicklei/go-restful"
	"gopkg.in/logger.v1"
	"html/template"

	"xgopkg.com/xgopkg/pkg/assets"
	"xgopkg.com/xgopkg/pkg/mapper"
)

//PackageView package view struct
type PackageView struct {
	Host   string
	Source string
	Name   string
}

// Render render pkg
func Render(req *restful.Request, resp *restful.Response) {
	if req.QueryParameter("go-get") != "1" {
		resp.WriteHeader(400)
		return
	}
	log.Debug(req.Request.URL)
	urls := strings.Split(req.Request.URL.String(), "?")

	// TODO: 从URL第一层开始解析 匹配数据库 支持 子包

	pkg := mapper.PackageMapperInstance().FindByName(urls[0])
	if pkg == nil {
		resp.WriteHeader(404)
		return
	}
	p := &PackageView{
		Host:   "changhong.io",
		Source: pkg.Source,
		Name:   pkg.Name,
	}
	// you might want to cache compiled templates
	text, err := assets.Asset("public/views/pkg.html")
	if err != nil {
		log.Error(err)
	}
	tmpl, err := template.New("pkgTemplate").Parse(string(text))
	// tmpl, err := template.New("mytmpl", assets.Asset).Parse("public/views/pkg.html")
	log.Debugf("tmpl text: %s", text)
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	tmpl.Execute(resp.ResponseWriter, p)
}
