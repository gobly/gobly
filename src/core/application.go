package core

import (
	"path/filepath"
	"runtime"
	"path"
	"io"
	"fmt"
)

type Application struct {
	Version string
	Name string
	Root string
}

var App = Application {
	"0.0.1",
	"Gobly Engine",
	filepath.Dir(CallerPath(0)),
}

var CallerPath = func(skip int) string {
	_, filename, _, success := runtime.Caller(skip)
	if !success {
		panic("No caller information")
	}

	return path.Dir(filename)
}

var CurrentPath = func(packageName string) string {
	return filepath.Join(App.Root, packageName)
}

func ShowWelcome(out io.Writer, router *Router) {
	fmt.Fprintln(out, "Gobly Web Framweork is up and running!")
	fmt.Fprintln(out, "Activated routes: ")
	router.FPrint(out)
}
