package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, pathExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, pathExist := r.FindHandler(request.URL.Path, request.Method)

	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		// URL exists but not with that method
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)

}
