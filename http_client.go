package main

import (
	"bufio"
	"log"
	"net/http"
)

type HTTPClient struct {
}

func NewHTTPClient() HTTPClient {
	return HTTPClient{}
}

func (id HTTPClient) get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	log.Println("Response status:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		log.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
}
