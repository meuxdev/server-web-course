package main

import (
	"fmt"
	"net/http"
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
