package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
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

func Logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() {
				log.Printf("Request \n------   Method --- %s --- %s --- Time Passed %s\n",
					strings.ToUpper(req.Method),
					req.RequestURI,
					time.Since(start))
			}()
			f(w, req)
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(req.URL.Path, time.Since(start))
			}()
			f(w, req)
		}
	}
}
