package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Midleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			flag := false
			fmt.Println("Checking authentication")
			if flag {
				f(w, req)
			} else {
				return
			}
		}
	}
}
