package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello root! 🚀")
}
