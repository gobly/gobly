package main

import (
	"github.com/gobly/core"
	"github.com/gobly/help"
	"github.com/gobly/ui"
	"net/http"
	"os"
)

var layoutSingle = ui.LoadSingle("html/main.html")

func main() {
	router := core.NewRouter()
	router.AddGetHandler("/", rootHandler)

	// Initialize modules
	ui.CreateContext("/ui", router)
	help.CreateContext("/help", router)

	core.App.RegisterModule("Home", "/", "")

	core.ShowWelcome(os.Stdout, router)
	http.ListenAndServe(":8080", router)
}

func rootHandler(out http.ResponseWriter, _ *http.Request) {
	layoutSingle.Execute(out, nil)
}
