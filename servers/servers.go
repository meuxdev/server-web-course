package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	initialTime := time.Now()

	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}

	for _, serverUrl := range servers {
		testServer(serverUrl)
	}

	timePassed := time.Since(initialTime)
	fmt.Printf("---------------------------\nExc. Time was %s\n", timePassed)
}

func testServer(serverURL string) {
	_, err := http.Get(serverURL)
	fmt.Printf("Testing Server %s\n---------\n", serverURL)

	if err != nil {
		fmt.Println("Server no available. 😭")
	} else {
		fmt.Println("Server working correctly. ✔️")
	}
}
