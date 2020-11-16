package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/blackestwhite/crawler/getch"
)

func main() {

	resp, err := http.Get("https://www.example.com")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	defer resp.Body.Close()

	// Read body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	theDoc := string(body[:])
	res, err := getch.Get("textarea", theDoc)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range res {
		fmt.Println(val)
	}

	links, err := getch.GetLinks(theDoc)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range links {
		fmt.Println(val)
	}
}
