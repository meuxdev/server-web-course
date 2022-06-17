package main

import (
	"fmt"
	"net/http"
	"strings"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			flag := true
			fmt.Println("Checking authentication")
			if flag {
				f(w, req)
			} else {
				return
			}
		}
	}
}

func PrintRequest() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			fmt.Printf("New Request  Method --- %s --- %s\n", strings.ToUpper(req.Method), req.RequestURI)
			f(w, req)
		}
	}
}
