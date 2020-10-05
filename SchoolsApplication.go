package main

import (
	router2 "go-module-example/router"
	"log"
	"net/http"
)

func main() {
	handlerRequest()
}

func handlerRequest() {
	router := router2.MakeRouter()
	log.Fatal(http.ListenAndServe(":1000", router))
}