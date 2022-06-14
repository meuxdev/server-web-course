package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type WriterWeb struct {
}

func (WriterWeb) Write(webBytes []byte) (int, error) {
	fmt.Println(string(webBytes))
	return len(webBytes), nil
}

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalln("Error with the get HTTP request")
	}
	reader := res.Body
	writer := WriterWeb{}
	io.Copy(writer, reader)
	//fmt.Println(res.Body)

}
