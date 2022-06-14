package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	initialTime := time.Now()
	var wg sync.WaitGroup
	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}
	workersAmount := len(servers)
	c := make(chan string, workersAmount)

	for _, serverUrl := range servers {
		wg.Add(1)
		//testServer(serverUrl, c, &wg)
		go testServer(serverUrl, c, &wg)
	}

	timePassed := time.Since(initialTime)
	fmt.Printf("---------------------------\nExc. Time was %s\n", timePassed)
	wg.Wait()

}

func testServer(serverURL string, c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(serverURL)
	fmt.Printf("Testing Server %s\n---------\n", serverURL)

	if err != nil {
		fmt.Println("Server no available. 😭")
		c <- ""

	} else {
		fmt.Println("Server working correctly. ✔️")
		c <- "OK"
	}

}
