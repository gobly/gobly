package ui

import (
	"html/template"
	"path"
	"path/filepath"
	"strings"
	"core"
)

// Package path depends on the caller, this cannot be determined on loading
var packagePath = func() string {
	pkgPath, err := filepath.Rel(core.App.Root, core.CallerPath(3))
	if err != nil {
		panic("Cannot determine package path from " + pkgPath)
	}

	// Get first component from pkgPath, and prepend absolute path of the source directory
	return filepath.Join(core.App.Root, strings.Split(filepath.ToSlash(pkgPath), "/")[0])
}

var funcMap = template.FuncMap{
	"App" : func() core.Application { return core.App },
}

func LoadSingle(mainTemplate string) *template.Template {
	mainTemplate = path.Join(packagePath(), mainTemplate)
	footerTemplate := path.Join(core.CurrentPath("ui"), "html", "footer.html")
	singleTemplate := path.Join(core.CurrentPath("ui"), "html", "single.html")

	return template.Must(template.New("single.html").Funcs(funcMap).ParseFiles(singleTemplate, mainTemplate, footerTemplate))
}

func LoadDouble(mainTemplate string, sidebarTemplate string) *template.Template {
	mainTemplate = path.Join(packagePath(), mainTemplate)
	sidebarTemplate = path.Join(packagePath(), sidebarTemplate)
	footerTemplate := path.Join(core.CurrentPath("ui"), "html", "footer.html")
	doubleTemplate := path.Join(core.CurrentPath("ui"), "html", "double.html")

	return template.Must(template.New("double.html").Funcs(funcMap).ParseFiles(doubleTemplate, mainTemplate, sidebarTemplate, footerTemplate))
}
