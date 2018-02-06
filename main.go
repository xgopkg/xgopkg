package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

const (
	//QTraderURL for qtrader.io home page URL
	QTraderURL = "https://www.xgopkg.com"
)

//PackageView package view struct
type PackageView struct {
	Title string
	Name  string
}

//Template template impl
type Template struct {
	templates *template.Template
}

//Render render template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		c.Redirect(301, QTraderURL)
		return nil
	})
	//support xgopkg.com/x,xgopkg.com/x/y,xgopkg.com/x/y/z, xgopkg/x/y/z/
	e.GET("/:pkg", handPkg)
	e.GET("/:pkg/:subPkg", handPkg)
	e.GET("/:pkg/:subPkg/:sSubPkg", handPkg)
	e.GET("/:pkg/:subPkg/:sSubPkg/:sSSubPkg", handPkg)

	e.Logger.Fatal(e.Start(":1323"))
}

func handPkg(c echo.Context) error {
	pkgName := c.Param("pkg")
	isGoGet := c.QueryParam("go-get")
	pkg := &PackageView{
		Title: pkgName,
		Name:  pkgName,
	}
	if pkg.Name != "" && isGoGet == "1" {
		return c.Render(http.StatusOK, "pkg.html", pkg)

	}
	c.Redirect(301, QTraderURL)
	return nil
}
