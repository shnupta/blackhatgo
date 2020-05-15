package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	// Print the HTTP status
	fmt.Println(resp.Status)

	// Read and display the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
