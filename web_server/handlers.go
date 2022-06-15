package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	httpStatus := http.StatusOK
	fmt.Printf("Status: %d --- %s --- %s\n", httpStatus, req.Method, req.RequestURI)
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, "Hello root! 🚀")
}

func HandleHome(w http.ResponseWriter, req *http.Request) {
	httpStatus := http.StatusAccepted
	fmt.Printf("Status: %d --- %s --- %s\n", httpStatus, req.Method, req.RequestURI)
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, "This is the API endpoint")
}
