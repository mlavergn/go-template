package main

import (
	"log"
	"net/http"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func (id Server) start() {
	http.HandleFunc("/", id.rootHandler)
	http.HandleFunc("/secure/", id.secureHandler)
	err := http.ListenAndServe(":80", nil)
	// err := http.ListenAndServeTLS(":443", "demo.crt", "demo.key", nil)
	if err != nil {
		log.Fatal("failed to start listener", err)
	}
}

func (id Server) rootHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	resp.Header().Set("Content-Type", "text/html")
	resp.Header().Set("Connection", "close")
	resp.Write([]byte("<head></head><body><h2>OK</h2></body>"))
}

func (id Server) secureHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	resp.Header().Set("Content-Type", "text/html")
	resp.Header().Set("Connection", "close")
	resp.Write([]byte("<head></head><body><h2>FORBIDDEN</h2></body>"))
}
