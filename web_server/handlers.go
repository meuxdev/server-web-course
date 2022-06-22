package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello root! 🚀")
}

func HandleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the API endpoint")
}

func PostRequest(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var metadata MetaData
	// send decoder as reference
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	// !Pending error -> fmt.Printf("Name: ", metadata["Name"])
	// %v so that the message is formated
	fmt.Fprintf(w, "Payload: %v", metadata)
}
