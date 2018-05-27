package resource

import (
	"strings"

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
	log.Debug(req.Request.URL)

	if req.QueryParameter("go-get") != "1" {
		resp.WriteHeader(400)
		return
	}
	pathQueryParams := strings.Split(req.Request.URL.String(), "?")

	// TODO: 从URL第一层开始解析 匹配数据库 支持 子包

	/* 	FIXME: /matrix/node/api
	 *  check /matrix from pkg table
	 *  if not existed  check /matrix/node
	 *  if not existed  check /matrix/node/api
	 */
	if len(pathQueryParams) < 2 {
		resp.WriteHeader(400)
		return
	}
	path := pathQueryParams[0]
	index := strings.Index(path, "/")
	lastIndex := strings.LastIndex(path, "/")
	var iPath string
	if index == -1 {
		resp.WriteHeader(400)
		return
	}

	//Remove last '/'
	//For example  /matrix/types/?go-get=1 to  /matrix/types
	if strings.HasSuffix(path, "/") {
		iPath = path[0:lastIndex]
	} else {
		iPath = path
	}
	log.Debugf("First / index = %d,last index = %d, path length = %d",
		index, lastIndex, len(path))
	log.Debug(iPath)
	// iPathArr := strings.Split(iPath, "/")
	pkg, _ := findPkg(iPath)
	log.Debugf("pkg %+v", pkg)
	// pkg := mapper.PackageMapperInst  ance().FindByName(pathQueryParams[0])
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
	log.Debugf("tmpl text: %s", text)
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	tmpl.Execute(resp.ResponseWriter, p)
}

// 递归查询包
func findPkg(iPath string) (*mapper.Package, string) {
	// TODO: 跳出条件
	// find a pkg in table or last item in iPath
	log.Debug(iPath)
	pkg := mapper.DefaultPackageMapper().GetByName(iPath)
	if pkg != nil {
		return pkg, ""
	}
	if strings.LastIndex(iPath, "/") == 0 {
		log.Debug("Package not found")
		return nil, ""
	}
	return findPkg(iPath[0:strings.LastIndex(iPath, "/")])
}
