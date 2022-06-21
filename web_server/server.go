package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, pathExists := s.router.rules[path]
	if !pathExists {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, midlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range midlewares {
		f = middleware(f)
	}
	return f
}

func (s *Server) Listen() error {
	// param1 port | param2 handlers
	http.Handle("/", s.router)
	fmt.Printf("Server Running --- http://localhost%s\n", s.port)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
