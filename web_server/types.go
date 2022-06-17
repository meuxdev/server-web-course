package main

import "net/http"

type Midleware func(http.HandlerFunc) http.HandlerFunc
