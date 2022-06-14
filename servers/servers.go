package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	initialTime := time.Now()
	c := make(chan string)

	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}

	for _, serverUrl := range servers {
		go testServer(serverUrl, c)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-c)
	}
	timePassed := time.Since(initialTime)
	fmt.Printf("---------------------------\nExc. Time was %s\n", timePassed)

}

func testServer(serverURL string, c chan<- string) {
	_, err := http.Get(serverURL)

	if err != nil {
		c <- serverURL + " ===  Server NOT working. ❌"
	} else {
		c <- serverURL + " ===  Server working. ✔️"
	}

}
