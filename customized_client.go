package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 100 * time.Microsecond,
		// Timeout: 100 * time.Second, -> timeout will not happen
	}
	response, err := client.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		urlErr := err.(*url.Error)
		if urlErr.Timeout() {
			fmt.Println("Error occured due to timeout")
		}
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("Success : status-code", response.StatusCode)
	}

}
