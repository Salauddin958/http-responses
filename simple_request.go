package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	requestBody := strings.NewReader(`
	{
		"name": "robo",
		"memory": "100TB",
		"ram": "100GB"
	}`)

	request, err := http.NewRequest("POST", "https://dummy.restapiexample.com/api/v1/create", requestBody)
	if err != nil {
		log.Fatal("Error in new request ", err)
	}
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal("error in response ", err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error in reading response ", err)
	}
	defer response.Body.Close()
	fmt.Println("response status code ", response.StatusCode)
	fmt.Printf("actual response: %s", data)
}
