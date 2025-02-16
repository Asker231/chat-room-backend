package main

import (
	"log"
	"net/http"
)


func main() {

	//init Handler
	router := http.NewServeMux()

	//init server 
	server := http.Server{
		Addr: ":80",
		Handler: router,
	}

	log.Panic(server.ListenAndServe())
}