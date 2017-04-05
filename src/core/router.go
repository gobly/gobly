package core

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"io"
	"strings"
	"path/filepath"
)

type Router struct {
	mux *mux.Router
}

func NewRouter() *Router {
	router := Router{
		mux.NewRouter(),
	}

	router.mux.StrictSlash(true)
	return &router
}

func (r *Router) AddHandler(path string, callback func(http.ResponseWriter, *http.Request)) {
	r.mux.HandleFunc(path, callback)
}

func (r *Router) AddSubRouter(prefix string, callback func(*Router)) {
	s := r.mux.PathPrefix(prefix).Subrouter()
	callback(&Router{s})
}

func (r *Router) AddStatic(prefix string, folder string) {
	callerPath := CallerPath(2)
	r.mux.PathPrefix(prefix).HandlerFunc(func(out http.ResponseWriter, in *http.Request) {
		rpath := strings.SplitAfterN(in.RequestURI, prefix, 2)

		if len(rpath) < 2 {
			http.NotFound(out, in)
			return
		}

		http.ServeFile(out, in, filepath.Join(callerPath, folder, rpath[1]))
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) FPrint(out io.Writer) {
	r.mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		fmt.Fprintln(out, t, route.GetHandler())
		return nil
	})
}